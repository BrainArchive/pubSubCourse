package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	connectionString := "amqp://guest:guest@localhost:5672/"

	fmt.Println("Starting Peril client...")
	connection, err := amqp.Dial(connectionString)
	if err != nil {
		log.Fatalf("could not connect to rabbitMQ?: %v", err)
	}
	defer connection.Close()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan
	fmt.Println("Rabbit Connection Closed")
}
