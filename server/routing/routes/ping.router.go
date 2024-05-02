package routes

import (
	"net/http"
	controller "server/api/controllers"
	mid "server/api/middlewares"
)

func PingRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /hello", controller.HandleGetPing)
	mux.Handle("GET /whoami", mid.IsAuth(http.HandlerFunc(controller.HandleWhoAmI), false))
}
