package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"os/signal"

	"golang.org/x/sync/errgroup"
	"tinygo.org/x/bluetooth"
	"zntr.io/timeular"
)

var (
	currentAdapter = bluetooth.DefaultAdapter
)

const (
	deviceName                = "Timeular Tra"
)

func main() {
	// Initialize default logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
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
		device *bluetooth.Device
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

	// Register orientation change callback.
	t.OnOrientationChanged(func(u uint8) {
		// Send notification via pubsub/etc.
		slog.DebugContext(ctx, "Orientation changed", "faceID", u)
	})

	// Run the event monitor.
	eg, egCtx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		return t.Run(egCtx)
	})

	if err := eg.Wait(); err != nil {
		if !errors.Is(err, context.Canceled) {
			return fmt.Errorf("error occured in event monitor: %w", err)
		}
	}

	// Disconnect the bluetooth device
	if err := device.Disconnect(); err != nil {
		return fmt.Errorf("unable to successfully disconnect from the device: %w", err)
	}

	return nil
}