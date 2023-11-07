package kafka

import (
	"fmt"
	"github.com/IBM/sarama"
)

type SyncProducer struct {
	producer sarama.SyncProducer
}

func NewSyncProducer(brokers []string, opts ...Option) (*SyncProducer, error) {
	c := sarama.NewConfig()

	c.Producer.Partitioner = sarama.NewHashPartitioner
	c.Producer.RequiredAcks = sarama.WaitForAll

	c.Net.MaxOpenRequests = 1
	c.Producer.Idempotent = true

	c.Producer.Return.Successes = true
	c.Producer.Return.Errors = true

	for _, opt := range opts {
		if err := opt.Apply(c); err != nil {
			return nil, err
		}
	}

	producer, err := sarama.NewSyncProducer(brokers, c)
	if err != nil {
		return nil, err
	}

	return &SyncProducer{producer: producer}, nil
}

func (sp *SyncProducer) Produce(topic string, key string, payload []byte) error {
	message := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.ByteEncoder(payload),
	}

	partition, offset, err := sp.producer.SendMessage(message)
	fmt.Println(topic, ":", key, " = ", partition, ":", offset)

	return err
}

func (sp *SyncProducer) Close() error {
	return sp.producer.Close()
}
