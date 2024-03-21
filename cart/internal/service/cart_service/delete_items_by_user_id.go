package cartservice

import (
	"context"
	"fmt"

	"route256.ozon.ru/project/cart/internal/pkg/logger"
)

func (s *service) DeleteItemsByUserID(ctx context.Context, userID int64) error {
	err := s.cartStorage.DeleteItemsByUserID(ctx, userID)
	if err != nil {
		logger.Error("cartService.DeleteItemsByUserID: failed to delete items by user id", err)
		return fmt.Errorf("failed to delete items by user id: %w", err)
	}

	return nil
}
