package orderapi

import (
	"context"
	"fmt"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	orderservice "route256.ozon.ru/project/loms/internal/service/order_service"
	"route256.ozon.ru/project/loms/pkg/api/order/v1"
)

type API struct {
	order.UnimplementedOrderServer
	orderService orderservice.Service
}

func NewAPI(
	orderService orderservice.Service,
) *API {
	return &API{
		orderService: orderService,
	}
}

func (a *API) RegisterGrpcServer(server *grpc.Server) {
	order.RegisterOrderServer(server, a)
}

func (a *API) RegisterHttpHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	err := order.RegisterOrderHandler(ctx, mux, conn)
	if err != nil {
		return fmt.Errorf("failed to register gateway: %w", err)
	}
	return nil
}
