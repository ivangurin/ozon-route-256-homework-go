package cartservice

import (
	"context"

	"route256.ozon.ru/project/cart/internal/config"
	httpserver "route256.ozon.ru/project/cart/internal/pkg/http_server"
	"route256.ozon.ru/project/cart/internal/pkg/logger"
	serviceprovider "route256.ozon.ru/project/cart/internal/pkg/service_provider"
)

type IApp interface {
	Run() error
}

type app struct {
	ctx context.Context
	sp  *serviceprovider.ServiceProvider
}

func NewApp(ctx context.Context) IApp {
	return &app{
		ctx: ctx,
		sp:  serviceprovider.GetServiceProvider(),
	}
}

func (a *app) Run() error {
	logger.Info("app cartService is starting...")
	defer logger.Info("app cartService finished")

	closer := a.sp.GetCloser()
	defer closer.Wait()

	cartAPI := a.sp.GetCartAPI()

	httpServer := httpserver.NewServer(config.CartServcePort)
	httpServer.AddHandlers(cartAPI.GetDescription().Handlers)
	closer.Add(httpServer.Stop)

	go func() {
		logger.Info("http cartService server is starting...")
		err := httpServer.Start()
		if err != nil {
			logger.Error("failed to start http server", err)
			closer.CloseAll()
		}
		logger.Info("http cartService server finished")
	}()

	return nil
}
