package cartapi

import (
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
		logger.Info(fmt.Sprintf("handleCheckout: start handle request: %s", r.RequestURI))
		defer logger.Info(fmt.Sprintf("handleCheckout: finish handle request: %s", r.RequestURI))

		req, err := toCheckoutRequest(r)
		if err != nil {
			logger.Error("handleCheckout: request is not valid", err)
			http.Error(w, fmt.Sprintf("request is not valid: %s", err), http.StatusBadRequest)
			return
		}

		orderID, err := a.cartService.Checkout(r.Context(), req.UserID)
		if err != nil {
			logger.Error("handleCheckout: failed to checkout", err)
			if errors.Is(err, model.ErrNotFound) {
				http.Error(w, "cart not found", http.StatusNotFound)
			} else {
				http.Error(w, "internal error", http.StatusInternalServerError)
			}
			return
		}

		err = toCheckoutResponse(w, orderID)
		if err != nil {
			logger.Error("failed to write response", err)
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
		logger.Error("handleCheckout: failed to unmarshal request body", err)
		return nil, err
	}

	v := validate.Struct(req)
	if !v.Validate() {
		return nil, v.Errors
	}

	return req, nil
}

func toCheckoutResponse(w http.ResponseWriter, orderID int64) error {

	resp := &CheckoutResponse{
		OrderID: orderID,
	}

	json, err := json.Marshal(resp)
	if err != nil {
		logger.Error("handleCheckout: failed to marshal response", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return err
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(json)
	if err != nil {
		logger.Error("failed to write response", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return err
	}

	return nil
}
