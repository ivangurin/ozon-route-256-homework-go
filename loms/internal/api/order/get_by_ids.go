package orderapi

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"route256.ozon.ru/project/loms/internal/model"
	"route256.ozon.ru/project/loms/pkg/api/order/v1"
)

func (a *API) GetByIDs(ctx context.Context, req *order.GetOrdersByIDsRequest) (*order.GetOrdersByIDsResponse, error) {
	orders, err := a.orderService.GetByIDs(ctx, req.GetId())
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, status.Errorf(codes.NotFound, err.Error())
		} else {
			return nil, status.Errorf(codes.Internal, err.Error())
		}
	}

	return &order.GetOrdersByIDsResponse{Orders: toRespOrders(orders)}, nil
}
