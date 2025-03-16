package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/segmentio/kafka-go"
)

func main() {
	// Create a reader for consuming messages from Kafka
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:9092"}, // Kafka broker addresses
		Topic:     "my-topic",                 // Topic to consume messages from
		Partition: 0,                          // Specify a specific partition (remove this line to read from all partitions)
		MinBytes:  10e3,                       // 10KB minimum batch size
		MaxBytes:  10e6,                       // 10MB maximum batch size
	})

	// Ensure the reader is closed when the program exits
	defer func() {
		if err := r.Close(); err != nil {
			log.Printf("failed to close reader: %v", err)
		}
	}()

	// Channel to handle termination signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Context for graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start a goroutine to handle termination signals
	go func() {
		<-sigChan
		log.Println("Received termination signal, shutting down...")
		cancel() // Cancel the context to stop reading
	}()

	// Consume messages
	for {
		select {
		case <-ctx.Done():
			log.Println("Stopping reader...")
			return
		default:
			m, err := r.ReadMessage(ctx)
			if err != nil {
				log.Printf("failed to read message: %v", err)
				continue // Continue reading after an error
			}
			fmt.Printf("message at offset %d: %s\n", m.Offset, string(m.Value))
		}
	}
}
