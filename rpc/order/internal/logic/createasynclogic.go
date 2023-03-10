package logic

import (
	"context"
	"dcs/common/define"
	"dcs/rpc/producer/producer"
	"encoding/json"

	"dcs/rpc/order/internal/svc"
	"dcs/rpc/order/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateAsyncLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateAsyncLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateAsyncLogic {
	return &CreateAsyncLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateAsyncLogic) CreateAsync(in *order.CreateOrderReq) (*order.CreateOrderResp, error) {
	// todo: add your logic here and delete this line
	type CreateOrder struct {
		ProductId int64 `json:"product_id"`
	}
	cr := CreateOrder{ProductId: in.ProductId}
	body, _ := json.Marshal(cr)
	if _, err := l.svcCtx.ProducerRpc.Publish(l.ctx, &producer.Request{
		Topic: define.CreateOrderTopic,
		Body:  body,
	}); err != nil {
		return nil, err
	}
	return &order.CreateOrderResp{}, nil
}
