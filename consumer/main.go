package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main() {
	connection, err := amqp.Dial("amqp://renas:root@localhost:5671/")
	if err != nil {
		panic(err)
	}

	defer connection.Close()

	fmt.Println("Connected")

	channel, err := connection.Channel()
	if err != nil {
		panic(err)
	}

	msgs, err := channel.Consume(
		"first-queue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println("Error msgs")
		panic(err)
	}

	counter := 1
	forever := make(chan bool)
	go func() {
		for msg := range msgs {
			fmt.Printf("Received Message for the %d time: %s\n", counter, msg.Body)
			counter++
		}
	}()

	fmt.Println("Waiting for messages...")
	<-forever
}
