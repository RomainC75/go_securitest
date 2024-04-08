package routes

import (
	"net/http"
	controller "server/api/controllers"
)

func UserRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /user/signup", controller.HandleSignupUser)
}
