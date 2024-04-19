package orderstorage

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/jackc/pgx/v5"
	"route256.ozon.ru/project/loms/internal/model"
	"route256.ozon.ru/project/loms/internal/repository/kafka_storage"
)

func (r *repository) insertOutboxOrderStatusChanged(ctx context.Context, tx pgx.Tx, orderID int64, status string) error {
	order := &model.OrderChangeStatusMessageOrder{
		ID:     orderID,
		Status: status,
	}

	json, err := json.Marshal(order)
	if err != nil {
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
		return err
	}

	return nil
}
