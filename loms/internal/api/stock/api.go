package stockapi

import (
	"context"
	"fmt"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	stockservice "route256.ozon.ru/project/loms/internal/service/stock_service"
	"route256.ozon.ru/project/loms/pkg/api/stock/v1"
)

type API struct {
	stock.UnimplementedStockAPIServer
	stockService stockservice.Service
}

func NewAPI(
	stockService stockservice.Service,
) *API {
	return &API{
		stockService: stockService,
	}
}

func (a *API) RegisterGrpcServer(server *grpc.Server) {
	stock.RegisterStockAPIServer(server, a)
}

func (a *API) RegisterHttpHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	err := stock.RegisterStockAPIHandler(ctx, mux, conn)
	if err != nil {
		return fmt.Errorf("failed to register gateway: %w", err)
	}
	return nil
}
