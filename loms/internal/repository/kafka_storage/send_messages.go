package kafka_storage

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"route256.ozon.ru/project/loms/internal/repository/kafka_storage/sqlc"
)

func (r *repository) SendMessages(ctx context.Context, callback func(ctx context.Context, message *sqlc.KafkaOutbox) error) error {

	pool := r.dbClient.GetWriterPool()

	err := pool.BeginFunc(ctx, func(tx pgx.Tx) error {
		qtx := sqlc.New(pool).WithTx(tx)

		var err error
		messages, err := qtx.SelectOutboxMessages(ctx, pgtype.Text{String: StatusNew, Valid: true})
		if err != nil {
			return fmt.Errorf("failed to select outbox messages: %w", err)
		}

		for _, message := range messages {

			err = callback(ctx, &message)
			if err == nil {
				err = UpdateOutboxMessageStatusTx(ctx, tx, message.ID, StatusSent, sql.NullString{})
			} else {
				err = UpdateOutboxMessageStatusTx(ctx, tx, message.ID, StatusFailed, sql.NullString{String: err.Error(), Valid: true})
			}

			if err != nil {
				return fmt.Errorf("failed to update outbox message status: %w", err)
			}

		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
