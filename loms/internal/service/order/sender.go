package order

import (
	"fmt"
	"route256/loms/internal/model"
	"route256/loms/pkg/kafka"
)

type StatusSender interface {
	Send(model.OrderID, model.OrderStatus) error
	Close() error
}

type statusSender struct {
	producer *kafka.SyncProducer
	topic    string
}

func NewKafkaStatusSender(brokers []string, topic string) (*statusSender, error) {
	producer, err := kafka.NewSyncProducer(brokers)
	if err != nil {
		return nil, err
	}

	return &statusSender{
		producer: producer,
		topic:    topic,
	}, nil
}

func (s *statusSender) Send(orderID model.OrderID, status model.OrderStatus) error {
	return s.producer.Produce(s.topic, fmt.Sprint(orderID), []byte(status))
}

func (s *statusSender) Close() error {
	return s.producer.Close()
}
