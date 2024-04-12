package kafka_service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"route256.ozon.ru/project/loms/internal/config"
	"route256.ozon.ru/project/loms/internal/model"
	"route256.ozon.ru/project/loms/internal/pkg/logger"
	"route256.ozon.ru/project/loms/internal/repository/kafka_storage/sqlc"
)

func (s *service) SendMessages(ctx context.Context) {
	logger.Info("kafka outbox sender is starting...")
	s.sendMessagesWG.Add(1)
	go func() {
		time.Sleep(10 * time.Second)
		s.sendMessages(ctx)
	}()
}

func (s *service) StopSendMessages() error {
	close(s.sendMessagesDone)
	s.sendMessagesWG.Wait()
	return nil
}

func (s *service) sendMessages(ctx context.Context) error {
	firstTime := true
	for {
		select {
		case <-s.sendMessagesDone:
			s.sendMessagesWG.Done()
			logger.Info("kafka outbox sender is stopped successfully")
			return nil
		default:
			if !firstTime {
				time.Sleep(1 * time.Second)
			}
			err := s.kafkaStorage.SendMessages(ctx, s.sendMessage)
			if err != nil {
				logger.Errorf("failed to send messages: %v", err)
			}
			firstTime = false
		}
	}
}

func (s *service) sendMessage(ctx context.Context, message *sqlc.KafkaOutbox) error {
	var err error
	switch message.Event.String {
	case model.EventOrderStatusChanged:
		err = s.sendOrderStatusChangedMessage(ctx, message)
	default:
		logger.Errorf("failed to send message: %v", err)
	}
	return err
}

func (s *service) sendOrderStatusChangedMessage(ctx context.Context, message *sqlc.KafkaOutbox) error {
	order := &model.OrderChangeStatusMessageOrder{}
	err := json.Unmarshal([]byte(message.Data.String), order)
	if err != nil {
		return fmt.Errorf("failed to unmarshal order data: %w", err)
	}

	orderStatusChangedMessage := model.OrderChangeStatusMessage{
		ID:         message.ID,
		Time:       message.CreatedAt.Time,
		Event:      message.Event.String,
		EntityType: message.EntityType.String,
		EntityID:   message.EntityID.String,
		Data: model.OrderChangeStatusMessageData{
			Order: *order,
		},
	}

	err = s.kafkaProducer.SendMessageWithKey(
		ctx,
		config.KafkaOrderEventsTopic,
		message.EntityID.String,
		orderStatusChangedMessage)
	if err != nil {
		logger.Errorf("failed to send OrderChangeStatusMessage: %v", err)
		return fmt.Errorf("failed to send OrderChangeStatusMessage: %w", err)
	}

	return nil
}