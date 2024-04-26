package cartapi

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gookit/validate"
	"route256.ozon.ru/project/cart/internal/model"
	"route256.ozon.ru/project/cart/internal/pkg/logger"
	cartservice "route256.ozon.ru/project/cart/internal/service/cart_service"
)

func (a *api) GetItemsByUserID() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		req, err := toGetItemsByUserIDRequest(r)
		if err != nil {
			logger.Errorf(ctx, "handleCartGet: request is not valid: %v", err)
			http.Error(w, fmt.Sprintf("request is not valid: %s", err), http.StatusBadRequest)
			return
		}

		cart, err := a.cartService.GetItemsByUserID(ctx, req.UserID)
		if err != nil {
			logger.Errorf(ctx, "handleCartGet: failed to get cart: %v", err)
			if errors.Is(err, model.ErrNotFound) {
				http.Error(w, fmt.Sprintf("cart for user %d not found", req.UserID), http.StatusNotFound)
			} else {
				http.Error(w, "internal error", http.StatusInternalServerError)
			}
			return
		}

		err = toGetItemsByUserIDResponse(ctx, w, cart)
		if err != nil {
			logger.Errorf(ctx, "failed to write response: %v", err)
		}

	}
}

func toGetItemsByUserIDRequest(r *http.Request) (*GetItemsByUserIDRequest, error) {
	userID, _ := strconv.ParseInt(r.PathValue(paramUserID), 10, 64)

	req := &GetItemsByUserIDRequest{
		UserID: userID,
	}

	v := validate.Struct(req)
	if !v.Validate() {
		return nil, v.Errors
	}

	return req, nil
}

func toGetItemsByUserIDResponse(ctx context.Context, w http.ResponseWriter, cart *cartservice.Cart) error {
	json, err := json.Marshal(cart)
	if err != nil {
		logger.Errorf(ctx, "handleCartGet: failed to marshal cart response: %v", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return nil
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(json)
	if err != nil {
		logger.Errorf(ctx, "failed to write response: %v", err)
		return err
	}

	return nil
}
