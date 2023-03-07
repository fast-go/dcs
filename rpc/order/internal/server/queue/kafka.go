package queue

import (
	"dcs/rpc/order/internal/config"
	"fmt"
	"github.com/zeromicro/go-queue/kq"
)

const TopicCreateOrder = "kq"

type KqPusherService struct {
	Pusher *kq.Pusher
}

func NewKafkaService(c config.Config, topic string) *KqPusherService {
	pusher := kq.NewPusher(c.Kq.Brokers, topic)
	return &KqPusherService{
		Pusher: pusher,
	}
}

func (k *KqPusherService) Publish(msg string) error {
	return k.Pusher.Push(msg)
}

func (k *KqPusherService) Close() {
	_ = k.Pusher.Close()
}

func NewKafkaConsumer(c config.Config) {
	fmt.Println("启动")
	q := kq.MustNewQueue(c.Kq, kq.WithHandle(func(k, v string) error {
		fmt.Println("============================")
		fmt.Printf("=> %s\n", v)
		return nil
	}))
	fmt.Println("3333")
	defer q.Stop()
	q.Start()

	fmt.Println("结束")
}
