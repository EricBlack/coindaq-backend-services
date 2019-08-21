package sdk

import "log"

func InitRabbitMQ() (RabbitMQ, error) {
	rq := RabbitMQ{
		conn:    nil,
		channel: nil,
		done:    make(chan error),
	}
	err := rq.Connect()
	if err != nil {
		return RabbitMQ{}, err
	}
	return rq, nil
}

func (r *RabbitMQ) InitExchange() (err error) {
	exchange := ExchangeEntity{
		Name:       "rabbit_demo",
		Type:       "topic",
		Durable:    true,
		AutoDelete: false,
		NoWait:     false,
	}
	err = r.DeclareExchange(exchange.Name, exchange.Type, exchange.Durable, exchange.AutoDelete, exchange.NoWait)
	return err
}

func (r *RabbitMQ) InitQueue() (err error) {
	var queueList []QueueEntity
	queue1 := QueueEntity{
		Name:       "rabbit_emailAuth",
		Durable:    true,
		AutoDelete: false,
		Exclusive:  false,
		NoWait:     false,
	}
	queueList = append(queueList, queue1)

	queue2 := QueueEntity{
		Name:       "rabbit_phoneAuth",
		Durable:    true,
		AutoDelete: false,
		Exclusive:  false,
		NoWait:     false,
	}
	queueList = append(queueList, queue2)

	queue3 := QueueEntity{
		Name:       "rabbit_googleAuth",
		Durable:    true,
		AutoDelete: false,
		Exclusive:  false,
		NoWait:     false,
	}
	queueList = append(queueList, queue3)
	for _, queue := range queueList {
		err = r.DeclareQueue(queue.Name, queue.Durable, queue.AutoDelete, queue.Exclusive, queue.NoWait)
		if err != nil {
			return err
		}
	}
	return err
}

func (r *RabbitMQ) InitBinding() (err error) {
	err = r.BindQueue("rabbit_emailAuth", "rabbit_demo", []string{"email", "Email", "EMAIL"}, false)
	if err != nil {
		return err
	}

	err = r.BindQueue("rabbit_phoneAuth", "rabbit_demo", []string{"phone", "Phone", "PHONE"}, false)
	if err != nil {
		return err
	}

	err = r.BindQueue("rabbit_googleAuth", "rabbit_demo", []string{"google", "Google", "GOOGLE"}, false)
	return err
}

func (r *RabbitMQ) RegisterConsumer(username, queue string) (err error) {
	log.Printf("Register user=%s to queue=%s", username, queue)
	message := make(chan []byte)
	err = r.ConsumeQueue(queue, message)

	for {
		select {
		case info := <-message:
			log.Printf("[%s]: %s", username, string(info))
		default:
		}
	}

	log.Printf("Complete register one consumer.")
	return err
}

func (r *RabbitMQ) ShutdownConsumer() (err error) {
	return err
}

func (r *RabbitMQ) ProduceInfo(key, message string) (err error) {
	err = r.Publish("rabbit_demo", key, 1, 1, message)
	log.Printf("Publish: key=%s  message:=%s", key, message)
	return err
}
