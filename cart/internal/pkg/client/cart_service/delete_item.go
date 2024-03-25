package cartservice

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"route256.ozon.ru/project/cart/internal/pkg/logger"
)

func (c *client) DeleteItem(ctx context.Context, UserID, SkuID int64) error {
	logger.Infof(ctx, "cartService.DeleteItem: start delete item %d/%d", UserID, SkuID)
	defer logger.Infof(ctx, "cartService.DeleteItem: finish delete item %d/%d", UserID, SkuID)

	httpReq, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("http://%s/user/%d/cart/%d", c.host, UserID, SkuID), nil)
	if err != nil {
		logger.Errorf(ctx, "cartService.AddItem: failed to create request: %v", err)
		return fmt.Errorf("cartService.AddItem: failed to create request: %w", err)
	}

	httpReq = httpReq.WithContext(ctx)
	client := http.DefaultClient
	httpResp, err := client.Do(httpReq)
	if err != nil {
		logger.Errorf(ctx, "cartService.DeleteItem: failed to do request: %v", err)
		return fmt.Errorf("cartService.DeleteItem: failed to do request: %w", err)
	}
	defer httpResp.Body.Close()

	respBody, err := io.ReadAll(httpResp.Body)
	if err != nil {
		logger.Errorf(ctx, "cartService.DeleteItem: failed to get response body: %v", err)
		return fmt.Errorf("cartService.DeleteItem: failed to get response body: %w", err)
	}

	if httpResp.StatusCode != http.StatusNoContent {
		logger.Errorf(ctx, "cartService.DeleteItem: %s", string(respBody))
		return fmt.Errorf("cartService.AddItem: %s", string(respBody))
	}

	return nil
}
