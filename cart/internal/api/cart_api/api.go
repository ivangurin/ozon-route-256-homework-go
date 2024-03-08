package cartapi

import (
	"fmt"
	"net/http"

	cartservice "route256.ozon.ru/project/cart/internal/service/cart_service"
)

const (
	paramUserID string = "user_id"
	paramSkuID  string = "sku_id"
	paramCount  string = "count"
)

type IAPI interface {
}

type api struct {
	cartService cartservice.IService
}

func NewAPI(cartService cartservice.IService) IAPI {
	api := &api{
		cartService: cartService,
	}

	http.HandleFunc(fmt.Sprintf("POST /user/{%s}/cart/{%s}", paramUserID, paramSkuID), api.AddItem())
	http.HandleFunc(fmt.Sprintf("DELETE /user/{%s}/cart/{%s}", paramUserID, paramSkuID), api.DeleteItem())
	http.HandleFunc(fmt.Sprintf("GET /user/{%s}/cart", paramUserID), api.GetItemsByUserID())
	http.HandleFunc(fmt.Sprintf("DELETE /user/{%s}/cart", paramUserID), api.DeleteItemsByUserID())

	return api
}
