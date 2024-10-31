package conf

import "shared/config"

var VarList = []config.ConfigVar{
	"POSTGRES_USER",
	"POSTGRES_PASSWORD",
	"POSTGRES_DB_NAME",
	"POSTGRES_HOST",
	"POSTGRES_PORT",

	"SERVER_PORT",
	"SERVER_JWT_SECRET",
}
