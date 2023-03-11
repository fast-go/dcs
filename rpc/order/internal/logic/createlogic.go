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
		newOrder := model.Order{
			Uid:       1,
			ProductId: in.ProductId,
			Status:    0,
			Num:       1,
			OrderNum:  in.OrderNum,
		}
		_, err = l.svcCtx.OrderModel.TxInsert(l.ctx, tx, &newOrder)
		if err != nil {
			logx.Error("create order err:", err)
			return fmt.Errorf("create order error")
		}
		return nil
	}); err != nil {
		return nil, status.Error(500, err.Error())
	}
	return &order.CreateOrderResp{}, nil
}
