package productservice

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"route256.ozon.ru/project/cart/internal/config"
	"route256.ozon.ru/project/cart/internal/model"
	"route256.ozon.ru/project/cart/internal/pkg/logger"
	"route256.ozon.ru/project/cart/internal/pkg/metrics"
)

const StatusEnhanceYourCalm = 420

func (c *client) GetProduct(ctx context.Context, skuID int64) (*GetProductResponse, error) {
	resp, exists := productStorage[skuID]
	if exists {
		return resp, nil
	}

	metrics.UpdateExternalRequestsTotal(
		ServiceName,
		"GetProduct",
	)
	defer metrics.UpdateExternalResponseTime(time.Now().UTC())

	req := &GetProductRequest{
		Token: config.ProductServiceToken,
		Sku:   skuID,
	}

	jsonReq, err := json.Marshal(req)
	if err != nil {
		logger.Errorf(ctx, "productService.getProduct: failed to marshal get product request: %v", err)
		return nil, fmt.Errorf("failed to marshal get product request: %w", err)
	}

	httpReq, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://%s/get_product", config.ProductServiceHost), bytes.NewBuffer(jsonReq))
	if err != nil {
		logger.Errorf(ctx, "productService.getProduct: failed to create request: %v", err)
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	httpReq = httpReq.WithContext(ctx)
	client := http.DefaultClient
	httpResp, err := client.Do(httpReq)
	if err != nil {
		logger.Errorf(ctx, "productService.getProduct: failed to do product request: %v", err)
		return nil, fmt.Errorf("failed to do product request: %w", err)
	}
	defer httpResp.Body.Close()

	metrics.UpdateExternalResponseCode(
		ServiceName,
		"GetProduct",
		http.StatusText(httpResp.StatusCode),
	)

	if httpResp.StatusCode == http.StatusOK {

		jsonResp, err := io.ReadAll(httpResp.Body)
		if err != nil {
			logger.Errorf(ctx, "productService.getProduct: failed to get product response body: %v", err)
			return nil, fmt.Errorf("failed to get product response body: %w", err)
		}

		resp = &GetProductResponse{}

		err = json.Unmarshal(jsonResp, resp)
		if err != nil {
			logger.Errorf(ctx, "productService.getProduct: failed to unmarshal product response body: %v", err)
			return nil, fmt.Errorf("failed to unmarshal product response body: %w", err)
		}

		return resp, nil
	} else if httpResp.StatusCode == http.StatusNotFound {
		logger.Warn(ctx, "productService.getProduct: product not found")
		return nil, model.ErrNotFound
	} else if httpResp.StatusCode == http.StatusTooManyRequests ||
		httpResp.StatusCode == StatusEnhanceYourCalm {
		logger.Warn(ctx, "productService.getProduct: too many requests")
		return nil, model.ErrTooManyRequests
	} else {
		logger.Error(ctx, "productService.getProduct: error")
		return nil, model.ErrUnknownError
	}
}

func (c *client) GetProductWithRetries(ctx context.Context, skuID int64) (*GetProductResponse, error) {
	for i := 0; i < config.ProductServiceRetries; i++ {
		logger.Infof(ctx, "productService.GetProduct: start %d try for product %d", i, skuID)

		resp, err := c.GetProduct(ctx, skuID)
		if err != nil {
			if errors.Is(err, model.ErrTooManyRequests) {
				if i == config.ProductServiceRetries-1 {
					return nil, err
				}
			} else {
				return nil, err
			}
		} else {
			return resp, nil
		}

	}

	return nil, model.ErrUnknownError
}
