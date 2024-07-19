package kafka

import (
	"encoding/json"
	"github.com/IBM/sarama"
	"log"
)

func SetUpConsumer() *sarama.Consumer {
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, nil)
	if err != nil {
		log.Fatalf("Error creating the consumer: %v", err)
	}
	return &consumer
}

func ReadMessage(consumer *sarama.Consumer, topic string) string {
	partitionConsumer, err := (*consumer).ConsumePartition("output_topic", 0, sarama.OffsetNewest)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := partitionConsumer.Close(); err != nil {
			panic(err)
		}
	}()

	for msg := range partitionConsumer.Messages() {
		var result map[string]interface{}
		json.Unmarshal(msg.Value, &result)
		return result["result"].(string)
	}
	return ""
}
