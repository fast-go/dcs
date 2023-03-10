package svc

import (
	"dcs/rpc/product/internal/config"
	"dcs/rpc/product/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	SqlConn      sqlx.SqlConn
	ProductModel model.ProductModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DataSource)
	return &ServiceContext{
		Config:       c,
		ProductModel: model.NewProductModel(sqlConn, c.Cache),
		SqlConn:      sqlConn,
	}
}
