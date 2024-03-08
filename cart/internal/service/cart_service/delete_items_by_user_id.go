package cartservice

import (
	"context"
	"fmt"

	"route256.ozon.ru/project/cart/internal/pkg/logger"
)

func (s *cartService) DeleteItemsByUserID(ctx context.Context, userID int64) error {
	logger.Info(fmt.Sprintf("cartService.DeleteItemsByUserID: start delete items by userID: %d", userID))
	defer logger.Info(fmt.Sprintf("cartService.DeleteItemsByUserID: finish delete items by userID: %d", userID))

	err := s.cartStorage.DeleteItemsByUserID(ctx, userID)
	if err != nil {
		logger.Error("cartService.DeleteItemsByUserID: failed to delete items by user id", err)
		return fmt.Errorf("failed to delete items by user id: %w", err)
	}

	return nil
}
