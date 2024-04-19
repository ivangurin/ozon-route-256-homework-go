package kafka_storage

import (
	"context"

	"route256.ozon.ru/project/loms/internal/db"
	"route256.ozon.ru/project/loms/internal/repository/kafka_storage/sqlc"
)

const KafkaOutboxTable = "kafka_outbox"

type Repository interface {
	SendMessages(ctx context.Context, callback func(ctx context.Context, message *sqlc.KafkaOutbox) error) error
}

type repository struct {
	ctx      context.Context
	dbClient db.Client
}

func NewRepository(ctx context.Context, dbClient db.Client) Repository {
	return &repository{
		ctx:      ctx,
		dbClient: dbClient,
	}
}
