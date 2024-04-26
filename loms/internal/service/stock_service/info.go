package stockservice

import (
	"context"

	"route256.ozon.ru/project/loms/internal/pkg/tracer"
)

func (s *service) Info(ctx context.Context, sku int64) (uint16, error) {
	ctx, span := tracer.StartSpanFromContext(ctx, "stockService.Info")
	defer span.End()

	quantity, err := s.stockStorage.GetBySku(ctx, sku)
	if err != nil {
		return 0, err
	}

	return quantity, nil
}
