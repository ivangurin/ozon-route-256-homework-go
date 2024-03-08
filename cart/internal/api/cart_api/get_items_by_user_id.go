package cartapi

import (
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

		logger.Info(fmt.Sprintf("handleCartGet: start handle request: %s", r.RequestURI))
		defer logger.Info(fmt.Sprintf("handleCartGet: finish handle request: %s", r.RequestURI))

		req, err := toGetItemsByUserIDRequest(r)
		if err != nil {
			logger.Error("handleCartGet: request is not valid", err)
			http.Error(w, fmt.Sprintf("request is not valid: %s", err), http.StatusBadRequest)
			return
		}

		cart, err := a.cartService.GetItemsByUserID(r.Context(), req.UserID)
		if err != nil {
			logger.Error("handleCartGet: failed to get cart", err)
			if errors.Is(err, model.ErrNotFound) {
				http.Error(w, fmt.Sprintf("cart for user %d not found", req.UserID), http.StatusNotFound)
			} else {
				http.Error(w, "interanl error", http.StatusInternalServerError)
			}
			return
		}

		err = toGetItemsByUserIDResponse(w, cart)
		if err != nil {
			logger.Error("failed to write response", err)
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

func toGetItemsByUserIDResponse(w http.ResponseWriter, cart *cartservice.Cart) error {
	json, err := json.Marshal(cart)
	if err != nil {
		logger.Error("handleCartGet: failed to marshal cart response", err)
		http.Error(w, "interanl error", http.StatusInternalServerError)
		return nil
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(json)
	if err != nil {
		logger.Error("failed to write response", err)
		return err
	}

	return nil
}
