package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf

	Amqp struct {
		Host string
	}

	ProductRpc zrpc.RpcClientConf
	OrderRpc   zrpc.RpcClientConf
}
