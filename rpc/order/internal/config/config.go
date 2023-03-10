package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	DataSource string

	Cache cache.CacheConf

	zrpc.RpcServerConf

	UserRpc     zrpc.RpcClientConf
	ProducerRpc zrpc.RpcClientConf
	ProductRpc  zrpc.RpcClientConf

	//Kq kq.KqConf
}
