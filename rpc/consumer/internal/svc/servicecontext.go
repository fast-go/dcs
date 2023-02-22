package svc

import "dcs/rpc/consumer/internal/config"

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	//对服务进行初始化
	return &ServiceContext{
		Config: c,
	}
}
