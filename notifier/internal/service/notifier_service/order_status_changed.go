package notifierservice

import (
	"context"

	"route256.ozon.ru/project/notifier/internal/pkg/logger"
	"route256.ozon.ru/project/notifier/internal/pkg/tracer"
)

func (s *service) OrderStatusChanges(ctx context.Context, orderID int64, status string) error {
	ctx, span := tracer.StartSpanFromContext(ctx, "notifierService.OrderStatusChanges")
	defer span.End()

	logger.Infof(ctx, "Got a new order status. OrderID: %d. Status: %s", orderID, status)
	return nil
}
