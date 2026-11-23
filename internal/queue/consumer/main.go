package main

import (
	"log"
	
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	_, err = ch.QueueDeclare(
		"tasks", // queue name
		true,  // durable
		false, // autoDelete
		false, // exclusive
		false, // noWait
		nil,   // args
	)

	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	messages, err :=ch.Consume(
		"tasks",
		"", // consumer name
		false, // auto acknowledge
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	log.Println("Waiting for messages...")

	for msg := range messages {
		log.Printf("Received a message: %s", string(msg.Body))
		msg.Ack(false)
	}
}