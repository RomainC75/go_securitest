package routes

import (
	"net/http"
	"server/internal/api/controllers"
	"server/internal/api/middlewares"
)

func AnalyseRoutes(mux *http.ServeMux) {
	controllers := controllers.NewScanCtrl()
	mux.Handle("POST /scan/{scenario}",
		middlewares.AuthMiddleware(
			http.HandlerFunc(controllers.HandleScan),
		),
	)
}
