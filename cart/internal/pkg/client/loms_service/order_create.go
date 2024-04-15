package lomsservice

import (
	"context"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"route256.ozon.ru/project/cart/internal/pkg/metrics"
)

func (c *Client) OrderCreate(ctx context.Context, user int64, items OrderItems) (int64, error) {
	metrics.UpdateExternalRequestsTotal(
		ServiceName,
		"OrderCreate",
	)
	defer metrics.UpdateExternalResponseTime(time.Now().UTC())

	req := ToOrderCreateRequest(user, items)

	resp, err := c.OrderAPI.Create(ctx, req)
	if err != nil {
		st, _ := status.FromError(err)
		metrics.UpdateExternalResponseCode(
			ServiceName,
			"OrderCreate",
			st.Code().String(),
		)
		return 0, err
	}

	metrics.UpdateExternalResponseCode(
		ServiceName,
		"OrderCreate",
		codes.OK.String(),
	)

	return resp.GetOrderId(), nil
}
