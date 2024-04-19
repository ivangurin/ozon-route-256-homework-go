package lomsservice

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	stock_api "route256.ozon.ru/project/cart/internal/pb/api/stock/v1"
	"route256.ozon.ru/project/cart/internal/pkg/metrics"
	"route256.ozon.ru/project/cart/internal/pkg/tracer"
)

func (c *Client) StockInfo(ctx context.Context, sku int64) (uint16, error) {
	metrics.UpdateExternalRequestsTotal(
		ServiceName,
		"StockInfo",
	)
	defer metrics.UpdateExternalResponseTime(time.Now().UTC())

	fmt.Println("TRACE ID", tracer.GetTraceID(ctx))
	ctx = metadata.AppendToOutgoingContext(ctx, "x-trace-id", tracer.GetTraceID(ctx))

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
