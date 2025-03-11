package main

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

func main() {
	// Create a writer for sending messages to Kafka
	w := &kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"), // Addresses of the brokers
		Topic:    "my-topic",                  // Topic to which messages will be sent
		Balancer: &kafka.LeastBytes{},         // Balancer for partitioning
	}

	// Send a message
	err := w.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("key"),
			Value: []byte("Hello, Kafka!"),
		},
	)
	if err != nil {
		log.Fatalf("Failed to write message: %s", err)
	}

	// Close the writer
	if err := w.Close(); err != nil {
		log.Fatalf("Failed to close writer: %s", err)
	}
}
