package kafka_storage

import (
	"context"
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"route256.ozon.ru/project/loms/internal/pkg/metrics"
)

func InsertOutboxMessageTx(ctx context.Context, tx pgx.Tx, message *Outbox) error {
	metrics.UpdateDatabaseRequestsTotal(
		RepositoryName,
		"InsertOutboxMessageTx",
		"insert",
	)

	defer metrics.UpdateDatabaseResponseTime(time.Now().UTC())

	builder := squirrel.
		Insert(KafkaOutboxTable).
		Columns("event", "entity_type", "entity_id", "data").
		Values(message.Event, message.EntityType, message.EntityID, message.Data).
		PlaceholderFormat(squirrel.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		metrics.UpdateDatabaseResponseCode(
			RepositoryName,
			"InsertOutboxMessageTx",
			"insert",
			"error",
		)
		return err
	}

	_, err = tx.Exec(ctx, query, args...)
	if err != nil {
		metrics.UpdateDatabaseResponseCode(
			RepositoryName,
			"InsertOutboxMessageTx",
			"insert",
			"error",
		)
		return fmt.Errorf("failed to insert outbox message with query %s, args %+v: %w", query, args, err)
	}

	metrics.UpdateDatabaseResponseCode(
		RepositoryName,
		"InsertOutboxMessageTx",
		"insert",
		"ok",
	)
	return nil
}
