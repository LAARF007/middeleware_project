package services

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
)

func initStream() error {

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		return fmt.Errorf("failed to connect to NATS server: %w", err)
	}

	// getting Jetstream context
	jsc, err := nc.JetStream()
	if err != nil {
		return fmt.Errorf("failed to get JetStream context: %w", err)
	}

	// Init stream
	// Try to get existing stream first, create if it doesn't exist
	streamInfo, err := jsc.StreamInfo("USERS")
	if err != nil {
		// Stream doesn't exist, create it
		_, err = jsc.AddStream(&nats.StreamConfig{
			Name:     "USERS",
			Subjects: []string{"USERS.>"},
		})
		if err != nil {
			return fmt.Errorf("failed to create stream: %w", err)
		}
		log.Println("Stream USERS created successfully")
	} else {
		log.Printf("Stream USERS already exists with %d messages", streamInfo.State.Msgs)
	}

	return nil
}
