package svc

import (
	"dcs/gen/model"
	"dcs/rpc/order/internal/config"
	"dcs/rpc/order/internal/server/queue"
	"dcs/rpc/product/productclient"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	ProductRpc productclient.Product

	OrderModel model.OrderModel

	KqCreateOrderPusherService *queue.KqPusherService
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DataSource)

	go func() {
		queue.NewKafkaConsumer(c)
	}()

	return &ServiceContext{
		Config:                     c,
		ProductRpc:                 productclient.NewProduct(zrpc.MustNewClient(c.ProductRpc)),
		OrderModel:                 model.NewOrderModel(sqlConn, c.Cache),
		KqCreateOrderPusherService: queue.NewKafkaService(c, queue.TopicCreateOrder),
	}
}
