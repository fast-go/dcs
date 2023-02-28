package user

import (
	"context"
	"dcs/common"
	"dcs/common/define"
	"dcs/rpc/producer/producer"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"dcs/api/frontend/internal/svc"
	"dcs/api/frontend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginLogic {
	return LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginReply, err error) {
	// todo: add your logic here and delete this line
	//这个服务可以挪到rpc当中
	if len(strings.TrimSpace(req.Username)) == 0 || len(strings.TrimSpace(req.Password)) == 0 {
		return nil, errors.New("参数错误")
	}

	type nameStruct struct {
		Id     int64
		Name   string
		Gender string
	}

	userInfo := nameStruct{Name: "jack", Id: 1, Gender: "woman"}
	//userInfo, err := l.svcCtx.UserModel.FindOneByNumber(l.ctx, req.Username)
	//switch err {
	//case nil:
	//case model.ErrNotFound:
	//	return nil, errors.New("用户名不存在")
	//default:
	//	return nil, err
	//}
	//
	//if userInfo.Password != req.Password {
	//	return nil, errors.New("用户密码不正确")
	//}

	// ---start---
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire
	jwtToken, err := common.JwtToken(l.svcCtx.Config.Auth.AccessSecret, now, l.svcCtx.Config.Auth.AccessExpire, userInfo.Id)
	if err != nil {
		return nil, err
	}
	// ---end---

	body, _ := json.Marshal(userInfo)

	logx.Error("用户登录成功")

	//登录成功发送事件到消息队列中
	_, err = l.svcCtx.ProducerRpc.Publish(l.ctx, &producer.Request{
		Topic: define.LoginTopic,
		Body:  body,
	})

	if err != nil {
		fmt.Printf("l.svcCtx.ProducerRpc.Publish err:%s", err)
	}

	return &types.LoginReply{
		Id:           userInfo.Id,
		Name:         userInfo.Name,
		Gender:       userInfo.Gender,
		AccessToken:  jwtToken,
		AccessExpire: now + accessExpire,
		RefreshAfter: now + accessExpire/2,
	}, nil
}
