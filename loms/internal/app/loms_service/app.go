package lomsservice

import (
	"context"
	"fmt"

	"route256.ozon.ru/project/loms/internal/config"
	grpcserver "route256.ozon.ru/project/loms/internal/pkg/grpc_server"
	httpserver "route256.ozon.ru/project/loms/internal/pkg/http_server"
	"route256.ozon.ru/project/loms/internal/pkg/logger"
	serviceprovider "route256.ozon.ru/project/loms/internal/pkg/service_provider"
)

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
	logger.Info(a.ctx, "app is starting...")
	defer logger.Info(a.ctx, "app finished")

	closer := a.sp.GetCloser()
	defer closer.Wait()

	// Grpc Server
	grpcServer := grpcserver.NewServer(a.ctx, config.LomsServiceGrpcPort)
	closer.Add(grpcServer.Stop)

	grpcServer.RegisterAPI([]grpcserver.API{
		a.sp.GetOrderAPI(a.ctx),
		a.sp.GetStockAPI(a.ctx),
	})

	go func() {
		logger.Info(a.ctx, "grpc server is starting...")
		err := grpcServer.Start()
		if err != nil {
			logger.Errorf(a.ctx, "failed to start grpc server: %v", err)
			closer.CloseAll()
			return
		}
		logger.Info(a.ctx, "grpc server finished")
	}()

	// Http Server
	httpServer, err := httpserver.NewServer(a.ctx, config.LomsServiceHttpPort, config.LomsServiceGrpcPort)
	if err != nil {
		logger.Errorf(a.ctx, "failed to create http server: %v", err)
		closer.CloseAll()
		return fmt.Errorf("failed to create http server: %w", err)
	}
	closer.Add(httpServer.Stop)

	err = httpServer.RegisterAPI([]httpserver.API{
		a.sp.GetOrderAPI(a.ctx),
		a.sp.GetStockAPI(a.ctx),
	})
	if err != nil {
		logger.Errorf(a.ctx, "failed to register api: %v", err)
		closer.CloseAll()
		return fmt.Errorf("failed to register api: %w", err)
	}

	go func() {
		logger.Info(a.ctx, "http server is starting...")
		err := httpServer.Start()
		if err != nil {
			logger.Errorf(a.ctx, "failed to start http server: %v", err)
			closer.CloseAll()
			return
		}
		logger.Info(a.ctx, "http server finished")
	}()

	return nil
}
