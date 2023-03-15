package logic

import (
	"context"
	"dcs/gen/model"
	"dcs/rpc/product/internal/svc"
	"dcs/rpc/product/product"
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/zeromicro/go-zero/core/logx"
	"reflect"
)

type FindPageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindPageLogic {
	return &FindPageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// FindPage 参考 https://blog.csdn.net/finghting321/article/details/103459353
func (l *FindPageLogic) FindPage(in *product.FindPageReq) (*product.FindPageRes, error) {
	// todo: add your logic here and delete this line
	var (
		searchResult *elastic.SearchResult
		scrollId     string
		list         []*product.ProductDetail
		err          error
	)
	if scrollId != "" {
		searchResult, err = l.svcCtx.ElasticClient.Scroll("1m").ScrollId(scrollId).Do(context.TODO())
	} else {
		query := elastic.NewWildcardQuery("name", fmt.Sprintf("*%s*", in.Keyword))
		//普通搜索
		searchService := l.svcCtx.ElasticClient.Search().Index(l.svcCtx.Config.ElasticSearch.Index).
			Query(query).Sort("id", false).From(int(in.Page * in.Limit)).Size(int(in.Limit))
		searchResult, err = searchService.Do(context.TODO())
		//游标搜索暂时不使用(性能更高)
		//scrollService := client.Scroll("test").Query(query).Scroll("1m").Sort("id", false).Size(int(in.Limit))
		//searchResult, err = scrollService.Do(context.TODO())
	}
	if err != nil {
		logx.Errorf("elastic首次查询游标失败:%v\n", err)
		return nil, err
	}
	scrollId = searchResult.ScrollId
	var p model.Product
	for _, item := range searchResult.Each(reflect.TypeOf(p)) {
		t := item.(model.Product)
		list = append(list, &product.ProductDetail{
			Id:    t.Id,
			Name:  t.Name,
			Price: t.Price,
			Stock: t.Stock,
		})
	}
	return &product.FindPageRes{
		List:  list,
		Total: searchResult.TotalHits(),
	}, nil
}
