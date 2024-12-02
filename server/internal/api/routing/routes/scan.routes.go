package routes

import (
	"net/http"
	"server/internal/api/cache"
	"server/internal/api/controllers"
	"server/internal/api/middlewares"
)

func AnalyseRoutes(mux *http.ServeMux) {
	controllers := controllers.NewScanCtrl()

	mux.Handle("GET /scan",
		middlewares.AuthMiddleware(
			cache.NewCache(
				http.HandlerFunc(controllers.HandleGetScan),
			),
		),
	)
	mux.Handle("POST /scan/{scenario}",
		middlewares.AuthMiddleware(
			http.HandlerFunc(controllers.HandleScan),
		),
	)
}
