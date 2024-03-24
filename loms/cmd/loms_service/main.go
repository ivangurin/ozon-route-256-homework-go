package main

import (
	"context"

	lomsservice "route256.ozon.ru/project/loms/internal/app/loms_service"
	"route256.ozon.ru/project/loms/internal/pkg/logger"
)

func main() {
	app := lomsservice.NewApp(context.Background())
	if err := app.Run(); err != nil {
		logger.Fatal("can't run app", err)
	}
}
