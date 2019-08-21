package main

import (
	"net/http"
	"log"

	"bx.com/msg-service/sdk"
)

func main() {
	//Start Consumer Service
	/*
	go func(){
		service.RegistServiceHandle("rabbit_service", "EmailQueue", []string{"Email", "EMAIL", "email"}, new(service.EmailService))
	}()

	go func(){
		service.RegistServiceHandle("rabbit_service", "PhoneQueue", []string{"Phone", "PHONE", "phone"}, new(service.PhoneService))
	}()
	*/

	// Register HTTP Handlers
	http.HandleFunc("/exchange", sdk.ExchangeHandler)
	http.HandleFunc("/queue/bind", sdk.QueueBindHandler)
	http.HandleFunc("/queue", sdk.QueueHandler)
	http.HandleFunc("/publish", sdk.PublishHandler)

	// Start HTTP Server
	log.Printf("server run %s (listen %s)\n", *sdk.Address, *sdk.AmqpUri)
	err := http.ListenAndServe(*sdk.Address, nil)
	if err != nil {
		log.Fatal(err)
	}
}
