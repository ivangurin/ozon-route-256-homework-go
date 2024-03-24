//go:generate mkdir -p ./mocks
//go:generate minimock -i "StockClientMock" -o ./mocks/ -s ".go" -g
//go:generate minimock -i "OrderClientMock" -o ./mocks/ -s ".go" -g
package lomsservice

import (
	"route256.ozon.ru/project/cart/internal/pb/api/order/v1"
	"route256.ozon.ru/project/cart/internal/pb/api/stock/v1"
)

// StockClientMock ...
type StockClientMock interface {
	stock.StockClient
}

// OrderClientMock ...
type OrderClientMock interface {
	order.OrderClient
}
