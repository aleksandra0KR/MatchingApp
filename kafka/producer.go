package kafka

import (
	"github.com/IBM/sarama"
	"log"
	"sync"
)

type Producer struct {
	responseChannels map[string]chan *sarama.ConsumerMessage
	mu               sync.Mutex
	producer         sarama.SyncProducer
}

func (p *Producer) NewProducer() *Producer {

	producer, err := sarama.NewSyncProducer([]string{"kafka:8085"}, nil)
	if err != nil {
		log.Fatalf("Failed to create producer: %v", err)
	}

	pr := Producer{
		responseChannels: make(map[string]chan *sarama.ConsumerMessage),
		mu:               sync.Mutex{},
		producer:         producer,
	}
	return &pr
}
