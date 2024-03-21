package lomsservice

import (
	"context"
)

func (c *Client) OrederCreate(ctx context.Context, user int64, items OrderItems) (int64, error) {
	req := toOrederCreateRequest(user, items)

	resp, err := c.OrderAPI.Create(ctx, req)
	if err != nil {
		return 0, err
	}

	return resp.GetOrderId(), nil
}
