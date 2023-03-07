package svc

import (
	"dcs/api/frontend/internal/config"
	"dcs/rpc/order/orderclient"
	"dcs/rpc/producer/producerclient"
	"dcs/rpc/product/productclient"
	"dcs/rpc/user/userclient"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	UserRpc     userclient.User
	ProducerRpc producerclient.Producer
	ProductRpc  productclient.Product
	OrderRpc    orderclient.Order
}

type MultiWriter struct {
	writer        logx.Writer
	consoleWriter logx.Writer
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		UserRpc:     userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		ProducerRpc: producerclient.NewProducer(zrpc.MustNewClient(c.ProducerRpc)),
		ProductRpc:  productclient.NewProduct(zrpc.MustNewClient(c.ProductRpc)),
		OrderRpc:    orderclient.NewOrder(zrpc.MustNewClient(c.OrderRpc)),
	}
}
