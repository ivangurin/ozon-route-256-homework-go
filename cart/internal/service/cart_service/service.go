package cartservice

import (
	"context"

	productservice "route256.ozon.ru/project/cart/internal/client/product_service"
	cartstorage "route256.ozon.ru/project/cart/internal/repository/cart_storage"
)

type IService interface {
	AddItem(ctx context.Context, userID int64, skuID int64, quantity uint16) error
	DeleteItem(ctx context.Context, userID int64, skuID int64) error
	GetItemsByUserID(ctx context.Context, userID int64) (*Cart, error)
	DeleteItemsByUserID(ctx context.Context, userID int64) error
}

type cartService struct {
	productService productservice.IClient
	cartStorage    cartstorage.IStorage
}

func NewService(
	productService productservice.IClient,
	cartStorage cartstorage.IStorage,
) IService {
	return &cartService{
		productService: productService,
		cartStorage:    cartStorage,
	}
}
