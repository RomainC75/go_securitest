package bootstrap

import (
	"fmt"
	"server/conf"
	db "server/db/sqlc"
	"server/internal/api"
	"server/internal/api/routing"
	"shared/config"
)

func Bootstrap() {
	fmt.Println("==BOOTSTRAP==")
	config.Set(conf.VarList)
	// cfg := config.GetConfig()
	// utils.PrettyDisplay(".env", cfg)

	db.Connect()
	// kafka.SetKafkaWriter()

	mux := routing.ConnectRoutes()
	api.RunApi(mux)
}
