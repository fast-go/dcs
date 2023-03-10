package config

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}

	Log logx.LogConf

	UserRpc     zrpc.RpcClientConf
	ProducerRpc zrpc.RpcClientConf
	ProductRpc  zrpc.RpcClientConf
	OrderRpc    zrpc.RpcClientConf
}
