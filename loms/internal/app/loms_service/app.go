package lomsservice

import (
	"context"
	"fmt"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"route256.ozon.ru/project/loms/internal/config"
	grpcserver "route256.ozon.ru/project/loms/internal/pkg/grpc_server"
	httpserver "route256.ozon.ru/project/loms/internal/pkg/http_server"
	"route256.ozon.ru/project/loms/internal/pkg/logger"
	serviceprovider "route256.ozon.ru/project/loms/internal/pkg/service_provider"
)

type api interface {
	RegisterGrpcServer(server *grpc.Server)
	RegisterHttpHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error
}

type App interface {
	Run() error
}

type app struct {
	ctx context.Context
	sp  *serviceprovider.ServiceProvider
}

func NewApp(ctx context.Context) App {
	return &app{
		ctx: ctx,
		sp:  serviceprovider.GetServiceProvider(),
	}
}

func (a *app) Run() error {
	logger.Info("app is starting...")
	defer logger.Info("app finished")

	closer := a.sp.GetCloser()
	defer closer.Wait()

	api := []api{
		a.sp.GetOrderAPI(a.ctx),
		a.sp.GetStockAPI(a.ctx),
	}

	// Grpc Server
	grpcServer := grpcserver.NewServer(a.ctx, config.LomsServiceGrpcPort)
	closer.Add(grpcServer.Stop)

	for _, singleAPI := range api {
		grpcServer.RegisterAPI(singleAPI)
	}

	go func() {
		logger.Info("grpc server is starting...")
		err := grpcServer.Start()
		if err != nil {
			logger.Error("failed to start grpc server", err)
			closer.CloseAll()
			return
		}
		logger.Info("grpc server finished")
	}()

	// Http Server
	httpServer, err := httpserver.NewServer(a.ctx, config.LomsServiceHttpPort, config.LomsServiceGrpcPort)
	if err != nil {
		logger.Error("failed to create http server", err)
		closer.CloseAll()
		return fmt.Errorf("failed to create http server: %w", err)
	}
	closer.Add(httpServer.Stop)

	for _, singleAPI := range api {
		err := httpServer.RegisterAPI(singleAPI)
		if err != nil {
			logger.Error("failed to register api", err)
			closer.CloseAll()
			return fmt.Errorf("failed to register api: %w", err)
		}
	}

	go func() {
		logger.Info("http server is starting...")
		err := httpServer.Start()
		if err != nil {
			logger.Error("failed to start http server", err)
			closer.CloseAll()
			return
		}
		logger.Info("http server finished")
	}()

	return nil
}
