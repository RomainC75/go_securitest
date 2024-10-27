package routes

import (
	"net/http"
	"server/internal/api/controllers"
)

func OpenRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/ping", controllers.PingController)
}
