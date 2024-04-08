package routing

import (
	"net/http"
	"server/routing/routes"
)

func Init() {
	mux = http.NewServeMux()
}

func GetRouter() *http.ServeMux {
	return mux
}

func RegisterRoutes() {
	routes.RegisterRoutes(GetRouter())
}
