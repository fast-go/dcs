package topic

import (
	"context"
	"dcs/common/define"
	"github.com/hashicorp/go-uuid"
	"log"
)

type LoginTopic struct {
	ctx     context.Context
	topicId string
}

func NewLoginTopic(ctx context.Context) *LoginTopic {
	topicId, _ := uuid.GenerateUUID()
	return &LoginTopic{
		topicId: topicId,
		ctx:     ctx,
	}
}

func (r *LoginTopic) TopicName() string { return define.LoginTopic }
func (r *LoginTopic) Consume(body []byte) {
	//todo 处理具体要执行的事件
	log.Printf("[%s:%s] Received a message: %s", r.TopicName(), r.topicId, body)
}
