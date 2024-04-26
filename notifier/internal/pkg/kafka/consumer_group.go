package kafka

import (
	"context"
	"fmt"
	"strings"

	"github.com/IBM/sarama"
	"route256.ozon.ru/project/notifier/internal/pkg/logger"
)

type ConsumerGroup interface {
	Run() error
	Close() error
}

type consumerGroup struct {
	ctx                  context.Context
	consumerGroup        sarama.ConsumerGroup
	consumerGroupHandler sarama.ConsumerGroupHandler
	topics               []string
}

func NewConsumerGroup(
	ctx context.Context,
	addrs []string,
	id string,
	topics []string,
	handler Handler,
	opts ...ConsumerGroupOption,
) (ConsumerGroup, error) {
	config := NewConfig(opts...)

	cg, err := sarama.NewConsumerGroup(addrs, id, config)
	if err != nil {
		return nil, err
	}

	return &consumerGroup{
		ctx:                  ctx,
		consumerGroup:        cg,
		consumerGroupHandler: NewConsumerGroupHandler(ctx, handler),
		topics:               topics,
	}, nil
}

func (cg *consumerGroup) Run() error {
	for {
		if err := cg.consumerGroup.Consume(cg.ctx, cg.topics, cg.consumerGroupHandler); err != nil {
			if err != sarama.ErrClosedConsumerGroup {
				logger.Errorf(cg.ctx, "error consume topics %v: %v", strings.Join(cg.topics, ", "), err)
				return fmt.Errorf("error consume topics %v: %w", strings.Join(cg.topics, ", "), err)
			}
		}
	}
}

func (cg *consumerGroup) Close() error {
	err := cg.consumerGroup.Close()
	if err != nil {
		logger.Errorf(cg.ctx, "failed to close consumer group: %v", err)
		return fmt.Errorf("failed to close consumer group: %w", err)
	}
	logger.Info(cg.ctx, "consumer group closed successfully")
	return nil
}
