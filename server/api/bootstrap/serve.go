package bootstrap

import (
	"server/config"
	"server/routing"
)

func Serve() {
	config.Set()
	// db.Connect()

	// fmt.Println("=> serve", config.Server.Port)

	routing.Init()
	routing.RegisterRoutes()
	routing.Serve()
}
