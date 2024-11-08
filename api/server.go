package api

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"example.com/myapp/api/handler"
)

func newMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/version", handler.HandleVersion)
	mux.HandleFunc("/users", handler.HandleUsers)
	mux.HandleFunc("/groups", handler.HandleGroups)
	return mux
}

func StartServer(ctx context.Context) {
	port := 8081
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: newMux(),
	}
	go func() {
		slog.Info(fmt.Sprintf("started http server => localhost:%d", port))
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			slog.Error("got an error providing http service", "err", err)
		}
		slog.Info("stopped http server normally")
	}()

	<-ctx.Done()
	slog.Info("caught the stop signal")
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctxTimeout); err != nil {
		slog.Error("failed to shutdown normally", "err", err)
		return
	}
	slog.Info("done graceful shutdown")
}
