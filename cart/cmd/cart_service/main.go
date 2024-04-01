package main

import (
	"context"

	cartservice "route256.ozon.ru/project/cart/internal/app/cart_service"
	"route256.ozon.ru/project/cart/internal/pkg/logger"
)

func main() {
	ctx := context.Background()
	app := cartservice.NewApp(ctx)
	if err := app.Run(); err != nil {
		logger.Fatalf("can't run app: %v", err)
	}
}
