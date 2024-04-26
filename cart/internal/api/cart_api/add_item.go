package cartapi

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gookit/validate"
	"route256.ozon.ru/project/cart/internal/model"
	"route256.ozon.ru/project/cart/internal/pkg/logger"
)

func (a *api) AddItem() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		req, err := toAddItemRequest(ctx, r)
		if err != nil {
			logger.Errorf(ctx, "handleAddItem: request is not valid: %v", err)
			http.Error(w, fmt.Sprintf("request is not valid: %s", err), http.StatusBadRequest)
			return
		}

		err = a.cartService.AddItem(r.Context(), req.UserID, req.SkuID, req.Quantity)
		if err != nil {
			logger.Errorf(ctx, "handleAddItem failed to add item: %v", err)
			if errors.Is(err, model.ErrNotFound) {
				http.Error(w, fmt.Sprintf("sku %d not found", req.SkuID), http.StatusNotFound)
			} else if errors.Is(err, model.ErrInsufficientSock) {
				http.Error(w, fmt.Sprintf("insufficient stock for %d", req.SkuID), http.StatusBadRequest)
			} else {
				http.Error(w, "internal error", http.StatusInternalServerError)
			}
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func toAddItemRequest(ctx context.Context, r *http.Request) (*AddItemRequest, error) {

	userID, _ := strconv.ParseInt(r.PathValue(paramUserID), 10, 64)
	skuID, _ := strconv.ParseInt(r.PathValue(paramSkuID), 10, 64)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	data := &AddItemRequestBody{}
	err = json.Unmarshal(body, data)
	if err != nil {
		logger.Errorf(ctx, "handleAddItem: failed to unmarshal body json: %v", err)
		return nil, err
	}

	req := &AddItemRequest{
		UserID:   userID,
		SkuID:    skuID,
		Quantity: data.Count,
	}

	v := validate.Struct(req)
	if !v.Validate() {
		return nil, v.Errors
	}

	return req, nil
}
