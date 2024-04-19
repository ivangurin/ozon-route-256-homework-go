package lomsservice

import (
	"context"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	stock_api "route256.ozon.ru/project/cart/internal/pb/api/stock/v1"
	"route256.ozon.ru/project/cart/internal/pkg/metrics"
)

func (c *Client) StockInfo(ctx context.Context, sku int64) (uint16, error) {
	metrics.UpdateExternalRequestsTotal(
		ServiceName,
		"StockInfo",
	)
	defer metrics.UpdateExternalResponseTime(time.Now().UTC())

	req := &stock_api.StockInfoRequest{
		Sku: sku,
	}

	resp, err := c.StockAPI.Info(ctx, req)
	if err != nil {
		st, _ := status.FromError(err)
		metrics.UpdateExternalResponseCode(
			ServiceName,
			"GetProduct",
			st.Code().String(),
		)
		return 0, err
	}

	metrics.UpdateExternalResponseCode(
		ServiceName,
		"GetProduct",
		codes.OK.String(),
	)

	return uint16(resp.GetCount()), nil
}
