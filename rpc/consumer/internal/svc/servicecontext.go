package svc

import (
	"context"
	"dcs/rpc/consumer/internal/config"
	"dcs/rpc/consumer/internal/server/queue"
	"dcs/rpc/consumer/internal/topic"
	"dcs/rpc/order/orderclient"
	"dcs/rpc/product/productclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	QueueAmqp *queue.Amqp

	ProductRpc productclient.Product
	OrderRpc   orderclient.Order
}

func NewServiceContext(c config.Config) *ServiceContext {
	productRpc := productclient.NewProduct(zrpc.MustNewClient(c.ProductRpc))
	orderRpc := orderclient.NewOrder(zrpc.MustNewClient(c.OrderRpc))

	queueAmqp := queue.NewAmqp(c)
	queueAmqp.Register(
		topic.NewLoginTopic(context.Background()),
		topic.NewRegisterTopic(context.Background()),
		topic.NewCreateOrderTopic(topic.CreateOrderOption{
			Config:     c,
			ProductRpc: productRpc,
			OrderRpc:   orderRpc,
		}),
	)

	return &ServiceContext{
		Config:     c,
		QueueAmqp:  queueAmqp,
		ProductRpc: productRpc,
		OrderRpc:   orderRpc,
	}
}

func (s *ServiceContext) Close() {
	s.QueueAmqp.Close()
}
