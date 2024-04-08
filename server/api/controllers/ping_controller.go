package controller

import (
	"encoding/json"
	"net/http"
)

type CustomError struct {
	StatusCode int
	Error      string
}

func HandleGetPing(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(w).Encode("ping")
	if err != nil {
		errorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func errorResponse(w http.ResponseWriter, statusCode int, errorString string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	encodingError := json.NewEncoder(w).Encode(CustomError{
		StatusCode: statusCode,
		Error:      errorString,
	})
	if encodingError != nil {
		http.Error(w, encodingError.Error(), http.StatusInternalServerError)
	}
}
