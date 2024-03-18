package cartservice

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"route256.ozon.ru/project/loms/internal/pkg/logger"
)

func (c *client) DeleteItem(ctx context.Context, UserID, SkuID int64) error {

	logger.Info(fmt.Sprintf("cartService.DeleteItem: start delete item %d/%d", UserID, SkuID))
	defer logger.Info(fmt.Sprintf("cartService.DeleteItem: finish delete item %d/%d", UserID, SkuID))

	httpReq, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("http://%s/user/%d/cart/%d", c.host, UserID, SkuID), nil)
	if err != nil {
		logger.Error("cartService.AddItem: failed to create request", err)
		return fmt.Errorf("cartService.AddItem: failed to create request: %w", err)
	}

	httpReq = httpReq.WithContext(ctx)
	client := http.DefaultClient
	httpResp, err := client.Do(httpReq)
	if err != nil {
		logger.Error("cartService.DeleteItem: failed to do request", err)
		return fmt.Errorf("cartService.DeleteItem: failed to do request: %w", err)
	}
	defer httpResp.Body.Close()

	respBody, err := io.ReadAll(httpResp.Body)
	if err != nil {
		logger.Error("cartService.DeleteItem: failed to get response body", err)
		return fmt.Errorf("cartService.DeleteItem: failed to get response body: %w", err)
	}

	if httpResp.StatusCode != http.StatusNoContent {
		logger.Error(fmt.Sprintf("cartService.DeleteItem: %s", string(respBody)), nil)
		return fmt.Errorf("cartService.AddItem: %s", string(respBody))
	}

	return nil

}
