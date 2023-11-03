package kafka

import (
	"context"
	"fmt"
	"github.com/IBM/sarama"
	"time"
)

type ConsumerGroup struct {
	consumerGroup sarama.ConsumerGroup
	handler       sarama.ConsumerGroupHandler
	topics        []string
}

func NewConsumerGroup(brokers []string, groupID string, topics []string, consumerGroupHandler sarama.ConsumerGroupHandler, opts ...Option) (*ConsumerGroup, error) {
	c := sarama.NewConfig()

	c.Version = sarama.MaxVersion

	c.Consumer.Offsets.Initial = sarama.OffsetOldest
	c.Consumer.Group.ResetInvalidOffsets = true
	c.Consumer.Group.Heartbeat.Interval = 3 * time.Second
	c.Consumer.Group.Session.Timeout = 60 * time.Second
	c.Consumer.Group.Rebalance.Timeout = 60 * time.Second

	c.Consumer.Return.Errors = true

	c.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategyRoundRobin()}

	for _, opt := range opts {
		if err := opt.Apply(c); err != nil {
			return nil, err
		}
	}

	consumerGroup, err := sarama.NewConsumerGroup(brokers, groupID, c)
	if err != nil {
		return nil, err
	}

	return &ConsumerGroup{
		consumerGroup: consumerGroup,
		handler:       consumerGroupHandler,
		topics:        topics,
	}, nil
}

func (cg *ConsumerGroup) Run(ctx context.Context) {
	go func() {
		for err := range cg.consumerGroup.Errors() {
			fmt.Println("Consumer error: ", err)
		}
	}()

	for {
		if err := cg.consumerGroup.Consume(ctx, cg.topics, cg.handler); err != nil {
			fmt.Println("Consumer error: ", err)
		}

		if ctx.Err() != nil {
			return
		}
	}
}

func (cg *ConsumerGroup) Close() error {
	return cg.consumerGroup.Close()
}
