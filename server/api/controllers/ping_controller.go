package controllers

import (
	"encoding/json"
	"net/http"
	"server/api/controllers/ctrl_utils"
)

func HandleGetPing(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(w).Encode("ping")
	if err != nil {
		ctrl_utils.SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
}
