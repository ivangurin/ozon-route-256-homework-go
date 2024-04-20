package kafka_service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"go.opentelemetry.io/otel/trace"
	"route256.ozon.ru/project/loms/internal/config"
	"route256.ozon.ru/project/loms/internal/model"
	"route256.ozon.ru/project/loms/internal/pkg/logger"
	"route256.ozon.ru/project/loms/internal/pkg/tracer"
	"route256.ozon.ru/project/loms/internal/repository/kafka_storage/sqlc"
)

func (s *service) SendMessages(ctx context.Context) {

	logger.Info(ctx, "kafka outbox sender is starting...")
	s.sendMessagesWG.Add(1)
	go func() {
		s.sendMessages(ctx)
	}()
}

func (s *service) StopSendMessages() error {
	close(s.sendMessageDone)
	s.sendMessagesWG.Wait()
	return nil
}

func (s *service) sendMessages(ctx context.Context) error {
	ticker := time.NewTicker(config.KafkaOutboxSenderTimeout * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-s.sendMessageDone:
			s.sendMessagesWG.Done()
			logger.Info(ctx, "kafka outbox sender is stopped successfully")
			return nil
		case <-ticker.C:
			err := s.kafkaStorage.SendMessages(ctx, s.sendMessage)
			if err != nil {
				logger.Errorf(ctx, "failed to send messages: %v", err)
			}
		}
	}
}

func (s *service) sendMessage(ctx context.Context, message *sqlc.KafkaOutbox) error {
	var err error

	if message.TraceID.String != "" {
		spanContext := trace.SpanContextConfig{
			TraceFlags: trace.FlagsSampled,
			Remote:     true,
		}
		spanContext.TraceID, err = trace.TraceIDFromHex(message.TraceID.String)
		if err != nil {
			return err
		}
		spanContext.SpanID, err = trace.SpanIDFromHex(message.SpanID.String)
		if err != nil {
			return err
		}
		ctx = trace.ContextWithSpanContext(ctx,
			trace.NewSpanContext(spanContext))
	}

	var span trace.Span
	ctx, span = tracer.StartSpanFromContext(ctx, "kafkaService.sendMessage",
		trace.WithSpanKind(trace.SpanKindProducer))
	defer span.End()

	switch message.Event.String {
	case model.EventOrderStatusChanged:
		err = s.sendOrderStatusChangedMessage(ctx, message)
	default:
		logger.Errorf(ctx, "failed to send message: %v", err)
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
		logger.Errorf(ctx, "failed to send OrderChangeStatusMessage: %v", err)
		return fmt.Errorf("failed to send OrderChangeStatusMessage: %w", err)
	}

	return nil
}
