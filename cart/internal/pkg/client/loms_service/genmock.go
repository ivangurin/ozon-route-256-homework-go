//go:generate mockery --name=(.+)Mock --case=underscore  --with-expecter
package lomsservice

import (
	"route256.ozon.ru/project/cart/internal/pb/api/order/v1"
	"route256.ozon.ru/project/cart/internal/pb/api/stock/v1"
)

// StockClientMock ...
type StockClientMock interface {
	stock.StockAPIClient
}

// OrderClientMock ...
type OrderClientMock interface {
	order.OrderAPIClient
}
