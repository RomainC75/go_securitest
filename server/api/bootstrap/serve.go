package bootstrap

import (
	"server/config"
	db "server/db/sqlc"
	"server/events"
	"server/routing"
)

func Serve() {
	config.Set()
	db.Connect()

	// fmt.Println("=> serve", config.Server.Port)

	events.InitTaskDistributor()
	go events.InitTaskProcessor(*db.DbStore)

	routing.Init()
	routing.RegisterRoutes()
	routing.Serve()
}
