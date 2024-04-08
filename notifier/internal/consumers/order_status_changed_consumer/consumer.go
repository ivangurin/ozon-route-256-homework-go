package order_status_changed_consumer

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/IBM/sarama"
	orderservice "route256.ozon.ru/project/notifier/internal/app/notifier_service/model/order_service"
	"route256.ozon.ru/project/notifier/internal/pkg/logger"
	notifierservice "route256.ozon.ru/project/notifier/internal/service/notifier_service"
)

type Consumer interface {
	Handle(ctx context.Context, msg *sarama.ConsumerMessage) (bool, error)
}

type consumer struct {
	notifierService notifierservice.Service
}

func NewConsumer(
	notifierService notifierservice.Service,
) Consumer {
	return &consumer{
		notifierService: notifierService,
	}
}

func (c *consumer) Handle(ctx context.Context, msg *sarama.ConsumerMessage) (bool, error) {
	logger.Infof("Got a new message. Offset: %d. Partition: %d", msg.Offset, msg.Partition)
	defer logger.Infof("The message is handled. Offset: %d. Partition: %d", msg.Offset, msg.Partition)

	genericMessage := &orderservice.GenericMessage{}
	err := json.Unmarshal(msg.Value, genericMessage)
	if err != nil {
		logger.Errorf("failed to unmarshal the message: %v", err)
		return false, fmt.Errorf("failed to unmarshal the message: %w", err)
	}

	switch genericMessage.Event {
	case orderservice.OrderEventStatusChanged:
		return c.handleOrderStatusChanged(ctx, msg)
	default:
		logger.Errorf("unknown event: %s", genericMessage.Event)
		return false, fmt.Errorf("unknown event: %s", genericMessage.Event)
	}
}

func (c *consumer) handleOrderStatusChanged(ctx context.Context, msg *sarama.ConsumerMessage) (bool, error) {
	orderStatusChangedMessage := &orderservice.OrderChangeStatusMessage{}
	err := json.Unmarshal(msg.Value, orderStatusChangedMessage)
	if err != nil {
		logger.Errorf("failed to unmarshal the message: %v", err)
		return false, fmt.Errorf("failed to unmarshal the message: %w", err)
	}

	err = c.notifierService.OrderStatusChanges(ctx, orderStatusChangedMessage.Data.OrderID, orderStatusChangedMessage.Data.Status)
	if err != nil {
		return false, fmt.Errorf("failed to handle the order status change: %w", err)
	}

	return true, nil
}
