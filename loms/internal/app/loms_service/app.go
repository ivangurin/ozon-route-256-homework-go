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
	ctx, cancel := context.WithCancel(ctx)

	sp := serviceprovider.GetServiceProvider()
	sp.GetCloser().Add(func() error {
		cancel()
		return nil
	})

	return &app{
		ctx: ctx,
		sp:  sp,
	}
}

func (a *app) Run() error {
	logger.Info("app is starting...")
	defer logger.Info("app finished")

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
		logger.Info("grpc server is starting...")
		err := grpcServer.Start()
		if err != nil {
			logger.Errorf("failed to start grpc server: %v", err)
			closer.CloseAll()
			return
		}
		logger.Info("grpc server finished")
	}()

	// Http Server
	httpServer, err := httpserver.NewServer(a.ctx, config.LomsServiceHttpPort, config.LomsServiceGrpcPort)
	if err != nil {
		logger.Errorf("failed to create http server: %v", err)
		closer.CloseAll()
		return fmt.Errorf("failed to create http server: %w", err)
	}
	closer.Add(httpServer.Stop)

	err = httpServer.RegisterAPI([]httpserver.API{
		a.sp.GetOrderAPI(a.ctx),
		a.sp.GetStockAPI(a.ctx),
	})
	if err != nil {
		logger.Errorf("failed to register api: %v", err)
		closer.CloseAll()
		return fmt.Errorf("failed to register api: %w", err)
	}

	go func() {
		logger.Info("http server is starting...")
		err := httpServer.Start()
		if err != nil {
			logger.Errorf("failed to start http server: %v", err)
			closer.CloseAll()
			return
		}
		logger.Info("http server finished")
	}()

	return nil
}
