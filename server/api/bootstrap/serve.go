package bootstrap

import (
	db "server/db/sqlc"
	"server/routing"
	"shared/config"
	"shared/events"
)

func Serve() {
	config.Set()
	db.Connect()

	// fmt.Println("=> serve", config.Server.Port)

	events.InitTaskDistributor()
	go events.InitTaskProcessor(false)

	routing.Init()
	routing.RegisterRoutes()
	routing.Serve()
}
