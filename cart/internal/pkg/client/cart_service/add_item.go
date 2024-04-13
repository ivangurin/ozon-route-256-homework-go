package cartservice

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"route256.ozon.ru/project/cart/internal/model"
	"route256.ozon.ru/project/cart/internal/pkg/logger"
)

func (c *client) AddItem(ctx context.Context, UserID, SkuID int64, Quantity uint16) error {
	logger.Infof(ctx, "cartService.AddItem: start add item %d/%d/%d", UserID, SkuID, Quantity)
	defer logger.Infof(ctx, "cartService.AddItem: start add item %d/%d/%d", UserID, SkuID, Quantity)

	body := &AddItemRequestBody{
		Count: Quantity,
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		logger.Errorf(ctx, "cartService.AddItem: failed to marshal request body: %v", err)
		return fmt.Errorf("cartService.AddItem: failed to marshal request body: %w", err)
	}

	httpReq, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://%s/user/%d/cart/%d", c.host, UserID, SkuID), bytes.NewBuffer(jsonBody))
	if err != nil {
		logger.Errorf(ctx, "cartService.AddItem: failed to create request: %v", err)
		return fmt.Errorf("cartService.AddItem: failed to create request: %w", err)
	}

	httpReq = httpReq.WithContext(ctx)
	client := http.DefaultClient
	httpResp, err := client.Do(httpReq)
	if err != nil {
		logger.Errorf(ctx, "cartService.AddItem: failed to do request: %v", err)
		return fmt.Errorf("cartService.AddItem: failed to do request: %w", err)
	}
	defer httpResp.Body.Close()

	respBody, err := io.ReadAll(httpResp.Body)
	if err != nil {
		logger.Errorf(ctx, "cartService.AddItem: failed to get response body: %v", err)
		return fmt.Errorf("failed to get esponse body: %w", err)
	}

	if httpResp.StatusCode == http.StatusNotFound {
		logger.Errorf(ctx, "cartService.AddItem: %s", string(respBody))
		return model.ErrNotFound
	} else if httpResp.StatusCode != http.StatusOK {
		logger.Errorf(ctx, "cartService.AddItem: %s", string(respBody))
		return fmt.Errorf("cartService.AddItem: %s", string(respBody))
	}

	return nil
}
