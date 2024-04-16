package bootstrap

import (
	"server/routing"
	"shared/config"
	db "shared/db/sqlc"
	"shared/events"
)

func Serve() {
	config.Set()
	db.Connect()

	// fmt.Println("=> serve", config.Server.Port)

	events.InitTaskDistributor()
	go events.InitTaskProcessor(*db.DbStore, false)

	routing.Init()
	routing.RegisterRoutes()
	routing.Serve()
}
