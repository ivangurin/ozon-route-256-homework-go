package cartservice

import (
	"context"
	"fmt"

	"route256.ozon.ru/project/cart/internal/pkg/logger"
	"route256.ozon.ru/project/cart/internal/pkg/tracer"
)

func (s *service) DeleteItemsByUserID(ctx context.Context, userID int64) error {
	ctx, span := tracer.StartSpanFromContext(ctx, "cartService.DeleteItemsByUserID")
	defer span.End()

	err := s.cartStorage.DeleteItemsByUserID(ctx, userID)
	if err != nil {
		logger.Errorf(ctx, "cartService.DeleteItemsByUserID: failed to delete items by user id: %v", err)
		return fmt.Errorf("failed to delete items by user id: %w", err)
	}

	return nil
}
