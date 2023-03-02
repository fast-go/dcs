package main

import (
	"dcs/api/frontend/internal/config"
	"dcs/api/frontend/internal/handler"
	"dcs/api/frontend/internal/svc"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/frontend-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	//禁用stat(每分钟控制台定时输出日志信息)
	logx.DisableStat()

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	//defer logx.Close()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)

	server.Start()
}
