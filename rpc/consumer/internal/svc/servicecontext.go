package svc

import (
	"context"
	"dcs/rpc/consumer/internal/config"
	"dcs/rpc/consumer/internal/server/queue"
	"dcs/rpc/consumer/internal/topic"
)

type ServiceContext struct {
	Config    config.Config
	QueueAmqp *queue.Amqp
}

func NewServiceContext(c config.Config) *ServiceContext {
	queueAmqp := queue.NewAmqp(c)
	queueAmqp.Register(
		topic.NewLoginTopic(context.Background()),
		topic.NewRegisterTopic(context.Background()),
	)

	return &ServiceContext{
		Config:    c,
		QueueAmqp: queueAmqp,
	}
}

func (s *ServiceContext) Close() {
	s.QueueAmqp.Close()
}
