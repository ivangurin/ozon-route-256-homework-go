package kafka_storage

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

func UpdateOutboxMessageStatusTx(ctx context.Context, tx pgx.Tx, id string, status string, error sql.NullString) error {
	builder := squirrel.
		Update(KafkaOutboxTable).
		Set("status", status).
		Set("error", error).
		Set("updated_at", time.Now().UTC()).
		Where(squirrel.Eq{"id": id}).
		PlaceholderFormat(squirrel.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to update outbox message status with query %s, args %+v: %w", query, args, err)
	}

	return nil
}
