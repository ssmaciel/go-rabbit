package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
	"github.com/wesleywillians/go-rabbitmq/queue"
)


func init() {
	log.Println("Init consumer")
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func main() {
	messageChannel := make(chan amqp.Delivery)

	rabbitMQ := queue.NewRabbitMQ()
	ch := rabbitMQ.Connect()
	defer ch.Close()

	rabbitMQ.Consume(messageChannel)

	for msg := range messageChannel {
		process(msg)
	}
}

func process(msg amqp.Delivery) {
	log.Println(string(msg.Body))
}
