package main

import (
	"bx.com/msg-service/sdk"
	"log"
	"os"
)

func main(){
	var rq sdk.RabbitMQ
	var err error
	rq, err = sdk.InitRabbitMQ()
	defer rq.Close()
	if err != nil{
		log.Printf("Error: %s", err)
	}
	err = rq.InitExchange()
	if err != nil{
		log.Printf("Error: %s", err)
	}

	log.Println("%s", os.Args)
	if len(os.Args) == 1{
		rq.ProduceInfo("test", "Hello Message")
	} else if len(os.Args) ==2 {
		rq.ProduceInfo(os.Args[1], "Message from" + string(os.Args[1]))
	} else {
		rq.ProduceInfo(os.Args[1], os.Args[2])
	}
}
