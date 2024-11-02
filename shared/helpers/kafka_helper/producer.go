package kafka_helper

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func (kh *KafkaHandler) Push(key string, data []byte) {

	kh.p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &kh.t, Partition: kafka.PartitionAny},
		Key:            []byte(key),
		Value:          []byte(data),
	}, nil)

	// Wait for all messages to be delivered
	// kh.p.Flush(15 * 1000)
	// kh.p.Close()
}
