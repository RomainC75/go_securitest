package ctrl_utils

import (
	"encoding/json"
	"net/http"
)

type CtrlResponse map[string]any

type CustomError struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

func SendJsonResponse(w http.ResponseWriter, status uint, response map[string]any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(status))
	json.NewEncoder(w).Encode(response)
}

func SendErrorResponse(w http.ResponseWriter, statusCode int, errorString string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	encodingError := json.NewEncoder(w).Encode(CustomError{
		Message: "error",
		Error:   errorString,
	})
	if encodingError != nil {
		http.Error(w, encodingError.Error(), http.StatusInternalServerError)
	}
}
