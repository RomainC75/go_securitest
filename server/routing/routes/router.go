package routes

import "net/http"

func RegisterRoutes(mux *http.ServeMux) {
	PingRoutes(mux)
	UserRoutes(mux)
	WorkRoutes(mux)
}
