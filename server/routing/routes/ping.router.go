package routes

import (
	"net/http"
	controller "server/api/controllers"
)

func PingRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /hello", controller.HandleGetPing)
}
