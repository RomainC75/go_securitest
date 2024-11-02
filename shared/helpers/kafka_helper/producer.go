package kafka_helper

import (
	"fmt"
	"os"
	"os/signal"
	"shared/config"
	"syscall"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/sirupsen/logrus"
	"golang.org/x/exp/rand"
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
	}
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

func (kh *KafkaHandler) Listen() {
	// Ctrl-C, etc
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	run := true
	for run {
		select {
		case sig := <-sigchan:
			fmt.Printf("Caught signal %v: terminating\n", sig)
			run = false
		default:
			ev, err := kh.c.ReadMessage(100 * time.Millisecond)
			if err != nil {
				// Errors are informational and automatically handled by the consumer
				continue
			}
			fmt.Printf("Consumed event from topic %s: key = %-10s value = %s\n",
				*ev.TopicPartition.Topic, string(ev.Key), string(ev.Value))
		}
	}

	kh.c.Close()
}

func (kh *KafkaHandler) Produce() {

	go func() {
		for e := range kh.p.Events() {
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

	users := [...]string{"eabara", "jsmith", "sgarcia", "jbernard", "htanaka", "awalther"}
	items := [...]string{"book", "alarm clock", "t-shirts", "gift card", "batteries"}

	for n := 0; n < 10; n++ {
		key := users[rand.Intn(len(users))]
		data := items[rand.Intn(len(items))]
		kh.p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &kh.t, Partition: kafka.PartitionAny},
			Key:            []byte(key),
			Value:          []byte(data),
		}, nil)
	}

	// Wait for all messages to be delivered
	kh.p.Flush(15 * 1000)
	kh.p.Close()
}
