package lomsservice

import (
	"context"

	stock_api "route256.ozon.ru/project/cart/internal/pb/api/stock/v1"
)

func (c *Client) StockInfo(ctx context.Context, sku int64) (uint16, error) {
	req := &stock_api.StockInfoRequest{
		Sku: sku,
	}

	resp, err := c.StockAPI.Info(ctx, req)
	if err != nil {
		return 0, err
	}

	return uint16(resp.GetCount()), nil
}
