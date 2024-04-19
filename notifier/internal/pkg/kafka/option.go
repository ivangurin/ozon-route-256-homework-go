package kafka

import (
	"github.com/IBM/sarama"
)

type ConsumerGroupOption func(c *sarama.Config)

func WithOffsetsInitial(v int64) ConsumerGroupOption {
	return func(c *sarama.Config) {
		c.Consumer.Offsets.Initial = v
	}
}

func WithReturnSuccessesEnabled(isEnabled bool) ConsumerGroupOption {
	return func(c *sarama.Config) {
		c.Producer.Return.Successes = isEnabled
	}
}
