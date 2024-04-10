package ctrl_utils

import (
	"encoding/json"
	"net/http"
)

type CtrlResponse map[string]any

func SendJsonResponse(w http.ResponseWriter, status uint, response map[string]any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(status))
	json.NewEncoder(w).Encode(response)
}
