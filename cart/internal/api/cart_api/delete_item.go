package cartapi

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gookit/validate"
	"route256.ozon.ru/project/cart/internal/pkg/logger"
)

func (a *api) DeleteItem() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		logger.Info(fmt.Sprintf("handleDeleteItem: start handle request: %s", r.RequestURI))
		defer logger.Info(fmt.Sprintf("handleDeleteItem: finish handle request: %s", r.RequestURI))

		req, err := toDeleteItemReq(r)
		if err != nil {
			logger.Error("handleDeleteItem: request is not valid", err)
			http.Error(w, fmt.Sprintf("request is not valid: %s", err), http.StatusBadRequest)
			return
		}

		err = a.cartService.DeleteItem(r.Context(), req.UserID, req.SkuID)
		if err != nil {
			logger.Error("handleDeleteItem: failed to delete item", err)
			http.Error(w, "interanl error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func toDeleteItemReq(r *http.Request) (*DeleteItemRequest, error) {
	userID, _ := strconv.ParseInt(r.PathValue(paramUserID), 10, 64)
	skuID, _ := strconv.ParseInt(r.PathValue(paramSkuID), 10, 64)

	req := &DeleteItemRequest{
		UserID: userID,
		SkuID:  skuID,
	}

	v := validate.Struct(req)
	if !v.Validate() {
		return nil, v.Errors
	}

	return req, nil
}
