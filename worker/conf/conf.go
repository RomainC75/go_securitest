package conffff

import "shared/config"

var VarList = []config.ConfigVar{
	"SERVER_PORT",
	"SERVER_JWT_SECRET",

	"KAFKA_TOPIC",
	"KAFKA_URL",
	"KAFKA_CONSUMER_GROUP_ID",
}
