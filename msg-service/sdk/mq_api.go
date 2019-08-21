ExchangeHandlerpackage sdk

import (
"github.com/streadway/amqp"
"log"
"net/http"
"io/ioutil"
"encoding/json"
"fmt"
)

func (r *RabbitMQ) Connect() (err error) {
	r.conn, err = amqp.Dial(*AmqpUri)
	if err != nil {
		log.Printf("[amqp] connect error: %s\n", err)
		return err
	}
	r.channel, err = r.conn.Channel()
	if err != nil {
		log.Printf("[amqp] get channel error: %s\n", err)
		return err
	}
	r.done = make(chan error)
	return nil
}

//Connect with parameter
func (r *RabbitMQ) ConnectInfo(connUrl string) (err error) {
	r.conn, err = amqp.Dial(connUrl)
	if err != nil {
		log.Printf("[amqp] connect error: %s\n", err)
		return err
	}
	r.channel, err = r.conn.Channel()
	if err != nil {
		log.Printf("[amqp] get channel error: %s\n", err)
		return err
	}
	r.done = make(chan error)
	return nil
}

/*
prefetchCount：会告诉RabbitMQ不要同时给一个消费者推送多于N个消息，即一旦有N个消息还没有ack，则该consumer将block掉，直到有消息ack
global：true\false 是否将上面设置应用于channel，简单点说，就是上面限制是channel级别的还是consumer级别
*/
func (r *RabbitMQ) AutoDisableQos() (err error){
	err = r.channel.Qos(10, 0, false)
	return err
}

/*
deliveryTag:该消息的index
multiple：是否批量.true:将一次性ack所有小于deliveryTag的消息。
*/
func (r *RabbitMQ) BasicAck(deliveryTag uint64, multiple bool) (err error) {
	err = r.channel.Ack(deliveryTag, multiple)
	return err
}

/*
deliveryTag:该消息的index
multiple：是否批量.true:将一次性拒绝所有小于deliveryTag的消息。
requeue：被拒绝的是否重新入队列
*/
func (r *RabbitMQ) BasicNack(deliveryTag uint64, multiple, requeue bool) (err error) {
	err = r.channel.Nack(deliveryTag, multiple, requeue)
	return err
}

/*
deliveryTag:该消息的index
requeue：被拒绝的是否重新入队列
channel.basicNack 与 channel.basicReject 的区别:
	basicNack可以拒绝多条消息
	basicReject一次只能拒绝一条消息
*/
func (r *RabbitMQ) BasicReject(deliveryTag uint64, requeue bool) (err error) {
	err = r.channel.Reject(deliveryTag, requeue)
	return err
}

/*autoAck：是否自动ack，如果不自动ack，需要使用channel.ack、channel.nack、channel.basicReject 进行消息应答*/

func (r *RabbitMQ) Publish(exchange, key string, deliverymode, priority uint8, body string) (err error) {
	err = r.channel.Publish(exchange, key, false, false,
		amqp.Publishing{
			//Headers:         amqp.Table{},
			ContentType:     "text/plain",
			//ContentEncoding: "",
			DeliveryMode:    deliverymode,
			Priority:        priority,
			Body:            []byte(body),
		},
	)
	if err != nil {
		log.Printf("[amqp] publish message error: %s\n", err)
		return err
	}
	return nil
}

/*
type：有direct、fanout、topic三种
durable：true、false true：在服务器重启时，能够存活
exclusive ： 是否为当前连接的专用队列，在连接断开后，会自动删除该队列，生产环境中应该很少用到吧。
autodelete： true:当没有任何消费者使用时，自动删除该队列。
*/
func (r *RabbitMQ) DeclareExchange(name, typ string, durable, autodelete, nowait bool) (err error) {
	err = r.channel.ExchangeDeclare(name, typ, durable, autodelete, false, nowait, nil)
	if err != nil {
		log.Printf("[amqp] declare exchange error: %s\n", err)
		return err
	}
	return nil
}

func (r *RabbitMQ) DeleteExchange(name string) (err error) {
	err = r.channel.ExchangeDelete(name, false, false)
	if err != nil {
		log.Printf("[amqp] delete exchange error: %s\n", err)
		return err
	}
	return nil
}

func (r *RabbitMQ) DeclareQueue(name string, durable, autodelete, exclusive, nowait bool) (err error) {
	_, err = r.channel.QueueDeclare(name, durable, autodelete, exclusive, nowait, nil)
	if err != nil {
		log.Printf("[amqp] declare queue error: %s\n", err)
		return err
	}
	return nil
}

