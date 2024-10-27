package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func PingController(w http.ResponseWriter, r *http.Request) {
	status := 200
	fmt.Println("hello ! ")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(status))
	json.NewEncoder(w).Encode(map[string]any{
		"mesage": "wooo",
	})
}
