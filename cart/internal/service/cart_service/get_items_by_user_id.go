package cartservice

import (
	"context"
	"fmt"
	"sync"

	productservice "route256.ozon.ru/project/cart/internal/pkg/client/product_service"
	"route256.ozon.ru/project/cart/internal/pkg/errgroup"
	"route256.ozon.ru/project/cart/internal/pkg/logger"
)

func (s *service) GetItemsByUserID(ctx context.Context, userID int64) (*Cart, error) {
	cart, err := s.cartStorage.GetItemsByUserID(ctx, userID)
	if err != nil {
		logger.Errorf(ctx, "cartService.GetItemsByUserID: failed to get items by user id: %v", err)
		return nil, fmt.Errorf("failed to get items by user id: %w", err)
	}

	mu := &sync.Mutex{}
	products := make(map[int64]*productservice.GetProductResponse)

	eg, ctx := errgroup.NewErrGroup(ctx, 10)

	for sku, _ := range cart.Items {
		eg.Go(func() error {
			product, err := s.productService.GetProductWithRetries(ctx, sku)
			if err != nil {
				return fmt.Errorf("failed to get product %d: %w", sku, err)
			}
			mu.Lock()
			defer mu.Unlock()
			products[sku] = product
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		return nil, err
	}

	resp, err := s.toGetCartResponse(ctx, cart, products)
	if err != nil {
		logger.Errorf(ctx, "cartService.GetItemsByUserID: failed to make response: %v", err)
		return nil, fmt.Errorf("failed to make response: %w", err)
	}

	return resp, nil
}
