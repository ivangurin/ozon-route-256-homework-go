package orderapi

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"route256.ozon.ru/project/loms/internal/model"
	"route256.ozon.ru/project/loms/pkg/api/order/v1"
)

func (a *API) Pay(ctx context.Context, req *order.OrderPayRequest) (*emptypb.Empty, error) {
	err := a.orderService.Pay(ctx, req.GetOrderId())
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, status.Errorf(codes.NotFound, err.Error())
		} else {
			return nil, status.Errorf(codes.Internal, err.Error())
		}
	}

	return &emptypb.Empty{}, nil
}
