package handler

import (
	"encoding/json"
	"net/http"
)

func HandleGroups(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	m := []map[string]any{
		{
			"name":        "badminton",
			"description": "enjoy sports",
		},
		{
			"name":    "mountain",
			"address": "enjoy climbing",
		},
	}

	b, err := json.Marshal(m)
	if err != nil {
		writeErrorMessage(w, err)
		return
	}
	writeResponse(w, b)
}
