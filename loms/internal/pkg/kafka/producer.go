package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/IBM/sarama"
	"route256.ozon.ru/project/loms/internal/config"
	"route256.ozon.ru/project/loms/internal/pkg/logger"
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
				Key:   []byte(appParam),
				Value: []byte(config.AppName),
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
	err := p.syncProducer.Close()
	if err != nil {
		logger.Errorf("failed to close kafka producer: %v", err)
		return fmt.Errorf("failed to close kafka producer: %w", err)
	}
	logger.Info("kafka producer closed successfully")
	return nil
}
