package routes

import (
	"net/http"
	controllers "server/api/controllers"
	mid "server/api/middlewares"
)

func WorkRoutes(mux *http.ServeMux) {
	workController := controllers.NewWorkCtrl()
	mux.Handle("POST /work/{work_code}", mid.IsAuth(http.HandlerFunc(workController.HandleWorkTest), false))
}
