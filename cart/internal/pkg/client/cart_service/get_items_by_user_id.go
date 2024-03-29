package cartservice

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"route256.ozon.ru/project/cart/internal/model"
	"route256.ozon.ru/project/cart/internal/pkg/logger"
)

func (c *client) GetItemsByUserID(ctx context.Context, UserID int64) (*GetItmesByUserIDResponse, error) {

	logger.Info(fmt.Sprintf("cartService.GetItemsByUserID: start get items %d", UserID))
	defer logger.Info(fmt.Sprintf("cartService.GetItemsByUserID: finish get items %d", UserID))

	httpReq, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://%s/user/%d/cart", c.host, UserID), nil)
	if err != nil {
		logger.Error("cartService.GetItemsByUserID: failed to create request", err)
		return nil, fmt.Errorf("cartService.GetItemsByUserID: failed to create request: %w", err)
	}

	httpReq = httpReq.WithContext(ctx)
	client := http.DefaultClient
	httpResp, err := client.Do(httpReq)
	if err != nil {
		logger.Error("cartService.GetItemsByUserID: failed to do request", err)
		return nil, fmt.Errorf("cartService.GetItemsByUserID: failed to do request: %w", err)
	}
	defer httpResp.Body.Close()

	respBody, err := io.ReadAll(httpResp.Body)
	if err != nil {
		logger.Error("cartService.GetItemsByUserID: failed to get response body", err)
		return nil, fmt.Errorf("cartService.GetItemsByUserID: failed to get response body: %w", err)
	}

	if httpResp.StatusCode == http.StatusNotFound {
		logger.Error(fmt.Sprintf("cartService.GetItemsByUserID: %s", string(respBody)), nil)
		return nil, model.ErrNotFound
	}

	if httpResp.StatusCode != http.StatusOK {
		logger.Error(fmt.Sprintf("cartService.GetItemsByUserID: %s", string(respBody)), nil)
		return nil, fmt.Errorf("cartService.GetItemsByUserID: %s", string(respBody))
	}

	resp := &GetItmesByUserIDResponse{}
	err = json.Unmarshal(respBody, resp)
	if err != nil {
		logger.Error("cartService.GetItemsByUserID: failed to unmarshal body", err)
		return nil, fmt.Errorf("cartService.GetItemsByUserID: failed to unmarshal body: %w", err)
	}

	return resp, nil

}
