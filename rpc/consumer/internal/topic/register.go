package topic

import (
	"context"
	"dcs/common/define"
	"dcs/rpc/consumer/internal/svc"
	"github.com/hashicorp/go-uuid"
	"log"
)

type RegisterTopic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext

	topicId string
}

func NewRegisterTopic(svcCtx *svc.ServiceContext) *RegisterTopic {
	topicId, _ := uuid.GenerateUUID()
	return &RegisterTopic{
		svcCtx:  svcCtx,
		topicId: topicId,
	}
}

func (r *RegisterTopic) TopicName() string { return define.RegisterTopic }
func (r *RegisterTopic) Consume(body []byte) {
	//todo 处理具体要执行的事件
	log.Printf("[%s:%s] Received a message: %s", r.TopicName(), r.topicId, body)
}
