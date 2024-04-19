package kafka

import (
	"time"

	"github.com/IBM/sarama"
)

type ProducerOption func(c *sarama.Config)

// WithProducerPartitioner ...
func WithProducerPartitioner(pfn sarama.PartitionerConstructor) ProducerOption {
	return func(c *sarama.Config) {
		c.Producer.Partitioner = pfn
	}
}

// WithRequiredAcks ...
func WithRequiredAcks(acks sarama.RequiredAcks) ProducerOption {
	return func(c *sarama.Config) {
		c.Producer.RequiredAcks = acks
	}
}

// WithIdempotent ...
func WithIdempotent() ProducerOption {
	return func(c *sarama.Config) {
		c.Producer.Idempotent = true
	}
}

// WithMaxRetries ...
func WithMaxRetries(n int) ProducerOption {
	return func(c *sarama.Config) {
		c.Producer.Retry.Max = n
	}
}

// WithRetryBackoff ...
func WithRetryBackoff(d time.Duration) ProducerOption {
	return func(c *sarama.Config) {
		c.Producer.Retry.Backoff = d
	}
}

// WithMaxOpenRequests ...
func WithMaxOpenRequests(n int) ProducerOption {
	return func(c *sarama.Config) {
		c.Net.MaxOpenRequests = n
	}
}

// WithProducerFlushMessages ...
func WithProducerFlushMessages(n int) ProducerOption {
	return func(c *sarama.Config) {
		c.Producer.Flush.Messages = n
	}
}

// WithProducerFlushFrequency ...
func WithProducerFlushFrequency(d time.Duration) ProducerOption {
	return func(c *sarama.Config) {
		c.Producer.Flush.Frequency = d
	}
}