func (r *RabbitMQ) DeleteQueue(name string) (err error) {
	// TODO: other property wrapper
	_, err = r.channel.QueueDelete(name, false, false, false)
	if err != nil {
		log.Printf("[amqp] delete queue error: %s\n", err)
		return err
	}
	return nil
}

/*
channel.queueBind(queueName, EXCHANGE_NAME, bindingKey);
用于通过绑定bindingKey将queue到Exchange，之后便可以进行消息接收
*/
func (r *RabbitMQ) BindQueue(queue, exchange string, keys []string, nowait bool) (err error) {
	for _, key := range keys {
		if err = r.channel.QueueBind(queue, key, exchange, nowait, nil); err != nil {
			log.Printf("[amqp] bind queue error: %s\n", err)
			return err
		}
	}
	return nil
}

func (r *RabbitMQ) UnBindQueue(queue, exchange string, keys []string) (err error) {
	for _, key := range keys {
		if err = r.channel.QueueUnbind(queue, key, exchange, nil); err != nil {
			log.Printf("[amqp] unbind queue error: %s\n", err)
			return err
		}
	}
	return nil
}

/*autoAck：是否自动ack，如果不自动ack，需要使用channel.ack、channel.nack、channel.basicReject 进行消息应答*/
func (r *RabbitMQ) ConsumeQueue(queue string, message chan []byte) (err error) {
	deliveries, err := r.channel.Consume(queue, "", true, false, false, false, nil)
	if err != nil {
		log.Printf("[amqp] consume queue error: %s\n", err)
		return err
	}
	go func(deliveries <-chan amqp.Delivery, done chan error, message chan []byte) {
		for d := range deliveries {
			message <- d.Body
		}
		done <- nil
	}(deliveries, r.done, message)
	return nil
}

func (r *RabbitMQ) Close() (err error) {
	err = r.conn.Close()
	if err != nil {
		log.Printf("[amqp] close error: %s\n", err)
		return err
	}
	return nil
}

// HTTP Handlers
func QueueHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" || r.Method == "DELETE" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		entity := new(QueueEntity)
		if err = json.Unmarshal(body, entity); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		rabbit := new(RabbitMQ)
		if err = rabbit.Connect(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rabbit.Close()

		if r.Method == "POST" {
			if err = rabbit.DeclareQueue(entity.Name, entity.Durable, entity.AutoDelete, entity.Exclusive, entity.NoWait); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Write([]byte("declare queue ok"))
		} else if r.Method == "DELETE" {
			if err = rabbit.DeleteQueue(entity.Name); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Write([]byte("delete queue ok"))
		}
	} else if r.Method == "GET" {
		r.ParseForm()

		rabbit := new(RabbitMQ)
		if err := rabbit.Connect(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rabbit.Close()

		message := make(chan []byte)

		for _, name := range r.Form["name"] {
			if err := rabbit.ConsumeQueue(name, message); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		w.Write([]byte(""))
		w.(http.Flusher).Flush()

		for {
			fmt.Fprintf(w, "%s\n", <-message)
			w.(http.Flusher).Flush()
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func QueueBindHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" || r.Method == "DELETE" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		entity := new(QueueBindEntity)
		if err = json.Unmarshal(body, entity); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		rabbit := new(RabbitMQ)
		if err = rabbit.Connect(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rabbit.Close()

		if r.Method == "POST" {
			if err = rabbit.BindQueue(entity.Queue, entity.Exchange, entity.Keys, entity.NoWait); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Write([]byte("bind queue ok"))
		} else if r.Method == "DELETE" {
			if err = rabbit.UnBindQueue(entity.Queue, entity.Exchange, entity.Keys); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Write([]byte("unbind queue ok"))
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func PublishHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		entity := new(MessageEntity)
		if err = json.Unmarshal(body, entity); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		rabbit := new(RabbitMQ)
		if err = rabbit.Connect(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rabbit.Close()

		if err = rabbit.Publish(entity.Exchange, entity.Key, entity.DeliveryMode, entity.Priority, entity.Body); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte("publish message ok"))
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func ExchangeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" || r.Method == "DELETE" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		entity := new(ExchangeEntity)
		if err = json.Unmarshal(body, entity); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		rabbit := new(RabbitMQ)
		if err = rabbit.Connect(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rabbit.Close()

		if r.Method == "POST" {
			if err = rabbit.DeclareExchange(entity.Name, entity.Type, entity.Durable, entity.AutoDelete, entity.NoWait); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Write([]byte("declare exchange ok"))
		} else if r.Method == "DELETE" {
			if err = rabbit.DeleteExchange(entity.Name); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Write([]byte("delete exchange ok"))
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

