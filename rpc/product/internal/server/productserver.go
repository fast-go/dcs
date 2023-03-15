// Code generated by goctl. DO NOT EDIT.
// Source: product.proto

package server

import (
	"context"

	"dcs/rpc/product/internal/logic"
	"dcs/rpc/product/internal/svc"
	"dcs/rpc/product/product"
)

type ProductServer struct {
	svcCtx *svc.ServiceContext
	product.UnimplementedProductServer
}

func NewProductServer(svcCtx *svc.ServiceContext) *ProductServer {
	return &ProductServer{
		svcCtx: svcCtx,
	}
}

func (s *ProductServer) GetProduct(ctx context.Context, in *product.DetailReq) (*product.ProductDetail, error) {
	l := logic.NewGetProductLogic(ctx, s.svcCtx)
	return l.GetProduct(in)
}

func (s *ProductServer) FindPage(ctx context.Context, in *product.FindPageReq) (*product.FindPageRes, error) {
	l := logic.NewFindPageLogic(ctx, s.svcCtx)
	return l.FindPage(in)
}

func (s *ProductServer) DecrStock(ctx context.Context, in *product.DecrStockReq) (*product.DecrStockResp, error) {
	l := logic.NewDecrStockLogic(ctx, s.svcCtx)
	return l.DecrStock(in)
}

func (s *ProductServer) DecrStockRevert(ctx context.Context, in *product.DecrStockReq) (*product.DecrStockResp, error) {
	l := logic.NewDecrStockRevertLogic(ctx, s.svcCtx)
	return l.DecrStockRevert(in)
}
