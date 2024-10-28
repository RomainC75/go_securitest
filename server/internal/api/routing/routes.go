package routing

import (
	"net/http"
	"server/internal/api/routing/routes"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func ConnectRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	// prometheus
	mux.Handle("/metrics", promhttp.Handler())

	routes.OpenRoutes(mux)
	routes.AuthRoutes(mux)

	return mux

}
