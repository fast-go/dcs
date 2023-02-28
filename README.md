# dcs

docker构建rpc应用
```
docker build -t rpcservice-user:v1 -f ./rpc/user/Dockerfile .
docker run -p 8080:8080 -d rpcservice-user:v1
```

示例Dockerfile代码，也可以通过 `goctl docker -go user.go` 生成
```
FROM golang:1.19-alpine as golang

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct

WORKDIR /www

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
COPY rpc/user/etc/user.yaml /www/etc/user.yaml
RUN  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -tags netgo -o /www/user rpc/user/user.go

# 暴露服务端口
#EXPOSE 8080

FROM scratch

WORKDIR /www

COPY --from=golang /www/user /www/user
COPY --from=golang /www/etc /www/etc

#RUN #chmod +x /app/user
ENTRYPOINT  ["./user", "-f", "etc/user.yaml"]
```

通过容器编排工具 ```docker-compose``` 启动docker镜像

```
docker-compose up -d
```

如果需要重新构建只需要加上 ```--force-recreate --build```参数
```
docker-compose up -d --force-recreate --build
```


goctl生成k8s脚本文件

```

[//]: # (goctl kube deploy --name greet --image greet:v1 --namespace default --port 8888 -nodePort 8888 --nodePort 32000 -0 greet.yam)
```

运行 rabbitmq

```
docker run -id --hostname myrabbit --name rabbitmq1 -p 15672:15672 -p 5672:5672 rabbitmq
```

gorm生成model代码

```
gentool -dsn "root:root@tcp(localhost:3306)/dcs?charset=utf8mb4&parseTime=True&loc=Local" -tables "user"
```


docker 之间网络不通剋有通过  ```docker inspect 容器id``` 查看容器的ip,然后将对应的ip地址更改就可以访问呢


启动日志同步到kafka服务

``` 
./filebeat -e -c filebeat.yaml -d publish
```

日志收集 参考 https://blog.csdn.net/jj546630576/article/details/123128581

