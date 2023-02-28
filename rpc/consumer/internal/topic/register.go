package topic

import (
	"context"
	"dcs/common/define"
	"github.com/hashicorp/go-uuid"
	"log"
)

type RegisterTopic struct {
	ctx context.Context

	topicId string
}

func NewRegisterTopic(ctx context.Context) *RegisterTopic {
	topicId, _ := uuid.GenerateUUID()
	return &RegisterTopic{
		ctx:     ctx,
		topicId: topicId,
	}
}

func (r *RegisterTopic) TopicName() string { return define.RegisterTopic }
func (r *RegisterTopic) Consume(body []byte) {
	//todo 处理具体要执行的事件
	log.Printf("[%s:%s] Received a message: %s", r.TopicName(), r.topicId, body)
}
