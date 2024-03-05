package cartservice

import (
	"context"
	"fmt"

	"route256.ozon.ru/project/cart/internal/pkg/logger"
)

func (s *cartService) DeleteItem(ctx context.Context, userID int64, skuID int64) error {
	logger.Info(fmt.Sprintf("cartService.DeleteItem: start delete item from cart userID: %d, skuID: %d", userID, skuID))
	defer logger.Info("cartService.DeleteItem: finish delete item from cart userID")

	err := s.cartStorage.DeleteItem(ctx, userID, skuID)
	if err != nil {
		logger.Error("cartService.DeleteItem:: faild to delete item", err)
		return fmt.Errorf("faild to delete item: %w", err)
	}

	return nil
}
