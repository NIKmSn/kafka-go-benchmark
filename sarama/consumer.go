package sarama

import (
	"context"

	"github.com/IBM/sarama"
)

type Consumer struct {
	c     sarama.ConsumerGroup
	topic string
}

func NewConsumer(groupID string, addr []string, topic string) (Consumer, error) {
	cfg := sarama.NewConfig()
	cfg.Consumer.Offsets.AutoCommit.Enable = true
	cfg.Consumer.Return.Errors = false

	cg, err := sarama.NewConsumerGroup(addr, groupID, cfg)
	if err != nil {
		return Consumer{}, err
	}

	return Consumer{c: cg, topic: topic}, nil
}

func (c Consumer) Consume(f func([]byte) bool) error {
	return c.c.Consume(context.Background(), []string{c.topic}, simpleHandler{f: f})
}

type simpleHandler struct {
	f func([]byte) bool
}

func (simpleHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (simpleHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (h simpleHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		if !h.f(msg.Value) {
			return nil
		}

		sess.MarkMessage(msg, "")
	}
	return nil
}
