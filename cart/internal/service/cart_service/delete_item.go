package cartservice

import (
	"context"
	"fmt"

	"route256.ozon.ru/project/cart/internal/pkg/logger"
)

func (s *service) DeleteItem(ctx context.Context, userID int64, skuID int64) error {
	err := s.cartStorage.DeleteItem(ctx, userID, skuID)
	if err != nil {
		logger.Error("cartService.DeleteItem:: failed to delete item", err)
		return fmt.Errorf("failed to delete item: %w", err)
	}

	return nil
}
