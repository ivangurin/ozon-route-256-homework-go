package cartservice

import (
	"context"
	"fmt"

	"route256.ozon.ru/project/cart/internal/pkg/logger"
)

func (s *cartService) AddItem(ctx context.Context, userID int64, skuID int64, quantity uint16) error {
	logger.Info(fmt.Sprintf("cartService.AddItem: start add item to cart userID: %d, skuID: %d, quantity: %v", userID, skuID, quantity))
	defer logger.Info("cartService.AddItem: finish add item to cart")

	_, err := s.productService.GetProductWithRetries(ctx, skuID)
	if err != nil {
		logger.Error("cartService.AddItem: failed to get product", err)
		return fmt.Errorf("failed to get product: %w", err)
	}

	err = s.cartStorage.AddItem(ctx, userID, skuID, quantity)
	if err != nil {
		logger.Error("cartService.AddItem: failed to add item", err)
		return fmt.Errorf("failed to add item: %w", err)
	}

	return nil
}
