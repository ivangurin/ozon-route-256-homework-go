package cartservice

import (
	"context"
	"fmt"

	"route256.ozon.ru/project/cart/internal/pkg/logger"
	"route256.ozon.ru/project/cart/internal/pkg/tracer"
)

func (s *service) DeleteItem(ctx context.Context, userID int64, skuID int64) error {
	ctx, span := tracer.StartSpanFromContext(ctx, "cartService.DeleteItem")
	defer span.End()

	err := s.cartStorage.DeleteItem(ctx, userID, skuID)
	if err != nil {
		logger.Errorf(ctx, "cartService.DeleteItem:: failed to delete item: %v", err)
		return fmt.Errorf("failed to delete item: %w", err)
	}

	return nil
}
