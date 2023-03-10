package order

import (
	"context"
	"dcs/rpc/order/order"

	"dcs/api/frontend/internal/svc"
	"dcs/api/frontend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) OrderLogic {
	return OrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderLogic) Order(req *types.CreateOrderReq) (resp *types.CreateOrderResp, err error) {
	// todo: add your logic here and delete this line
	result, err := l.svcCtx.OrderRpc.CreateAsync(l.ctx, &order.CreateOrderReq{ProductId: req.ProductId})

	if err != nil {
		return nil, err
	}
	return &types.CreateOrderResp{
		Id:          result.Id,
		ProductName: result.ProductName,
		ProductId:   result.ProductId,
		Uid:         result.Uid,
		Num:         result.Num,
	}, nil
}
