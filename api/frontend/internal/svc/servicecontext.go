package svc

import (
	"dcs/api/frontend/internal/config"
	"dcs/rpc/producer/producerclient"
	"dcs/rpc/user/userclient"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	UserRpc     userclient.User
	ProducerRpc producerclient.Producer
}

type MultiWriter struct {
	writer        logx.Writer
	consoleWriter logx.Writer
}

func NewServiceContext(c config.Config) *ServiceContext {
	//logx.MustSetup(logx.LogConf{
	//	Mode: "file",
	//	Stat: true,
	//})
	logx.Error("错误的")

	return &ServiceContext{
		Config:      c,
		UserRpc:     userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		ProducerRpc: producerclient.NewProducer(zrpc.MustNewClient(c.ProducerRpc)),
	}
}
