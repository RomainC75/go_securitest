package main

import (
	// "conf"
	"fmt"
	"os"
	"os/signal"
	"shared/config"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var VarList = []config.ConfigVar{
	"SERVER_PORT",
	"SERVER_JWT_SECRET",

	"KAFKA_TOPIC",
	"KAFKA_URL",
	"KAFKA_CONSUMER_GROUP_ID",
}

func main() {

	config.Set(VarList)

	configuration := kafka.ConfigMap{
		"bootstrap.servers": viper.GetString(string(config.KAFKA_URL)),
	}
	configuration["group.id"] = viper.GetString(string(config.KAFKA_CONSUMER_GROUP_ID))
	configuration["auto.offset.reset"] = "earliest"

	c, err := kafka.NewConsumer(&configuration)

	if err != nil {
		fmt.Printf("Failed to create consumer: %s", err)
		os.Exit(1)
	}

	topic := viper.GetString(string(config.KAFKA_TOPIC))
	err = c.SubscribeTopics([]string{topic}, nil)

	// Ctrl+C
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	// Process messages
	run := true
	for run {
		select {
		case sig := <-sigchan:
			logrus.Errorf("Caught signal %v: terminating\n", sig)
			run = false
		default:
			ev, err := c.ReadMessage(100 * time.Millisecond)
			if err != nil {
				// Errors are informational and automatically handled by the consumer
				continue
			}
			logrus.Warnf("Consumed event from topic %s: key = %-10s value = %s\n",
				*ev.TopicPartition.Topic, string(ev.Key), string(ev.Value))
		}
	}

	c.Close()
}
