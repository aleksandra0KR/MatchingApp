package kafka

import (
	"MatchingApp/internal/model/kafka"
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	"log"
)

func SetUpProducer() *sarama.SyncProducer {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 10
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"},
		config)
	if err != nil {
		log.Fatalf("failed to setup producer: %v", err)
	}
	return &producer
}

func SendMessage(producer *sarama.SyncProducer, topic string, message kafka.Message) {
	msg := &sarama.ProducerMessage{}
	msg.Topic = "input_topic"
	bytes, _ := json.Marshal(message)
	msg.Value = sarama.ByteEncoder(bytes)

	partition, offset, err := (*producer).SendMessage(msg)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Message is stored in partition %d at offset %d\n", partition, offset)
}
