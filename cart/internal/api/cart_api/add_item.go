package cartapi

import (
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
		logger.Infof("handleAddItem: start handle request: %s", r.RequestURI)
		defer logger.Infof("handleAddItem: finish handle request: %s", r.RequestURI)

		req, err := toAddItemRequest(r)
		if err != nil {
			logger.Errorf("handleAddItem: request is not valid: %v", err)
			http.Error(w, fmt.Sprintf("request is not valid: %s", err), http.StatusBadRequest)
			return
		}

		err = a.cartService.AddItem(r.Context(), req.UserID, req.SkuID, req.Quantity)
		if err != nil {
			logger.Errorf("handleAddItem failed to add item: %v", err)
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

func toAddItemRequest(r *http.Request) (*AddItemRequest, error) {

	userID, _ := strconv.ParseInt(r.PathValue(paramUserID), 10, 64)
	skuID, _ := strconv.ParseInt(r.PathValue(paramSkuID), 10, 64)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	data := &AddItemRequestBody{}
	err = json.Unmarshal(body, data)
	if err != nil {
		logger.Errorf("handleAddItem: failed to unmarshal body json: %v", err)
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
