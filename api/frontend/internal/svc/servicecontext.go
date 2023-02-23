package svc

import (
	"dcs/api/frontend/internal/config"
	"dcs/rpc/producer/producerclient"
	"dcs/rpc/user/userclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	UserRpc     userclient.User
	ProducerRpc producerclient.Producer
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		UserRpc:     userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		ProducerRpc: producerclient.NewProducer(zrpc.MustNewClient(c.ProducerRpc)),
	}
}
