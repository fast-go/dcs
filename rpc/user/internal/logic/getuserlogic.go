package logic

import (
	"context"
	"dcs/common"
	"dcs/rpc/user/internal/svc"
	"dcs/rpc/user/user"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"os"
)

import _ "github.com/dtm-labs/driver-gozero"

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *user.IdReq) (*user.UserInfoReply, error) {
	// todo: add your logic here and delete this line
	//获取Id,查询数据库获取用户信息
	var ip string
	if localIp, err := common.GetLocalIP(); err == nil && len(localIp) > 0 {
		ip = localIp[0]
	}

	u, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return &user.UserInfoReply{
		Id:     in.GetId(),
		Name:   fmt.Sprintf("username %s:ip:%s,TZ = %s", u.Username, ip, os.Getenv("TZ")),
		Number: "",
		Gender: "woman",
	}, nil
}
