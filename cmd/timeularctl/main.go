package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"os"
	"os/signal"

	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	eventsv1 "zntr.io/timeular/api/timeular/events/v1"
)

func main() {
	// Initialize default logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: false,
	})).With(
		"service", "timeularctl",
	)
	slog.SetDefault(logger)

	if err := run(); err != nil {
		logger.Error("unable to run the application", "error", err)
	}
}

func run() error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Kill, os.Interrupt)
	defer cancel()

	c, err := grpc.Dial("127.0.0.1:27045", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}

	srv := eventsv1.NewEventHubServiceClient(c)

	// Subscribe to events
	stream, err := srv.Subscribe(ctx)
	if err != nil {
		panic(err)
	}

	eg, _ := errgroup.WithContext(ctx)

	eg.Go(func() error {
		for {
			resp, err := stream.Recv()
			if err != nil {
				if errors.Is(err, io.EOF) {
					return nil
				}
				return fmt.Errorf("unable to receive message: %w", err)
			}

			slog.Info("Message received", "msg", resp.Event)
		}
	})

	return eg.Wait()
}
