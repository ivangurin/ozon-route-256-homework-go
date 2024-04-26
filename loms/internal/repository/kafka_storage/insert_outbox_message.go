package kafka_storage

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"route256.ozon.ru/project/loms/internal/pkg/metrics"
	"route256.ozon.ru/project/loms/internal/pkg/tracer"
)

func InsertOutboxMessage(ctx context.Context, tx pgx.Tx, message *Outbox) error {
	ctx, span := tracer.StartSpanFromContext(ctx, "kafkaStorage:InsertOutboxMessage")
	defer span.End()

	metrics.UpdateDatabaseRequestsTotal(
		RepositoryName,
		"InsertOutboxMessageTx",
		"insert",
	)
	defer metrics.UpdateDatabaseResponseTime(time.Now().UTC())

	traceID := tracer.GetTraceID(ctx)
	var sqlTraceID sql.NullString
	if traceID != "" {
		sqlTraceID = sql.NullString{
			String: traceID,
			Valid:  true,
		}
	}

	spanID := tracer.GetSpanID(ctx)
	var sqlSpanID sql.NullString
	if spanID != "" {
		sqlSpanID = sql.NullString{
			String: spanID,
			Valid:  true,
		}
	}

	builder := squirrel.
		Insert(KafkaOutboxTable).
		Columns("event", "entity_type", "entity_id", "data", "trace_id", "span_id").
		Values(message.Event, message.EntityType, message.EntityID, message.Data, sqlTraceID, sqlSpanID).
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
