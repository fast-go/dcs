package logic

import (
	"context"
	"dcs/gen/model"
	"dcs/rpc/product/internal/svc"
	"dcs/rpc/product/product"
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logx"
	"reflect"
	"strings"
)

type FindPageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

const cursorSeparator = "_"

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
		list         []*product.ProductDetail
		err          error
		cursor       string
		hasMore      bool = true
		query        elastic.Query
	)
	if in.Keyword != "" {
		query = elastic.NewWildcardQuery("name", fmt.Sprintf("*%s*", in.Keyword))
	} else {
		query = elastic.NewMatchAllQuery()
	}
	//普通搜索
	//searchService := l.svcCtx.ElasticClient.Search().Index(l.svcCtx.Config.ElasticSearch.Index).
	//	Query(query).Sort("id", false).From(int(in.Page * in.Limit)).Size(int(in.Limit))
	//searchResult, err = searchService.Do(context.TODO())
	//游标搜索暂时不使用(性能更高,占用资源,存在过期问题)
	//scrollService := l.svcCtx.ElasticClient.Scroll("test").Query(query).Scroll("1m").Sort("id", false).Size(int(in.Limit))
	//searchResult, err = scrollService.Do(context.TODO())
	//深度搜索
	searchService := l.svcCtx.ElasticClient.Search(l.svcCtx.Config.ElasticSearch.Index).Query(query).
		Sort("create_at", true).
		Sort("id", true).
		Size(int(in.Limit))
	if in.Cursor != "" && strings.Contains(in.Cursor, cursorSeparator) {
		cursorSlice := strings.Split(in.Cursor, cursorSeparator)
		createAt := cursorSlice[0]
		id := cursorSlice[1]
		searchService.SearchAfter(createAt, id)
	}
	if searchResult, err = searchService.Do(context.TODO()); err != nil {
		logx.Errorf("elastic首次查询游标失败:%v\n", err)
		return nil, err
	}
	hitsLen := len(searchResult.Hits.Hits)
	if hitsLen <= 0 || hitsLen < int(in.Limit) {
		hasMore = false
	}
	if hitsLen > 0 && len(searchResult.Hits.Hits[len(searchResult.Hits.Hits)-1].Sort) > 1 {
		cursor = strings.Join(cast.ToStringSlice(searchResult.Hits.Hits[len(searchResult.Hits.Hits)-1].Sort), cursorSeparator)
	}
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
	return &product.FindPageRes{List: list,
		Total:   searchResult.TotalHits(),
		Cursor:  cursor,
		HasMore: hasMore,
	}, nil
}
