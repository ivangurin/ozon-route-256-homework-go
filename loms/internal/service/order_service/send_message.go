package orderservice

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"route256.ozon.ru/project/loms/internal/config"
	"route256.ozon.ru/project/loms/internal/model"
	"route256.ozon.ru/project/loms/internal/pkg/logger"
)

func (s *service) sendMessageStatusChanged(ctx context.Context, orderID int64, status string) error {
	msg := model.OrderChangeStatusMessage{
		Event:  model.OrderEventStatusChanged,
		Entity: model.OrderEntityOrder,
		ID:     fmt.Sprintf("%d", orderID),
		UUID:   uuid.NewString(),
		Time:   time.Now(),
		Data: model.OrderChangeStatusMessageData{
			OrderID: orderID,
			Status:  status,
		},
	}

	err := s.kafkaProducer.SendMessageWithKey(ctx, config.KafkaOrderEventsTopic, fmt.Sprintf("%d", orderID), msg)
	if err != nil {
		logger.Errorf("failed to send message: %v", err)
		return fmt.Errorf("failed to send message: %w", err)
	}

	return nil
}
