package handler

import (
	"encoding/json"
	"net/http"
)

func HandleVersion(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	b, err := json.Marshal(map[string]any{"version": "0.0.1"})
	if err != nil {
		writeErrorMessage(w, err)
		return
	}
	writeResponse(w, b)
}
