package main

import (
	// "conf"

	"fmt"
	"os"
	"shared/config"
	"shared/helpers/kafka_helper"
	"worker/queue"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var VarList = []config.ConfigVar{
	"SERVER_PORT",
	"SERVER_JWT_SECRET",

	"KAFKA_TOPIC_REQ",
	"KAFKA_URL",
	"KAFKA_CONSUMER_GROUP_ID",
}

var queueRes *queue.SQueue

func main() {
	config.Set(VarList)

	envMp := map[config.ConfigVar]string{
		config.KAFKA_URL:               viper.GetString(string(config.KAFKA_URL)),
		config.KAFKA_TOPIC_REQ:         viper.GetString(string(config.KAFKA_TOPIC_REQ)),
		config.KAFKA_TOPIC_RES:         viper.GetString(string(config.KAFKA_TOPIC_RES)),
		config.KAFKA_CONSUMER_GROUP_ID: viper.GetString(string(config.KAFKA_CONSUMER_GROUP_ID)),
	}

	kafkaEnv, err := kafka_helper.NewKafkaHandler(
		viper.GetString(string(config.KAFKA_TOPIC_RES)),
		viper.GetString(string(config.KAFKA_TOPIC_REQ)),
		envMp,
	)
	if err != nil {
		logrus.Errorf("boostrap error : ", err.Error())
		os.Exit(1)
	}

	queue.SetQueue(kafkaEnv)
	queue.GetQueue().Strategy.Listen()

	fmt.Println("fin ^^")

}
