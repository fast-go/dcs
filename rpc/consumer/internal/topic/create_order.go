package topic

import (
	"context"
	"dcs/common/define"
	"dcs/rpc/consumer/internal/config"
	"dcs/rpc/order/order"
	"dcs/rpc/order/orderclient"
	"dcs/rpc/product/product"
	"dcs/rpc/product/productclient"
	"encoding/json"
	"fmt"
	"github.com/dtm-labs/client/dtmgrpc"
	_ "github.com/dtm-labs/driver-gozero"
	"github.com/hashicorp/go-uuid"
	"github.com/zeromicro/go-zero/core/logx"
	"log"
)

type CreateOrderOption struct {
	ctx context.Context

	Config config.Config

	ProductRpc productclient.Product
	OrderRpc   orderclient.Order
}

type CreateOrderTopic struct {
	topicId string
	Option  CreateOrderOption
}

func NewCreateOrderTopic(option CreateOrderOption) *CreateOrderTopic {
	topicId, _ := uuid.GenerateUUID()
	option.ctx = context.Background()
	return &CreateOrderTopic{
		topicId: topicId,
		Option:  option,
	}
}

func (r *CreateOrderTopic) TopicName() string { return define.CreateOrderTopic }

type CreateOrder struct {
	ProductId int64 `json:"product_id"`
}

func (r *CreateOrderTopic) Consume(body []byte) {
	//todo 处理具体要执行的事件
	log.Printf("[%s:%s] Received a message: %s", r.TopicName(), r.topicId, body)

	var (
		co  CreateOrder
		err error
		//pd  *product.ProductDetail
		//createOrderResp *order.CreateOrderResp
		//decrStockRes    *product.DecrStockResp
	)

	if err = json.Unmarshal(body, &co); err != nil {
		return
	}

	//pd, err = r.Option.ProductRpc.GetProduct(r.Option.ctx, &product.DetailReq{Id: co.ProductId})
	//
	//fmt.Println(err)
	//
	//if err != nil {
	//	logx.Errorf(fmt.Sprintf("find product err: %s", err))
	//	return
	//}
	////
	////fmt.Println(11)
	////
	//if pd.Stock <= 0 {
	//	fmt.Println("no stock")
	//	return
	//}
	//
	//fmt.Println(22)
	//
	//// dtm 服务的 etcd 注册地址
	var dtmServer = "etcd://localhost:2379/dtmservice"
	//var dtmServer = "http://localhost:36790/api/dtmsvr"

	productRpcService, err := r.Option.Config.ProductRpc.BuildTarget()

	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	// 创建一个gid
	gid := dtmgrpc.MustGenGid(dtmServer)
	////
	//fmt.Println(gid)
	//fmt.Println(productRpcService)
	//
	//return
	orderRpcService, _ := r.Option.Config.OrderRpc.BuildTarget()

	//orderRpcService = "http://127.0.0.1:9090"

	//fmt.Println("======")
	fmt.Println(orderRpcService)
	//fmt.Println("======")
	//
	//var dtmServer = "http://localhost:36789/api/dtmsvr"

	//// 创建一个saga协议的事务
	saga := dtmgrpc.NewSagaGrpc(dtmServer, gid).
		Add(orderRpcService+"/order.order/create", orderRpcService+"/order.order/createRevert", &order.CreateOrderReq{ProductId: co.ProductId}).
		Add(productRpcService+"/product.product/decrStock", productRpcService+"/product.product/decrStockRevert", &product.DecrStockReq{
			Id:  co.ProductId,
			Num: 1,
		})

	// 事务提交
	err = saga.Submit()
	if err != nil {
		logx.Error(err)
		return
	}

	if err != nil {
		return
	}

}
