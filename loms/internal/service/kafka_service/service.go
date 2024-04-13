package kafka_service

import (
	"context"
	"sync"

	"route256.ozon.ru/project/loms/internal/pkg/kafka"
	"route256.ozon.ru/project/loms/internal/repository/kafka_storage"
)

type Service interface {
	SendMessages(ctx context.Context)
	StopSendMessages() error
}

type service struct {
	kafkaStorage    kafka_storage.Repository
	kafkaProducer   kafka.Producer
	sendMessagesWG  sync.WaitGroup
	sendMessageDone chan struct{}
}

func NewService(
	kafkaStorage kafka_storage.Repository,
	kafkaProducer kafka.Producer,
) Service {
	return &service{
		kafkaStorage:    kafkaStorage,
		kafkaProducer:   kafkaProducer,
		sendMessagesWG:  sync.WaitGroup{},
		sendMessageDone: make(chan struct{}),
	}
}
