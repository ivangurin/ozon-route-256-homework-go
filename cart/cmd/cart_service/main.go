package main

import (
	"context"

	cartservice "route256.ozon.ru/project/cart/internal/app/cart_service"
	"route256.ozon.ru/project/cart/internal/pkg/logger"
)

func main() {
	app := cartservice.NewApp(context.Background())
	if err := app.Run(); err != nil {
		logger.Fatal("can't run app", err)
	}
}
