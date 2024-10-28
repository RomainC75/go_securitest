package utils

import (
	"encoding/json"
	"net/http"
)

func SendJson(w http.ResponseWriter, status int, content map[string]any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(content)
}
