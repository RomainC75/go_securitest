package controllers

import (
	"net/http"
	"server/utils"
)

func PingController(w http.ResponseWriter, r *http.Request) {
	res := map[string]any{
		"mesage": "wooo",
	}
	utils.SendJson(w, http.StatusOK, res)
}
