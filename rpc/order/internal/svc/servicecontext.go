package svc

import (
	"dcs/gen/model"
	"dcs/rpc/order/internal/config"
	"dcs/rpc/producer/producerclient"
	"dcs/rpc/product/productclient"
	"dcs/rpc/user/userclient"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	SqlConn     sqlx.SqlConn
	UserRpc     userclient.User
	ProducerRpc producerclient.Producer
	ProductRpc  productclient.Product

	OrderModel model.OrderModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DataSource)

	return &ServiceContext{
		Config:      c,
		SqlConn:     sqlConn,
		UserRpc:     userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		ProducerRpc: producerclient.NewProducer(zrpc.MustNewClient(c.ProducerRpc)),
		ProductRpc:  productclient.NewProduct(zrpc.MustNewClient(c.ProductRpc)),
		OrderModel:  model.NewOrderModel(sqlConn, c.Cache),
	}
}
