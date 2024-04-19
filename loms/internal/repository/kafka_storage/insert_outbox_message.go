package kafka_storage

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

func InsertOutboxMessageTx(ctx context.Context, tx pgx.Tx, message *Outbox) error {
	builder := squirrel.
		Insert(KafkaOutboxTable).
		Columns("event", "entity_type", "entity_id", "data").
		Values(message.Event, message.EntityType, message.EntityID, message.Data).
		PlaceholderFormat(squirrel.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to insert outbox message with query %s, args %+v: %w", query, args, err)
	}

	return nil
}
