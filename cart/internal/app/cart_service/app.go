package cartservice

import (
	"context"

	"route256.ozon.ru/project/cart/internal/config"
	httpserver "route256.ozon.ru/project/cart/internal/pkg/http_server"
	"route256.ozon.ru/project/cart/internal/pkg/logger"
	serviceprovider "route256.ozon.ru/project/cart/internal/pkg/service_provider"
	"route256.ozon.ru/project/cart/internal/pkg/tracer"
)

type IApp interface {
	Run() error
}

type app struct {
	ctx context.Context
	sp  *serviceprovider.ServiceProvider
}

func NewApp(ctx context.Context) IApp {
	ctx, cancel := context.WithCancel(ctx)

	sp := serviceprovider.GetServiceProvider(ctx)
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
	logger.Info(a.ctx, "app cartService is starting...")
	defer logger.Info(a.ctx, "app cartService finished")

	closer := a.sp.GetCloser()
	defer closer.Wait()

	cartAPI := a.sp.GetCartAPI()

	// http server
	httpServer := httpserver.NewServer(a.ctx, config.CartServiceHttpPort)
	httpServer.AddHandlers(cartAPI.GetDescription().Handlers)
	closer.Add(httpServer.Stop)

	go func() {
		logger.Info(a.ctx, "http cartService server is starting...")
		err := httpServer.Start()
		if err != nil {
			logger.Errorf(a.ctx, "failed to start http serve: %v", err)
			closer.CloseAll()
		}
	}()

	// logger
	closer.Add(logger.Close)

	// tracer
	closer.Add(tracer.Close)

	return nil
}
