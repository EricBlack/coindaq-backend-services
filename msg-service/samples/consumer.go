package main

import (
	"github.com/streadway/amqp"
	"log"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"msg_topic", // name
		"topic",      // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		true,        // no-wait
		nil,          // arguments
	)
	failOnError(err, "Failed to declare an exchange")


	q1, err := ch.QueueDeclare(
		"test1",    // name
		false, // durable
		false, // delete when unused
		false,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue1")

	q2, err := ch.QueueDeclare(
		"test2",    // name
		false, // durable
		false, // delete when unused
		false,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue2")

	bind1 :=[]string{"email", "phone"}
	for _, s := range bind1 {
		log.Printf("Binding queue %s to exchange %s with routing key %s", q1.Name, "msg_topic", s)
		err = ch.QueueBind(
			q1.Name,       // queue name
			s,            // routing key
			"msg_topic", // exchange
			false,
			nil)
		failOnError(err, "Failed to bind a queue")
	}
	bind2 :=[]string{"google"}
	for _, s := range bind2 {
		log.Printf("Binding queue %s to exchange %s with routing key %s", q2.Name, "msg_topic", s)
		err = ch.QueueBind(
			q2.Name,       // queue name
			s,            // routing key
			"msg_topic", // exchange
			false,
			nil)
		failOnError(err, "Failed to bind a queue")
	}

	msgs1, err := ch.Consume(
		q1.Name, // queue
		"consumer1",     // consumer
		true,   // auto ack
		false,  // exclusive
		false,  // no local
		false,  // no wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	msgs2, err := ch.Consume(
		q2.Name, // queue
		"consumer2",     // consumer
		true,   // auto ack
		false,  // exclusive
		false,  // no local
		false,  // no wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs1 {
			log.Printf(" [consumer1]: %s", d.Body)
		}
	}()

	go func() {
		for d := range msgs2 {
			log.Printf(" [consumer2]: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}

