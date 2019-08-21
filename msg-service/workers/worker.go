package workers

import (
	"bytes"
	"log"
	"strconv"

	"bx.com/msg-service/sdk"
)

var workHandlerMapping = make(map[string]interface{})

type Handler interface {
	handle(msg []byte) error
}

type WorkerConfig struct {
	name     string
	exchange string
	queue    string
	keys     []string
}

type Worker struct {
	Name        string
	Queue       string
	Keys        []string
	EventStruct interface{}
	RabbitInfo  *sdk.RabbitMQ
	handler     Handler
}

func RegisterWorker(name string, inter interface{}) {
	workHandlerMapping[name] = inter
}

func GetWorkerStructMap() map[string]interface{} {
	return workHandlerMapping
}

func GetRabbitConnStr(config ConnInfo) string {
	var buffer bytes.Buffer
	buffer.WriteString("amqp://")
	buffer.WriteString(config.User)
	buffer.WriteString(":")
	buffer.WriteString(config.Password)
	buffer.WriteString("@")
	buffer.WriteString(config.Host)
	buffer.WriteString(":")
	buffer.WriteString(strconv.Itoa(config.Port))
	buffer.WriteString("/")

	return buffer.String()
}

func NewWorker(connConfig ConnInfo, workConfig WorkerInfo, handler Handler) (Worker, error) {
	connStr := GetRabbitConnStr(connConfig)
	rabbit := new(sdk.RabbitMQ)
	if err := rabbit.ConnectInfo(connStr); err != nil {
		return Worker{}, err
	}

	work := Worker{
		Name:       workConfig.Name,
		Queue:      workConfig.Queue,
		Keys:       workConfig.Keys,
		RabbitInfo: rabbit,
		handler:    handler,
	}
	return work, nil
}

func InitWorker(rq *sdk.RabbitMQ, exchange, queue string, keys []string) (err error) {
	ex := sdk.ExchangeEntity{
		Name:       exchange,
		Type:       "topic",
		Durable:    true,
		AutoDelete: false,
		NoWait:     false,
	}
	err = rq.DeclareExchange(ex.Name, "topic", ex.Durable, ex.AutoDelete, ex.NoWait)
	if err != nil {
		return err
	}

	que := sdk.QueueEntity{
		Name:       queue,
		Durable:    true,
		AutoDelete: false,
		Exclusive:  false,
		NoWait:     false,
	}
	err = rq.DeclareQueue(que.Name, que.Durable, que.AutoDelete, que.Exclusive, que.NoWait)
	if err != nil {
		return err
	}

	qb := sdk.QueueBindEntity{
		Queue:    queue,
		Exchange: exchange,
		Keys:     keys,
		NoWait:   false,
	}
	err = rq.BindQueue(qb.Queue, qb.Exchange, qb.Keys, qb.NoWait)
	if err != nil {
		return err
	}

	return nil
}

func (w *Worker) Run() {
	//Init Exchange Info
	err := InitWorker(w.RabbitInfo, w.Name, w.Queue, w.Keys)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	//Regist Consumer Operation
	message := make(chan []byte)
	if err := w.RabbitInfo.ConsumeQueue(w.Queue, message); err != nil {
		log.Fatal(err)
		panic(err)
	}
	for {
		select {
		case info := <-message:
			log.Printf("New queue message: %s", string(info))
			if err := w.handler.handle(info); err != nil {
				log.Printf("Handle Error: %v", err.Error())
			}
		default:
		}
	}
}

func (w *Worker) Stop() {
	err := w.RabbitInfo.Close()
	if err != nil {
		log.Printf("Stop consumer error: %s", err)
	} else {
		log.Printf("Stop consumer success.")
	}
}
