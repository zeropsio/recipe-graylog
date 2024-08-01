package main

import (
	"log"
	"time"

	"github.com/Graylog2/go-gelf/gelf"
)

func sendTestMessage(gelfWriter *gelf.Writer) {
	// Test message
	msg := &gelf.Message{
		Version:  "1.1",
		Host:     "localhost",
		Short:    "Test message",
		Full:     "This is a test message sent to Graylog",
		TimeUnix: float64(time.Now().UnixNano()) / float64(time.Second),
		Level:    6, // 1=Alert, 2=Critical, 3=Error, 4=Warning, 5=Notice, 6=Informational, 7=Debug
		Facility: "example",
	}

	// Send the message
	err := gelfWriter.WriteMessage(msg)
	if err != nil {
		log.Printf("Failed to send message: %v", err)
	} else {
		log.Println("Message successfully sent to Graylog")
	}
}

func main() {
	// Graylog server address
	graylogAddr := "graylog:12201"

	// Create a new GELF writer
	gelfWriter, err := gelf.NewWriter(graylogAddr)
	if err != nil {
		log.Fatalf("Failed to create GELF writer: %v", err)
	}

	// Send a test message every 5 seconds
	for {
		sendTestMessage(gelfWriter)
		time.Sleep(10 * time.Second)
	}
}
