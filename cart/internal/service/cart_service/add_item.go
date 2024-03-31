package cartservice

import (
	"context"
	"fmt"

	"route256.ozon.ru/project/cart/internal/model"
	"route256.ozon.ru/project/cart/internal/pkg/logger"
)

func (s *service) AddItem(ctx context.Context, userID int64, skuID int64, quantity uint16) error {
	_, err := s.productService.GetProductWithRetries(ctx, skuID)
	if err != nil {
		logger.Errorf("cartService.AddItem: failed to get product: %v", err)
		return fmt.Errorf("failed to get product: %w", err)
	}

	freeStock, err := s.lomsService.StockInfo(ctx, skuID)
	if err != nil {
		logger.Errorf("lomsService.StockInfo: failed to get free stock: %v", err)
		return fmt.Errorf("failed to get free stock: %w", err)
	}
	if freeStock < quantity {
		return fmt.Errorf("insufficient stock for %d: %w", skuID, model.ErrInsufficientSock)
	}

	err = s.cartStorage.AddItem(ctx, userID, skuID, quantity)
	if err != nil {
		logger.Errorf("cartService.AddItem: failed to add item: %v", err)
		return fmt.Errorf("failed to add item: %w", err)
	}

	return nil
}
