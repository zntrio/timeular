package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"time"

	"github.com/gofrs/uuid"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"tinygo.org/x/bluetooth"
	"zntr.io/timeular"
	eventsv1 "zntr.io/timeular/api/timeular/events/v1"
)

var (
	currentAdapter = bluetooth.DefaultAdapter
)

const (
	deviceName = "Timeular Tra"
)

func main() {
	// Initialize default logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: false,
	})).With(
		"service", "timeulard",
	)
	slog.SetDefault(logger)

	if err := run(); err != nil {
		logger.Error("unable to run the application", err)
	}
}

func run() error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	// Start the bluetooth device
	if err := currentAdapter.Enable(); err != nil {
		return fmt.Errorf("unable to enable to bluetooth device: %w", err)
	}

	slog.InfoContext(ctx, "Trying to connect to the device", "deviceName", deviceName)

	ch := make(chan bluetooth.ScanResult, 1)

	// // Try to discover the tracker.
	address := ""
	if err := currentAdapter.Scan(func(a *bluetooth.Adapter, sr bluetooth.ScanResult) {
		slog.DebugContext(ctx, "Identified device", "deviceName", sr.LocalName(), "address", sr.Address.String())

		if deviceName == sr.Address.String() || deviceName == sr.LocalName() {
			address = sr.Address.String()
			slog.InfoContext(ctx, "Tracker identified", "deviceName", deviceName, "address", address)
			if err := a.StopScan(); err != nil {
				panic(err)
			}
			ch <- sr
		}
	}); err != nil {
		return fmt.Errorf("unable to scan for bluetooth devices: %w", err)
	}

	var (
		device        *bluetooth.Device
		connectionErr error
	)

	// Connect to the device.
	select {
	case result := <-ch:
		device, connectionErr = currentAdapter.Connect(result.Address, bluetooth.ConnectionParams{})
		if connectionErr != nil {
			return connectionErr
		}

		slog.InfoContext(ctx, "Connected to the tracker", "address", result.Address.String())
	}

	// Wrap the device connection with service decorator.
	t, err := timeular.New(device)
	if err != nil {
		return fmt.Errorf("unable to initialize timeular driver: %w", err)
	}

	// Retrieve initial face orientation.
	faceID, err := t.GetOrientation()
	if err != nil {
		return fmt.Errorf("unable to retrieve initial orientation: %w", err)
	}
	slog.DebugContext(ctx, "Initial orientation", "faceID", faceID)

	// Prepare send queue
	sendQueue := make(chan *eventsv1.Event, 100)

	// Send initial orientation
	sendQueue <- &eventsv1.Event{
		EventId:   uuid.Must(uuid.NewV7()).Bytes(),
		Timestamp: uint64(time.Now().Unix()),
		EventType: eventsv1.EventType_EVENT_TYPE_ORIENTATION_CHANGED,
		Payload: &eventsv1.Event_OrientationChanged{
			OrientationChanged: &eventsv1.OrientationChangedPayload{
				FaceId: uint32(faceID),
			},
		},
	}

	// Register orientation change callback.
	t.OnOrientationChanged(func(u uint8) {
		// Send notification via pubsub/etc.
		slog.DebugContext(ctx, "Orientation changed", "faceID", u)
		sendQueue <- &eventsv1.Event{
			EventId:   uuid.Must(uuid.NewV7()).Bytes(),
			Timestamp: uint64(time.Now().Unix()),
			EventType: eventsv1.EventType_EVENT_TYPE_ORIENTATION_CHANGED,
			Payload: &eventsv1.Event_OrientationChanged{
				OrientationChanged: &eventsv1.OrientationChangedPayload{
					FaceId: uint32(u),
				},
			},
		}
	})

	// Register battery level change callback.
	t.OnBatteryLevelChanged(func(u uint8) {
		// Send notification via pubsub/etc.
		slog.DebugContext(ctx, "Battery level changed", "level", u)
		sendQueue <- &eventsv1.Event{
			EventId:   uuid.Must(uuid.NewV7()).Bytes(),
			Timestamp: uint64(time.Now().Unix()),
			EventType: eventsv1.EventType_EVENT_TYPE_BATTERY_LEVEL_CHANGED,
			Payload: &eventsv1.Event_BatteryLevelChanged{
				BatteryLevelChanged: &eventsv1.BatteryLevelChangedPayload{
					Level: uint32(u),
				},
			},
		}
	})

	// Create the grpc stream server
	server := grpc.NewServer()
	eventsv1.RegisterEventHubServiceServer(server, &eventHub{
		sendQueue: sendQueue,
	})

	// Run the event monitor.
	eg, egCtx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		slog.InfoContext(ctx, "Publisher started and listening")

		// Prepare listener
		lis, err := net.Listen("tcp", ":27045")
		if err != nil {
			return fmt.Errorf("unable to initialize timeular hub: %w", err)
		}

		return server.Serve(lis)
	})

	eg.Go(func() error {
		return t.Run(egCtx)
	})

	eg.Go(func() error {
		<-egCtx.Done()
		// Close the gRPC service
		server.GracefulStop()
		return nil
	})

	if err := eg.Wait(); err != nil {
		if !errors.Is(err, context.Canceled) {
			return fmt.Errorf("error occured in event monitor: %w", err)
		}
	}

	// Close the broadcast channel
	close(sendQueue)

	// Disconnect the bluetooth device
	if err := device.Disconnect(); err != nil {
		return fmt.Errorf("unable to successfully disconnect from the device: %w", err)
	}

	return nil
}

type eventHub struct {
	eventsv1.EventHubServiceServer

	sendQueue chan *eventsv1.Event
}

func (hub *eventHub) Subscribe(listener eventsv1.EventHubService_SubscribeServer) error {
	ctx := listener.Context()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case m := <-hub.sendQueue:
			// Send to the listener
			if err := listener.Send(&eventsv1.SubscribeResponse{
				Event: m,
			}); err != nil {
				return fmt.Errorf("unable to broadcast event to the listener: %w", err)
			}
		}
	}
}
