package kafka_helper

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"golang.org/x/exp/rand"
)

func (kh *KafkaHandler) Push() {

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
