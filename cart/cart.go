package main

import (
	"context"

	"route256.ozon.ru/project/cart/internal/app"
	"route256.ozon.ru/project/cart/internal/pkg/logger"
)

func main() {
	ctx := context.Background()
	app := app.NewApp(ctx)
	if err := app.Run(); err != nil {
		logger.Error("can't run app", err)
		return
	}
}
