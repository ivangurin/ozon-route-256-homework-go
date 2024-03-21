package cartservice

import (
	"context"

	lomsservice "route256.ozon.ru/project/cart/internal/pkg/client/loms_service"
	productservice "route256.ozon.ru/project/cart/internal/pkg/client/product_service"
	cartstorage "route256.ozon.ru/project/cart/internal/repository/cart_storage"
)

type Service interface {
	AddItem(ctx context.Context, userID int64, skuID int64, quantity uint16) error
	DeleteItem(ctx context.Context, userID int64, skuID int64) error
	GetItemsByUserID(ctx context.Context, userID int64) (*Cart, error)
	DeleteItemsByUserID(ctx context.Context, userID int64) error
	Checkout(ctx context.Context, userID int64) (int64, error)
}

type service struct {
	productService productservice.Client
	cartStorage    cartstorage.Storage
	lomsService    *lomsservice.Client
}

func NewService(
	productService productservice.Client,
	cartStorage cartstorage.Storage,
	lomsService *lomsservice.Client,
) Service {
	return &service{
		productService: productService,
		cartStorage:    cartStorage,
		lomsService:    lomsService,
	}
}
