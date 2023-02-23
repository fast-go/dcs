package topic

import (
	"context"
	"dcs/common/define"
	"dcs/rpc/consumer/internal/svc"
	"github.com/hashicorp/go-uuid"
	"log"
)

type LoginTopic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext

	topicId string
}

func NewLoginTopic(svcCtx *svc.ServiceContext) *LoginTopic {
	topicId, _ := uuid.GenerateUUID()
	return &LoginTopic{
		svcCtx:  svcCtx,
		topicId: topicId,
	}
}

func (r *LoginTopic) TopicName() string { return define.LoginTopic }
func (r *LoginTopic) Consume(body []byte) {
	//todo 处理具体要执行的事件
	log.Printf("[%s:%s] Received a message: %s", r.TopicName(), r.topicId, body)
}
