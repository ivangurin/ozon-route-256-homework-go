package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/IBM/sarama"
	"route256.ozon.ru/project/loms/internal/config"
	"route256.ozon.ru/project/loms/internal/pkg/logger"
	"route256.ozon.ru/project/loms/internal/pkg/tracer"
)

const (
	appParam = "app-name"
)

type Producer interface {
	SendMessageWithKey(ctx context.Context, topic string, key string, message interface{}) error
	Close() error
}

type producer struct {
	syncProducer sarama.SyncProducer
}

func NewSyncProducer(addr string, opts ...ProducerOption) (Producer, error) {
	config := NewConfig(opts...)

	syncProducer, err := sarama.NewSyncProducer([]string{addr}, config)
	if err != nil {
		return nil, err
	}

	return &producer{
		syncProducer: syncProducer,
	}, nil
}

func (p *producer) SendMessage(topic string, message interface{}) error {
	return nil
}

func (p *producer) SendMessageWithKey(ctx context.Context, topic string, key string, message interface{}) error {
	err := ctx.Err()
	if err != nil {
		return err
	}

	msg, err := json.Marshal(message)
	if err != nil {
		return err
	}

	pm := &sarama.ProducerMessage{
		Headers: []sarama.RecordHeader{
			{
				Key:   sarama.ByteEncoder(appParam),
				Value: sarama.ByteEncoder(config.AppName),
			},
			{
				Key:   sarama.ByteEncoder("x-trace-id"),
				Value: sarama.ByteEncoder(tracer.GetTraceID(ctx)),
			},
		},
		Timestamp: time.Now(),
		Topic:     topic,
		Key:       sarama.StringEncoder(key),
		Value:     sarama.StringEncoder(msg),
	}

	_, _, err = p.syncProducer.SendMessage(pm)
	if err != nil {
		return err
	}

	return nil
}

func (p *producer) Close() error {
	ctx := context.Background()
	err := p.syncProducer.Close()
	if err != nil {
		logger.Errorf(ctx, "failed to close kafka producer: %v", err)
		return fmt.Errorf("failed to close kafka producer: %w", err)
	}
	logger.Info(ctx, "kafka producer is closed successfully")
	return nil
}
