package orderstorage

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"route256.ozon.ru/project/loms/internal/model"
	"route256.ozon.ru/project/loms/internal/pkg/metrics"
	"route256.ozon.ru/project/loms/internal/pkg/tracer"
	"route256.ozon.ru/project/loms/internal/repository/kafka_storage"
)

func (r *repository) insertOutboxOrderStatusChanged(ctx context.Context, tx pgx.Tx, orderID int64, status string) error {
	ctx, span := tracer.StartSpanFromContext(ctx, "orderStorage:insertOutboxOrderStatusChanged")
	defer span.End()

	metrics.UpdateDatabaseRequestsTotal(
		RepositoryName,
		"insertOutboxOrderStatusChanged",
		"insert",
	)
	defer metrics.UpdateDatabaseResponseTime(time.Now().UTC())

	order := &model.OrderChangeStatusMessageOrder{
		ID:     orderID,
		Status: status,
	}

	json, err := json.Marshal(order)
	if err != nil {
		metrics.UpdateDatabaseResponseCode(
			RepositoryName,
			"insertOutboxOrderStatusChanged",
			"insert",
			"error",
		)
		return err
	}

	outbox := &kafka_storage.Outbox{
		Event:      model.EventOrderStatusChanged,
		EntityType: model.EntityTypeOrder,
		EntityID:   fmt.Sprintf("%d", orderID),
		Data:       string(json),
	}

	err = kafka_storage.InsertOutboxMessageTx(ctx, tx, outbox)
	if err != nil {
		metrics.UpdateDatabaseResponseCode(
			RepositoryName,
			"insertOutboxOrderStatusChanged",
			"insert",
			"error",
		)
		return err
	}

	metrics.UpdateDatabaseResponseCode(
		RepositoryName,
		"insertOutboxOrderStatusChanged",
		"insert",
		"ok",
	)
	return nil
}
