package routes

import (
	"net/http"
	"server/internal/api/controllers"
	"server/internal/api/middlewares"
)

func AuthRoutes(mux *http.ServeMux) {
	controllers := controllers.NewAuthController()
	mux.HandleFunc("POST /auth/signup", controllers.HandleAuthSignup)
	mux.HandleFunc("POST /auth/login", controllers.HandleAuthLogin)
	mux.Handle("POST /auth/whoami",
		middlewares.AuthMiddleware(
			http.HandlerFunc(controllers.HandleWhoAmI),
		),
	)
}
