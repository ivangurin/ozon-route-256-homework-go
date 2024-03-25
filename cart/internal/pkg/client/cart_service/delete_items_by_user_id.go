package cartservice

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"route256.ozon.ru/project/cart/internal/pkg/logger"
)

func (c *client) DeleteItemsByUserID(ctx context.Context, UserID int64) error {
	logger.Infof(ctx, "cartService.DeleteItemsByUserID: start delete items %d", UserID)
	defer logger.Infof(ctx, "cartService.DeleteItemsByUserID: finish delete items %d", UserID)

	httpReq, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("http://%s/user/%d/cart", c.host, UserID), nil)
	if err != nil {
		logger.Errorf(ctx, "cartService.DeleteItemsByUserID: failed to create request: %v", err)
		return fmt.Errorf("cartService.DeleteItemsByUserID: failed to create request: %w", err)
	}

	httpReq = httpReq.WithContext(ctx)
	client := http.DefaultClient
	httpResp, err := client.Do(httpReq)
	if err != nil {
		logger.Errorf(ctx, "cartService.DeleteItemsByUserID: failed to do request: %v", err)
		return fmt.Errorf("cartService.DeleteItemsByUserID: failed to do request: %w", err)
	}
	defer httpResp.Body.Close()

	respBody, err := io.ReadAll(httpResp.Body)
	if err != nil {
		logger.Errorf(ctx, "cartService.DeleteItemsByUserID: failed to get response body: %v", err)
		return fmt.Errorf("cartService.DeleteItemsByUserID: failed to get response body: %w", err)
	}

	if httpResp.StatusCode != http.StatusNoContent {
		logger.Errorf(ctx, "cartService.DeleteItemsByUserID: %s", string(respBody))
		return fmt.Errorf("cartService.DeleteItemsByUserID: %s", string(respBody))
	}

	return nil

}
