package product

import (
	"context"
	"dcs/rpc/product/product"

	"dcs/api/frontend/internal/svc"
	"dcs/api/frontend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) FindPageLogic {
	return FindPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindPageLogic) FindPage(req *types.ProductFindPageReq) (resp *types.ProductFindPageResp, err error) {
	// todo: add your logic here and delete this line
	res, _ := l.svcCtx.ProductRpc.FindPage(l.ctx, &product.FindPageReq{
		Limit:   req.Limit,
		Page:    req.Page,
		Keyword: req.Keyword,
	})

	var d types.ProductFindPageResp

	for _, v := range res.List {
		d.List = append(d.List, types.ProductGetDetailResp{
			Id:    v.Id,
			Name:  v.Name,
			Price: v.Price,
			Stock: v.Stock,
		})
	}

	d.Total = res.Total

	return &d, nil
}
