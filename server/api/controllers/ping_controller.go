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

func HandleWhoAmI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(w).Encode("ping")

	response := struct {
		Email  string
		UserId int32
	}{
		Email:  r.Context().Value("user_email").(string),
		UserId: r.Context().Value("user_id").(int32),
	}
	ctrl_utils.SendJsonResponse(w, http.StatusCreated, ctrl_utils.CtrlResponse{"message": response})

	if err != nil {
		ctrl_utils.SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
}
