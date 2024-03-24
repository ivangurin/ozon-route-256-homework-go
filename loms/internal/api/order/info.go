package orderapi

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"route256.ozon.ru/project/loms/internal/model"
	"route256.ozon.ru/project/loms/pkg/api/order/v1"
)

func (a *API) Info(ctx context.Context, req *order.OrderInfoRequest) (*order.OrderInfoResponse, error) {
	order, err := a.orderService.Info(ctx, req.GetOrderId())
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, status.Errorf(codes.NotFound, err.Error())
		} else {
			return nil, status.Errorf(codes.Internal, err.Error())
		}
	}

	return toRespOrder(order), nil
}
