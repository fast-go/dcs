package logic

import (
	"context"
	"dcs/rpc/producer/internal/svc"
	"dcs/rpc/producer/producer"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPublishLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishLogic {
	return &PublishLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PublishLogic) Publish(in *producer.Request) (*producer.Response, error) {
	// todo: add your logic here and delete this line
	//这里可以对投递的任务数据进行限制处理
	l.svcCtx.QueueAmqp.Publish(in.Topic, in.Body)
	return &producer.Response{}, nil
}
