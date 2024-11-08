package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"example.com/myapp/api"
)

func initLog() {
	logOptions := &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	}
	handler := slog.NewTextHandler(os.Stderr, logOptions)
	slog.SetDefault(slog.New(handler))
	slog.Debug("set debug level")
}

func main() {
	initLog()

	ctx, stop := signal.NotifyContext(context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer stop()

	api.StartServer(ctx)
}
