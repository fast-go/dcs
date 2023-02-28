package queue

import (
	"dcs/rpc/consumer/internal/config"
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

type TopicHandle interface {
	Consume([]byte)
	TopicName() string
}

type Amqp struct {
	conn         *amqp.Connection
	topicHandles []TopicHandle
}

func NewAmqp(c config.Config) *Amqp {
	//对服务进行初始化
	conn, err := amqp.Dial(c.Amqp.Host)

	if err != nil {
		fmt.Println("消息队列启动失败:", err)
	}
	return &Amqp{conn: conn}
}

func (a *Amqp) Close() {
	if !a.conn.IsClosed() {
		_ = a.conn.Close()
	}
}

func (a *Amqp) getChannel(topic string) (<-chan amqp.Delivery, error) {
	ch, err := a.conn.Channel()

	if err != nil {
		return nil, err
	}
	//defer ch.Close()
	q, err := ch.QueueDeclare(
		topic, true, false, false, false, nil,
	)

	consume, _ := ch.Consume(
		q.Name, "", true, false, false, false, nil,
	)

	return consume, nil
}

func (a *Amqp) handle(h TopicHandle, consume <-chan amqp.Delivery) {
	for d := range consume {
		h.Consume(d.Body)
	}
}

func (a *Amqp) Register(topicHandles ...TopicHandle) {
	a.topicHandles = append(a.topicHandles, topicHandles...)
}

func (a *Amqp) StartConsume() {
	for key, h := range a.topicHandles {
		log.Printf("message queue [%s] start", h.TopicName())
		go func(h TopicHandle) {
			var (
				consume <-chan amqp.Delivery
				err     error
			)
			if consume, err = a.getChannel(h.TopicName()); err != nil {
				return
			}
			a.handle(h, consume)
		}(a.topicHandles[key])
	}
}
