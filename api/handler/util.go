package handler

import (
	"fmt"
	"log/slog"
	"net/http"
)

func writeErrorMessage(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	message := fmt.Sprintf(`{"message":"%s"}`, err.Error())
	if _, err := w.Write([]byte(message)); err != nil {
		slog.Error("failed to write error", "err", err)
		return
	}
	w.Write([]byte("\n"))
}

func writeResponse(w http.ResponseWriter, b []byte) {
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		slog.Error("failed to write response", "err", err)
		return
	}
	w.Write([]byte("\n"))
}
