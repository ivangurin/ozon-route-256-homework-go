package orderapi

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"route256.ozon.ru/project/loms/pkg/api/order/v1"
)

func (a *API) Create(ctx context.Context, req *order.OrderCreateRequest) (*order.OrderCreateResponse, error) {
	oredrID, err := a.orderService.Create(req.GetUser(), toItems(req.GetItems()))
	if err != nil {
		return nil, status.Errorf(codes.FailedPrecondition, err.Error())
	}

	return &order.OrderCreateResponse{
		OrderId: oredrID,
	}, nil
}
