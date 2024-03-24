package stockapi

import (
	"context"
	"errors"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"route256.ozon.ru/project/loms/internal/model"
	"route256.ozon.ru/project/loms/pkg/api/stock/v1"
)

func (a *API) Info(ctx context.Context, req *stock.StockInfoRequest) (*stock.StockInfoResponse, error) {
	quantity, err := a.stockService.Info(ctx, req.GetSku())
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, status.Errorf(codes.NotFound, fmt.Sprintf("sku %d not found", req.GetSku()))
		} else {
			return nil, status.Errorf(codes.Internal, fmt.Sprintf("internal error: %v", err))
		}
	}

	return &stock.StockInfoResponse{
		Count: uint64(quantity),
	}, nil
}
