package routes

import (
	"net/http"
	controllers "server/api/controllers"
)

func UserRoutes(mux *http.ServeMux) {
	authController := controllers.NewAuthCtrl()
	mux.HandleFunc("POST /auth/signup", authController.HandleSignupUser)
}
