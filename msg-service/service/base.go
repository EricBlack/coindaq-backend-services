package service

import (
	"bx.com/msg-service/sdk"
	"time"
	"math/rand"
	"log"
)

type Service interface {
	RegisterConsumeEvent(body string)
	//UninstallConsumeEvent(exchange string)
}

var serviceMapping = make(map[string]interface{})

func GetServiceMapping() map[string]interface{}{
	return serviceMapping
}

func RegisterServiceMapping(name string, inter interface{}){
	serviceMapping[name] = inter
}

func GetRandomString () string{
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 6; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func InitServiceEvent(rq *sdk.RabbitMQ, exchange, queue string, keys []string) (err error){
	ex := sdk.ExchangeEntity{
		Name:		exchange,
		Type:		"topic",
		Durable:	true,
		AutoDelete:	false,
		NoWait:		false,
	}
	err = rq.DeclareExchange(ex.Name, "topic", ex.Durable, ex.AutoDelete, ex.NoWait)
	if err != nil{
		return err
	}

	que := sdk.QueueEntity{
		Name:		queue,
		Durable:	true,
		AutoDelete:	false,
		Exclusive:	false,
		NoWait:		false,
	}
	err = rq.DeclareQueue(que.Name, que.Durable, que.AutoDelete, que.Exclusive, que.NoWait)
	if err != nil{
		return err
	}

	qb := sdk.QueueBindEntity{
		Queue:		queue,
		Exchange:	exchange,
		Keys:		keys,
		NoWait:		false,
	}
	err = rq.BindQueue(qb.Queue, qb.Exchange, qb.Keys, qb.NoWait)
	if err != nil{
		return err
	}

	return nil
}

func RegistServiceHandle(exchange, queue string, keys []string, service Service) (err error) {
	queue = queue +"_" +GetRandomString()
	log.Printf("Register to queue=%s", queue)
	message := make(chan []byte)

	rq := new(sdk.RabbitMQ)
	err = rq.Connect()
	if err != nil{
		return err
	}
	//Regist New Service

	err = InitServiceEvent(rq, exchange, queue, keys)
	if err != nil {
		return err
	}

	//Regist Consumer Operation
	err = rq.ConsumeQueue(queue, message)
	if err != nil {
		return err
	}

	for {
		select {
		case info := <-message:
			log.Printf("New queue message: %s", string(info))
			go service.RegisterConsumeEvent(string(info))
		default:
		}
	}

	log.Printf("Complete register one consumer.")
	return err
}