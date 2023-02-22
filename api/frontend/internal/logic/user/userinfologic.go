package user

import (
	"context"
	"dcs/api/frontend/internal/svc"
	"dcs/api/frontend/internal/types"
	"dcs/common"
	"dcs/rpc/user/user"
	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserinfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserinfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserinfoLogic {
	return UserinfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserinfoLogic) Userinfo(req *types.IdentificationReq) (resp *types.UserinfoResp, err error) {
	// todo: add your logic here and delete this line

	claims, err := common.ParseToken(l.svcCtx.Config.Auth.AccessSecret, req.Authorization)

	if err != nil {
		return nil, err
	}
	reply, err := l.svcCtx.UserRpc.GetUser(l.ctx, &user.IdReq{Id: cast.ToInt64((*claims)["userId"])})

	if err != nil {
		return nil, err
	}

	return &types.UserinfoResp{
		Id:     1,
		Name:   reply.Name,
		Gender: reply.Gender,
	}, nil
}
