package cartapi

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gookit/validate"
	"route256.ozon.ru/project/cart/internal/pkg/logger"
)

func (a *api) DeleteItemsByUserID() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		logger.Info(fmt.Sprintf("handleDeleteItemsByUserID: start handle request: %s", r.RequestURI))
		defer logger.Info(fmt.Sprintf("handleDeleteItemsByUserID: finish handle request: %s", r.RequestURI))

		req, err := toDeleteItemsByUserIDRequest(r)
		if err != nil {
			logger.Error("handleDeleteItemsByUserID: request is not valid", err)
			http.Error(w, fmt.Sprintf("request is not valid: %s", err), http.StatusBadRequest)
			return
		}

		err = a.cartService.DeleteItemsByUserID(r.Context(), req.UserID)
		if err != nil {
			logger.Error("handleDeleteItemsByUserID: failed to delete items", err)
			http.Error(w, "internal error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func toDeleteItemsByUserIDRequest(r *http.Request) (*DeleteItemsByUserIDRequest, error) {
	userID, _ := strconv.ParseInt(r.PathValue(paramUserID), 10, 64)

	req := &DeleteItemsByUserIDRequest{
		UserID: userID,
	}

	v := validate.Struct(req)
	if !v.Validate() {
		return nil, v.Errors
	}

	return req, nil
}
