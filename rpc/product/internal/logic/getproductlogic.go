package logic

import (
	"context"
	"dcs/rpc/product/internal/svc"
	"dcs/rpc/product/product"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductLogic {
	return &GetProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetProductLogic) GetProduct(in *product.DetailReq) (*product.ProductDetail, error) {
	// todo: add your logic here and delete this line
	result, err := l.svcCtx.ProductModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &product.ProductDetail{
		Id:    result.Id,
		Name:  result.Name,
		Price: result.Price,
		Stock: result.Stock,
	}, nil
}
