package cartapi

import (
	"fmt"

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
	cartService cartservice.IService
}

func NewAPI(cartService cartservice.IService) IAPI {
	return &api{
		cartService: cartService,
	}
}

func (a *api) GetDescription() *model.HttpAPIDescription {
	return &model.HttpAPIDescription{
		Handlers: model.HttpApiHandlers{
			{
				Pattern: fmt.Sprintf("POST /user/{%s}/cart/{%s}", paramUserID, paramSkuID),
				Handler: a.AddItem(),
			},
			{
				Pattern: fmt.Sprintf("DELETE /user/{%s}/cart/{%s}", paramUserID, paramSkuID),
				Handler: a.DeleteItem(),
			},
			{
				Pattern: fmt.Sprintf("GET /user/{%s}/cart", paramUserID),
				Handler: a.GetItemsByUserID(),
			},
			{
				Pattern: fmt.Sprintf("DELETE /user/{%s}/cart", paramUserID),
				Handler: a.DeleteItemsByUserID(),
			},
		},
	}
}
