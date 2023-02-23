package queue

import (
	"dcs/rpc/producer/internal/config"
	"fmt"
	"github.com/streadway/amqp"
)

type Amqp struct {
	conn *amqp.Connection
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

func (a *Amqp) Publish(topic string, body []byte) {
	ch, err := a.conn.Channel()

	if err != nil {
		return
	}
	defer ch.Close()
	q, err := ch.QueueDeclare(
		topic, true, false, false, false, nil,
	)

	err = ch.Publish(
		"", q.Name, false, false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})

	if err != nil {
		fmt.Printf("public err %s", err)
	}
}
