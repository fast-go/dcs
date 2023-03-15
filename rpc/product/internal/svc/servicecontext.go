package svc

import (
	"dcs/rpc/product/internal/config"
	"dcs/rpc/product/model"
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	SqlConn      sqlx.SqlConn
	ProductModel model.ProductModel

	ElasticClient *elastic.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DataSource)

	elasticClient, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL(c.ElasticSearch.Hosts...),
	)
	if err != nil {
		panic(fmt.Sprintf("elastic start err: %s", err))
	}

	return &ServiceContext{
		Config:        c,
		ProductModel:  model.NewProductModel(sqlConn, c.Cache),
		SqlConn:       sqlConn,
		ElasticClient: elasticClient,
	}
}
