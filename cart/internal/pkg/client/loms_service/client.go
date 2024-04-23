package lomsservice

import (
	"google.golang.org/grpc"
	order_api "route256.ozon.ru/project/cart/internal/pb/api/order/v1"
	stock_api "route256.ozon.ru/project/cart/internal/pb/api/stock/v1"
)

type Client struct {
	StockAPI stock_api.StockAPIClient
	OrderAPI order_api.OrderAPIClient
}

const (
	ServiceName = "loms-service"
)

func NewClient(conn *grpc.ClientConn) *Client {
	return &Client{
		StockAPI: stock_api.NewStockAPIClient(conn),
		OrderAPI: order_api.NewOrderAPIClient(conn),
	}
}
