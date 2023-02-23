package svc

import (
	"dcs/rpc/consumer/internal/config"
	"dcs/rpc/consumer/internal/server/queue"
)

type ServiceContext struct {
	Config    config.Config
	QueueAmqp *queue.Amqp
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		QueueAmqp: queue.NewAmqp(c),
	}
}

func (s *ServiceContext) Close() {
	s.QueueAmqp.Close()
}
