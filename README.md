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

```jsunicoderegexp
goctl model mysql datasource -c -url="root:root@tcp(127.0.0.1:3306)/dcs" -table="*"  -dir="./model"
```

docker 之间网络不通剋有通过  ```docker inspect 容器id``` 查看容器的ip,然后将对应的ip地址更改就可以访问呢

启动日志同步到kafka服务

``` 
./filebeat -e -c filebeat.yaml -d publish
```

日志收集 参考 https://blog.csdn.net/jj546630576/article/details/123128581

### 链路追踪

docker 启动 jaeger服务

```
 docker run -d --name jaeger -e COLLECTOR_ZIPKIN_HTTP_PORT=9411 -p 5775:5775/udp -p 6831:6831/udp -p 6832:6832/udp -p 5778:5778 -p 16686:16686 -p 14268:14268 -p 9411:9411 jaegertracing/all-in-one:1.6
```
config.yaml 文件中配置

```
Telemetry:
  Name: user.api
  Endpoint: http://127.0.0.1:14268/api/traces
  Sampler: 1.6
  Batcher: jaeger
  ```

访问 jaeger webUi界面查看 `http://localhost:16686/`

运行 es

参考 `https://www.elastic.co/guide/en/elasticsearch/reference/current/docker.html`

```
docker run -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" docker.elastic.co/elasticsearch/elasticsearch:7.2.0
```

安装 kabana 
参考 `https://blog.csdn.net/qq_34285557/article/details/127242907`

```
docker run -d --name kabana -v ./config/:/usr/share/kibana/config -p 5601:5601 kibana:7.2.0
```

