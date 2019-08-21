package main

import (
	"bx.com/msg-service/sdk"
	"log"
)

func main() {
	var rq sdk.RabbitMQ
	var err error
	rq, err = sdk.InitRabbitMQ()
	if err != nil {
		log.Printf("Error: %s", err)
	}
	err = rq.InitExchange()
	err = rq.InitQueue()
	err = rq.InitBinding()

	forever := make(chan bool)
	go func(){
		err = rq.RegisterConsumer("consumer1", "rabbit_emailAuth")
		if err != nil {
			log.Printf("Error: %s", err)
			return
		}
	}()

	go func() {
		err = rq.RegisterConsumer("consumer2", "rabbit_phoneAuth")
		if err != nil {
			log.Printf("Error: %s", err)
			return
		}
	}()

	go func() {
		err = rq.RegisterConsumer("consumer3", "rabbit_googleAuth")
		if err != nil {
			log.Printf("Error: %s", err)
			return
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}
