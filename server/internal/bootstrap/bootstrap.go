package bootstrap

import (
	"fmt"
	"os"
	"server/conf"
	db "server/db/sqlc"
	"server/internal/api"
	validator_helper "server/internal/api/dtos/validator"
	"server/internal/api/routing"
	"server/internal/queue"
	"shared/config"
	"shared/helpers/kafka_helper"
	shared_utils "shared/utils"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Bootstrap() {
	fmt.Println("==BOOTSTRAP==")

	validator_helper.SetValidate()
	config.Set(conf.VarList)
	db.Connect()

	envMp := map[config.ConfigVar]string{
		config.KAFKA_URL:               viper.GetString(string(config.KAFKA_URL)),
		config.KAFKA_TOPIC_REQ:         viper.GetString(string(config.KAFKA_TOPIC_REQ)),
		config.KAFKA_CONSUMER_GROUP_ID: viper.GetString(string(config.KAFKA_CONSUMER_GROUP_ID)),
	}

	shared_utils.PrettyDisplay("envMp", envMp)

	kafkaEnv, err := kafka_helper.NewKafkaHandler(
		viper.GetString(string(config.KAFKA_TOPIC_REQ)),
		viper.GetString(string(config.KAFKA_TOPIC_RES)),
		envMp,
	)
	if err != nil {
		logrus.Errorf("boostrap error : ", err.Error())
		os.Exit(1)
	}

	queue.SetQueue(kafkaEnv)

	mux := routing.ConnectRoutes()
	api.RunApi(mux)
}
