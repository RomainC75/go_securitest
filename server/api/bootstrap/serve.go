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

	events.InitTaskDistributor()

	// pass routing // logic ??
	go events.InitTaskProcessor(false)

	routing.Init()
	routing.RegisterRoutes()
	routing.Serve()
}
