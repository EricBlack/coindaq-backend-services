package sdk

import (
	"flag"
	"github.com/streadway/amqp"
)

var (
	Address = flag.String("address", "127.0.0.1:8080", "Bind Web API address")
	AmqpUri = flag.String("amqp", "amqp://guest:guest@127.0.0.1:5672/", "Amqp Uri Connection info")
)

func init() {
	flag.Parse()
}

// Entity for HTTP Request Body: Message/Exchange/Queue/QueueBind JSON Input
type MessageEntity struct {
	Exchange     string `json:"exchange"`
	Key          string `json:"key"`
	DeliveryMode uint8  `json:"deliverymode"`
	Priority     uint8  `json:"priority"`
	Body         string `json:"body"`
}

/*
Queue作为真实存储消息的队列和某个Exchange绑定，具体如何路由到感兴趣的Queue则由Exchange的三种模式决定：
fanout	topic	direct
Exchange为fanout时，生产者往此Exchange发送的消息会发给每个和其绑定的Queue，此时RoutingKey并不起作用；
Exchange为topic时，生产者可以指定一个支持通配符的RoutingKey（如demo.*）发向此Exchange，凡是Exchange上RoutingKey满足此通配符的Queue就会收到消息；
Exchange为direct是最直接最简单的，生产者指定Exchange和RoutingKey，然后往其发送消息，消息只能被绑定的满足RoutingKey的Queue接受消息。
(通常如果不指定RoutingKey的具体名字，那么默认的名字其实是Queue的名字）
*/
type ExchangeEntity struct {
	Name       string `json:"name"`
	Type       string `json:"type"`
	Durable    bool   `json:"durable"`
	AutoDelete bool   `json:"autodelete"`
	NoWait     bool   `json:"nowait"`
}

type QueueEntity struct {
	Name       string `json:"name"`
	Durable    bool   `json:"durable"`
	AutoDelete bool   `json:"autodelete"`
	Exclusive  bool   `json:"exclusive"`
	NoWait     bool   `json:"nowait"`
}

type QueueBindEntity struct {
	Queue    string   `json:"queue"`
	Exchange string   `json:"exchange"`
	NoWait   bool     `json:"nowait"`
	Keys     []string `json:"keys"` // bind/routing keys
}

// RabbitMQ Operate Wrapper
type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	done    chan error
}

/*
Mandatory：
	true：如果exchange根据自身类型和消息routeKey无法找到一个符合条件的queue，那么会调用basic.return方法将消息返还给生产者。
	false：出现上述情形broker会直接将消息扔掉

Immediate：
	true：如果exchange在将消息route到queue(s)时发现对应的queue上没有消费者，那么这条消息不会放入队列中。
	当与消息routeKey关联的所有queue(一个或多个)都没有消费者时，该消息会通过basic.return方法返还给生产者。
BasicProperties ：需要注意的是BasicProperties.deliveryMode，
	0:不持久化
	1：持久化 这里指的是消息的持久化，配合channel(durable=true),queue(durable)可以实现，即使服务器宕机，消息仍然保留
简单来说：
	mandatory标志告诉服务器至少将该消息route到一个队列中，否则将消息返还给生产者；
	immediate标志告诉服务器如果该消息关联的queue上有消费者，则马上将消息投递给它，如果所有queue都没有消费者，直接把消息返还给生产者，不用将消息入队列等待消费者了。
*/

