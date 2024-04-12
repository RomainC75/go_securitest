package ctrl_utils

import (
	"encoding/json"
	"net/http"
)

type CustomErrorType string

const (
	NormalErrorType     CustomErrorType = "error"
	ValidationErrorType CustomErrorType = "validation_error"
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

func SendErrorResponse(w http.ResponseWriter, statusCode uint, errorString string, customMessage ...CustomErrorType) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(statusCode))

	var encodingError error
	if len(customMessage) > 0 {
		encodingError = json.NewEncoder(w).Encode(CustomError{
			Message: string(customMessage[0]),
			Error:   errorString,
		})
	} else {
		encodingError = json.NewEncoder(w).Encode(CustomError{
			Message: string(NormalErrorType),
			Error:   errorString,
		})
	}

	if encodingError != nil {
		http.Error(w, encodingError.Error(), http.StatusInternalServerError)
	}
}
