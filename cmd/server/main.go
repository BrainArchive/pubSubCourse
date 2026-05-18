package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/bootdotdev/learn-pub-sub-starter/internal/pubsub"
	"github.com/bootdotdev/learn-pub-sub-starter/internal/routing"
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

	ch, err := connection.Channel()
	if err != nil {
		log.Fatalf("could not create new Channel: %v", err)
	}
	defer ch.Close()
	err = pubsub.PublishJSON(ch, routing.ExchangePerilDirect, routing.PauseKey, routing.PlayingState{IsPaused: true})
	if err != nil {
		log.Fatalf("error publishing JSON: %v", err)
	}

	<-signalChan
	fmt.Println("Rabbit Connection Closed")
}
