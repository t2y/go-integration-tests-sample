package handler

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"net/http"
)

func HandleUsers(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	m := []map[string]any{
		{
			"name":    "john",
			"address": "kobe",
		},
		{
			"name":    "bob",
			"address": "sanda",
		},
	}

	b, err := json.Marshal(m)

	if rand.N(3) == 1 {
		err = fmt.Errorf("random error")
	}
	if err != nil {
		writeErrorMessage(w, err)
		return
	}
	writeResponse(w, b)
}
