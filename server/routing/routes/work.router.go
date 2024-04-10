package routes

import (
	"net/http"
	controllers "server/api/controllers"
)

func WorkRoutes(mux *http.ServeMux) {
	workController := controllers.NewWorkCtrl()
	mux.HandleFunc("POST /work", workController.HandleWorkTest)
}
