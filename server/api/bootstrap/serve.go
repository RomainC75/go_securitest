package bootstrap

import (
	"server/config"
	db "server/db/sqlc"
	"server/routing"
	"server/worker"
)

func Serve() {
	config.Set()
	db.Connect()

	// fmt.Println("=> serve", config.Server.Port)

	worker.InitTaskDistributor()
	go worker.InitTaskProcessor(*db.DbStore)

	routing.Init()
	routing.RegisterRoutes()
	routing.Serve()
}
