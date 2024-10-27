package bootstrap

import (
	"fmt"
	db "server/db/sqlc"
	"server/internal/api"
	"server/internal/api/routing"
	"shared/config"
)

func Bootstrap() {
	fmt.Println("==BOOTSTRAP==")
	config.Set()
	// cfg := config.GetConfig()
	// utils.PrettyDisplay(".env", cfg)

	db.Connect()
	// kafka.SetKafkaWriter()

	mux := routing.ConnectRoutes()
	api.RunApi(mux)
}
