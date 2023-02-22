package topic

import (
	"context"
	"dcs/rpc/consumer/internal/svc"
)

type RegisterTopic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterTopic(svcCtx *svc.ServiceContext) *RegisterTopic {
	return &RegisterTopic{
		svcCtx: svcCtx,
	}
}

func (r *RegisterTopic) listen() {

}
