package cartservice

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"route256.ozon.ru/project/loms/internal/pkg/logger"
)

func (c *client) DeleteItemsByUserID(ctx context.Context, UserID int64) error {

	logger.Info(fmt.Sprintf("cartService.DeleteItemsByUserID: start delete items %d", UserID))
	defer logger.Info(fmt.Sprintf("cartService.DeleteItemsByUserID: finish delete items %d", UserID))

	httpReq, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("http://%s/user/%d/cart", c.host, UserID), nil)
	if err != nil {
		logger.Error("cartService.DeleteItemsByUserID: failed to create request", err)
		return fmt.Errorf("cartService.DeleteItemsByUserID: failed to create request: %w", err)
	}

	httpReq = httpReq.WithContext(ctx)
	client := http.DefaultClient
	httpResp, err := client.Do(httpReq)
	if err != nil {
		logger.Error("cartService.DeleteItemsByUserID: failed to do request", err)
		return fmt.Errorf("cartService.DeleteItemsByUserID: failed to do request: %w", err)
	}
	defer httpResp.Body.Close()

	respBody, err := io.ReadAll(httpResp.Body)
	if err != nil {
		logger.Error("cartService.DeleteItemsByUserID: failed to get response body", err)
		return fmt.Errorf("cartService.DeleteItemsByUserID: failed to get response body: %w", err)
	}

	if httpResp.StatusCode != http.StatusNoContent {
		logger.Error(fmt.Sprintf("cartService.DeleteItemsByUserID: %s", string(respBody)), nil)
		return fmt.Errorf("cartService.DeleteItemsByUserID: %s", string(respBody))
	}

	return nil

}
