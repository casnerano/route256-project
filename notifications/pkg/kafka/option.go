package kafka

import "github.com/IBM/sarama"

type Option interface {
	Apply(*sarama.Config) error
}

type fnOption func(*sarama.Config) error

func (fn fnOption) Apply(config *sarama.Config) error {
	return fn(config)
}
