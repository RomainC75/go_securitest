package routes

import (
	"net/http"
	"server/internal/api/controllers"
)

func AuthRoutes(mux *http.ServeMux) {
	controllers := controllers.NewAuthController()
	mux.HandleFunc("POST /auth/signup", controllers.HandleAuthSignup)

}
