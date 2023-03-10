package logic

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/dtm-labs/client/dtmgrpc"
	_ "github.com/dtm-labs/driver-gozero"
	"google.golang.org/grpc/status"

	"dcs/rpc/order/internal/svc"
	"dcs/rpc/order/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateRevertLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateRevertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateRevertLogic {
	return &CreateRevertLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateRevertLogic) CreateRevert(in *order.CreateOrderReq) (*order.CreateOrderResp, error) {
	// todo: add your logic here and delete this line
	// 获取 RawDB
	db, err := l.svcCtx.SqlConn.RawDB()
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	// 获取子事务屏障对象
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}
	// 开启子事务屏障
	if err := barrier.CallWithDB(db, func(tx *sql.Tx) error {
		// 查询用户是否存在
		//_, err := l.svcCtx.UserRpc.GetUser(l.ctx, &user.UserInfoRequest{
		//	Id: in.Uid,
		//})
		//if err != nil {
		//	return fmt.Errorf("用户不存在")
		// 查询用户最新创建的订单
		resOrder, err := l.svcCtx.OrderModel.FindOne(l.ctx, in.ProductId)
		if err != nil {
			return fmt.Errorf("no order record")
		}
		// 修改订单状态9，标识订单已失效，并更新订单
		resOrder.Status = 9
		err = l.svcCtx.OrderModel.TxUpdate(l.ctx, tx, resOrder)
		if err != nil {
			return fmt.Errorf("update order status fail")
		}
		return nil
	}); err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &order.CreateOrderResp{}, nil
}
