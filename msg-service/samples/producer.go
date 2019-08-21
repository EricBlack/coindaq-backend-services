package main

import (
	"github.com/streadway/amqp"
	"log"
	"os"
	"strings"
)

func failOnError1(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError1(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError1(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"msg_topic", // name
		"topic",      // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	failOnError1(err, "Failed to declare an exchange")

	err = ch.Publish(
		"msg_topic",          // exchange
		"email", // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Email info test"),
		})
	failOnError1(err, "Failed to publish a message")

	log.Printf(" [x] Sent %s", "Email info test")

	err = ch.Publish(
		"msg_topic",          // exchange
		"phone", // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Phone info test"),
		})
	failOnError1(err, "Failed to publish a message")

	log.Printf(" [x] Sent %s", "Phone info test")

	err = ch.Publish(
		"msg_topic",          // exchange
		"google", // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Google info test"),
		})
	failOnError1(err, "Failed to publish a message")

	log.Printf(" [x] Sent %s", "Google info test")
}

func bodyFrom(args []string) string {
	var s string
	if (len(args) < 3) || os.Args[2] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[2:], " ")
	}
	return s
}

func severityFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "anonymous.info"
	} else {
		s = os.Args[1]
	}
	return s
}

