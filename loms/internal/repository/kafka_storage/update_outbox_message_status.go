package kafka_storage

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"route256.ozon.ru/project/loms/internal/pkg/metrics"
)

func UpdateOutboxMessageStatusTx(ctx context.Context, tx pgx.Tx, id string, status string, error sql.NullString) error {
	metrics.UpdateDatabaseRequestsTotal(
		RepositoryName,
		"UpdateOutboxMessageStatusTx",
		"update",
	)

	defer metrics.UpdateDatabaseResponseTime(time.Now().UTC())
	
	builder := squirrel.
		Update(KafkaOutboxTable).
		Set("status", status).
		Set("error", error).
		Set("updated_at", time.Now().UTC()).
		Where(squirrel.Eq{"id": id}).
		PlaceholderFormat(squirrel.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		metrics.UpdateDatabaseResponseCode(
			RepositoryName,
			"UpdateOutboxMessageStatusTx",
			"update",
			"error",
		)
		return err
	}

	_, err = tx.Exec(ctx, query, args...)
	if err != nil {
		metrics.UpdateDatabaseResponseCode(
			RepositoryName,
			"UpdateOutboxMessageStatusTx",
			"update",
			"error",
		)
		return fmt.Errorf("failed to update outbox message status with query %s, args %+v: %w", query, args, err)
	}

	metrics.UpdateDatabaseResponseCode(
		RepositoryName,
		"UpdateOutboxMessageStatusTx",
		"update",
		"ok",
	)
	return nil
}
