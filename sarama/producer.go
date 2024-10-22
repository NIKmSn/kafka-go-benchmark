package sarama

import (
	"time"

	"github.com/IBM/sarama"
)

type Producer struct {
	p sarama.AsyncProducer
}

func NewProducer(addrs []string) (Producer, error) {
	sarCfg := sarama.NewConfig()

	sarCfg.ChannelBufferSize = 1024

	sarCfg.Producer.RequiredAcks = sarama.WaitForLocal
	sarCfg.Producer.Timeout = 10 * time.Second
	sarCfg.Producer.MaxMessageBytes = 1000000

	sarCfg.Producer.Return.Errors = false
	sarCfg.Producer.Return.Successes = false

	p, err := sarama.NewAsyncProducer(addrs, sarCfg)
	if err != nil {
		return Producer{}, nil
	}

	return Producer{p: p}, nil
}

func (p Producer) SendAsync(msg []byte, key string, topic string) error {
	sarMsg := &sarama.ProducerMessage{
		Topic:    topic,
		Key:      sarama.StringEncoder(key),
		Value:    sarama.ByteEncoder(msg),
		Metadata: nil,
	}

	p.p.Input() <- sarMsg

	return nil
}

func (p Producer) SendSync(msg []byte, key string, topic string) error {
	expectation := make(chan error, 1)
	sarMsg := &sarama.ProducerMessage{
		Topic:    topic,
		Key:      sarama.StringEncoder(key),
		Value:    sarama.ByteEncoder(msg),
		Metadata: expectation,
	}

	p.p.Input() <- sarMsg

	return <-expectation
}

func (p Producer) Close() {
	p.p.Close()
}
