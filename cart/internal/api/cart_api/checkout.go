package cartapi

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/gookit/validate"
	"route256.ozon.ru/project/cart/internal/model"
	"route256.ozon.ru/project/cart/internal/pkg/logger"
)

func (a *api) Checkout() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Infof(r.Context(), "handleCheckout: start handle request: %s", r.RequestURI)
		defer logger.Infof(r.Context(), "handleCheckout: finish handle request: %s", r.RequestURI)

		req, err := toCheckoutRequest(r)
		if err != nil {
			logger.Errorf(r.Context(), "handleCheckout: request is not valid: %v", err)
			http.Error(w, fmt.Sprintf("request is not valid: %s", err), http.StatusBadRequest)
			return
		}

		orderID, err := a.cartService.Checkout(r.Context(), req.UserID)
		if err != nil {
			logger.Errorf(r.Context(), "handleCheckout: failed to checkout: %v", err)
			if errors.Is(err, model.ErrNotFound) {
				http.Error(w, "cart not found", http.StatusNotFound)
			} else {
				http.Error(w, "internal error", http.StatusInternalServerError)
			}
			return
		}

		err = toCheckoutResponse(r.Context(), w, orderID)
		if err != nil {
			logger.Errorf(r.Context(), "failed to write response: %v", err)
		}

	}
}

func toCheckoutRequest(r *http.Request) (*CheckoutRequest, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	req := &CheckoutRequest{}
	err = json.Unmarshal(body, req)
	if err != nil {
		logger.Errorf(r.Context(), "handleCheckout: failed to unmarshal request body: %v", err)
		return nil, err
	}

	v := validate.Struct(req)
	if !v.Validate() {
		return nil, v.Errors
	}

	return req, nil
}

func toCheckoutResponse(ctx context.Context, w http.ResponseWriter, orderID int64) error {

	resp := &CheckoutResponse{
		OrderID: orderID,
	}

	json, err := json.Marshal(resp)
	if err != nil {
		logger.Errorf(ctx, "handleCheckout: failed to marshal response: %v", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return err
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(json)
	if err != nil {
		logger.Errorf(ctx, "failed to write response: %v", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return err
	}

	return nil
}
