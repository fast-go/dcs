package logic

import (
	"context"
	"dcs/rpc/order/internal/svc"
	"dcs/rpc/order/order"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateLogic) Create(in *order.CreateOrderReq) (*order.CreateOrderResp, error) {
	// todo: add your logic here and delete this line

	fmt.Println(l.svcCtx.KqCreateOrderPusherService.Publish("test"))
	return &order.CreateOrderResp{}, nil

	// 查询产品是否存在
	//productRes, err := l.svcCtx.ProductRpc.GetProduct(l.ctx, &product.DetailReq{
	//	Id: in.ProductId,
	//})
	//if err != nil {
	//	return nil, err
	//}
	//// 判断产品库存是否充足
	//if productRes.Stock <= 0 {
	//	return nil, status.Error(500, "产品库存不足")
	//}
	//
	//newOrder := model.Order{
	//	ProductName: productRes.Name,
	//	ProductId:   productRes.Id,
	//	Status:      0,
	//	Num:         1,
	//}
	//
	//// 创建订单
	//res, err := l.svcCtx.OrderModel.Insert(l.ctx, &newOrder)
	//if err != nil {
	//	return nil, status.Error(500, err.Error())
	//}
	//
	//newOrder.Id, err = res.LastInsertId()
	//if err != nil {
	//	return nil, status.Error(500, err.Error())
	//}
	//
	//// 更新产品库存
	//_, err = l.svcCtx.ProductRpc.Update(l.ctx, &product.UpdateRequest{
	//	Id:     productRes.Id,
	//	Name:   productRes.Name,
	//	Desc:   productRes.Desc,
	//	Stock:  productRes.Stock - 1,
	//	Amount: productRes.Amount,
	//	Status: productRes.Status,
	//})
	//if err != nil {
	//	return nil, err
	//}
	//
	//return &order.CreateOrderResp{
	//	Id: newOrder.Id,
	//}, nil

}
