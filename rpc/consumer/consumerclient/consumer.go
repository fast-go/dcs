// Code generated by goctl. DO NOT EDIT.
// Source: consumer.proto

package consumerclient

import (
	"context"

	"dcs/rpc/consumer/consumer"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Request  = consumer.Request
	Response = consumer.Response

	Consumer interface {
		Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	}

	defaultConsumer struct {
		cli zrpc.Client
	}
)

func NewConsumer(cli zrpc.Client) Consumer {
	return &defaultConsumer{
		cli: cli,
	}
}

func (m *defaultConsumer) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	client := consumer.NewConsumerClient(m.cli.Conn())
	return client.Ping(ctx, in, opts...)
}