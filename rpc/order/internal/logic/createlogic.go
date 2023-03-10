package logic

import (
	"context"
	"database/sql"
	"dcs/gen/model"
	"dcs/rpc/order/internal/svc"
	"dcs/rpc/order/order"
	"fmt"
	"github.com/dtm-labs/client/dtmgrpc"
	_ "github.com/dtm-labs/driver-gozero"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
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
	// 获取 RawDB
	fmt.Println(888)
	db, err := l.svcCtx.SqlConn.RawDB()
	if err != nil {
		return nil, status.Error(500, err.Error())
	}
	fmt.Println(7777)
	// 获取子事务屏障对象
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}
	fmt.Println(333333)
	// 开启子事务屏障
	if err := barrier.CallWithDB(db, func(tx *sql.Tx) error {
		// 查询用户是否存在
		//_, err := l.svcCtx.UserRpc.GetUser(l.ctx,&user.IdReq{Id: 1})
		//if err != nil {
		//	return fmt.Errorf("用户不存在")
		//}
		newOrder := model.Order{
			Uid:       1,
			ProductId: in.ProductId,
			Status:    0,
		}
		fmt.Println(888999)
		// create order
		_, err = l.svcCtx.OrderModel.TxInsert(l.ctx, tx, &newOrder)
		fmt.Println(11111)
		fmt.Println(err)
		if err != nil {
			return fmt.Errorf("create order error")
		}
		return nil
	}); err != nil {
		return nil, status.Error(500, err.Error())
	}

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
