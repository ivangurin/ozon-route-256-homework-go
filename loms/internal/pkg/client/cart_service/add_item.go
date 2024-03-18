package cartservice

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"route256.ozon.ru/project/loms/internal/model"
	"route256.ozon.ru/project/loms/internal/pkg/logger"
)

func (c *client) AddItem(ctx context.Context, UserID, SkuID int64, Quantity uint16) error {

	logger.Info(fmt.Sprintf("cartService.AddItem: start add item %d/%d/%d", UserID, SkuID, Quantity))
	defer logger.Info(fmt.Sprintf("cartService.AddItem: start add item %d/%d/%d", UserID, SkuID, Quantity))

	body := &AddItemRequestBody{
		Count: Quantity,
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		logger.Error("cartService.AddItem: failed to marshal request body", err)
		return fmt.Errorf("cartService.AddItem: failed to marshal request body: %w", err)
	}

	httpReq, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://%s/user/%d/cart/%d", c.host, UserID, SkuID), bytes.NewBuffer(jsonBody))
	if err != nil {
		logger.Error("cartService.AddItem: failed to create request", err)
		return fmt.Errorf("cartService.AddItem: failed to create request: %w", err)
	}

	httpReq = httpReq.WithContext(ctx)
	client := http.DefaultClient
	httpResp, err := client.Do(httpReq)
	if err != nil {
		logger.Error("cartService.AddItem: failed to do request", err)
		return fmt.Errorf("cartService.AddItem: failed to do request: %w", err)
	}
	defer httpResp.Body.Close()

	respBody, err := io.ReadAll(httpResp.Body)
	if err != nil {
		logger.Error("cartService.AddItem: failed to get response body", err)
		return fmt.Errorf("failed to get esponse body: %w", err)
	}

	if httpResp.StatusCode == http.StatusNotFound {
		logger.Error(fmt.Sprintf("cartService.AddItem: %s", string(respBody)), nil)
		return model.ErrNotFound
	} else if httpResp.StatusCode != http.StatusOK {
		logger.Error(fmt.Sprintf("cartService.AddItem: %s", string(respBody)), nil)
		return fmt.Errorf("cartService.AddItem: %s", string(respBody))
	}

	return nil

}
