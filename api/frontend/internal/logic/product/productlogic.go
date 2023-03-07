package product

import (
	"context"
	"dcs/rpc/product/product"

	"dcs/api/frontend/internal/svc"
	"dcs/api/frontend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) ProductLogic {
	return ProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductLogic) Product(req *types.ProductGetDetailReq) (resp *types.ProductGetDetailResp, err error) {
	// todo: add your logic here and delete this line

	result, err := l.svcCtx.ProductRpc.GetProduct(l.ctx, &product.DetailReq{Id: req.ProductId})

	if err != nil {
		return nil, err
	}

	return &types.ProductGetDetailResp{
		Id:    result.Id,
		Name:  result.Name,
		Price: result.Price,
		Stock: result.Stock,
	}, nil
}
