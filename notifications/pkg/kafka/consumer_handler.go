package kafka

import (
	"fmt"
	"github.com/IBM/sarama"
)

var _ sarama.ConsumerGroupHandler = (*ConsumerGroupHandler)(nil)

type ConsumerGroupHandler struct {
	ready chan bool
}

func NewConsumerGroupHandler() *ConsumerGroupHandler {
	return &ConsumerGroupHandler{
		ready: make(chan bool),
	}
}

func (h *ConsumerGroupHandler) Ready() <-chan bool {
	return h.ready
}

func (h *ConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error {
	close(h.ready)

	return nil
}

func (h *ConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (h *ConsumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case message := <-claim.Messages():
			fmt.Printf("Order #%s recived \"%s\" status.\n",
				string(message.Key),
				string(message.Value),
			)

			session.MarkMessage(message, "")
		case <-session.Context().Done():
			return nil
		}
	}
}
