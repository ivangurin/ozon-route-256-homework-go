package stockservice

import "context"

func (s *service) Info(ctx context.Context, sku int64) (uint16, error) {
	quantity, err := s.stockStorage.GetBySku(ctx, sku)
	if err != nil {
		return 0, err
	}

	return quantity, nil
}
