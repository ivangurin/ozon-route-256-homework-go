package cartapi

import (
	"fmt"
	"net/http"

	"route256.ozon.ru/project/cart/internal/model"
	cartservice "route256.ozon.ru/project/cart/internal/service/cart_service"
)

const (
	paramUserID string = "user_id"
	paramSkuID  string = "sku_id"
	paramCount  string = "count"
)

type IAPI interface {
	GetDescription() *model.HttpAPIDescription
}

type api struct {
	cartService cartservice.Service
}

func NewAPI(cartService cartservice.Service) IAPI {
	return &api{
		cartService: cartService,
	}
}

func (a *api) GetDescription() *model.HttpAPIDescription {
	return &model.HttpAPIDescription{
		Handlers: model.HttpApiHandlers{
			{
				Pattern: fmt.Sprintf("%s /user/{%s}/cart/{%s}", http.MethodPost, paramUserID, paramSkuID),
				Handler: a.AddItem(),
			},
			{
				Pattern: fmt.Sprintf("%s /user/{%s}/cart/{%s}", http.MethodDelete, paramUserID, paramSkuID),
				Handler: a.DeleteItem(),
			},
			{
				Pattern: fmt.Sprintf("%s /user/{%s}/cart", http.MethodGet, paramUserID),
				Handler: a.GetItemsByUserID(),
			},
			{
				Pattern: fmt.Sprintf("%s /user/{%s}/cart", http.MethodDelete, paramUserID),
				Handler: a.DeleteItemsByUserID(),
			},
		},
	}
}
