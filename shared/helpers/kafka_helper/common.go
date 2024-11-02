package kafka_helper

import (
	"fmt"
	"shared/config"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/sirupsen/logrus"
)

type KafkaHandler struct {
	p *kafka.Producer
	c *kafka.Consumer
	t string
}

func NewKafkaHandler(env map[config.ConfigVar]string) (*KafkaHandler, error) {
	newP, err := createProducer(env)
	if err != nil {
		return nil, err
	}

	newC, err := createConsumer(env)
	if err != nil {
		return nil, err
	}

	return &KafkaHandler{
		p: newP,
		c: newC,
		t: env[config.KAFKA_TOPIC],
	}, nil
}

func createProducer(env map[config.ConfigVar]string) (*kafka.Producer, error) {
	conf := kafka.ConfigMap{
		"bootstrap.servers": env[config.KAFKA_URL],
	}
	p, err := kafka.NewProducer(&conf)

	if err != nil {
		logrus.Warnf("Failed to create producer: %s", err)
		// os.Exit(1)
		return nil, err
	}

	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Failed to deliver message: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Produced event to topic %s: key = %-10s value = %s\n",
						*ev.TopicPartition.Topic, string(ev.Key), string(ev.Value))
				}
			}
		}
	}()

	return p, err
}

func createConsumer(env map[config.ConfigVar]string) (*kafka.Consumer, error) {
	conf := kafka.ConfigMap{
		"bootstrap.servers": env[config.KAFKA_URL],
	}
	conf["group.id"] = env[config.KAFKA_CONSUMER_GROUP_ID]
	conf["auto.offset.reset"] = "earliest"

	c, err := kafka.NewConsumer(&conf)

	if err != nil {
		logrus.Warnf("Failed to create consumer: %s", err)
		// os.Exit(1)
		return nil, err
	}
	topic := env[config.KAFKA_TOPIC]
	err = c.SubscribeTopics([]string{topic}, nil)
	return c, err

}
