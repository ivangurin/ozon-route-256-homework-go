package main

import (
	"context"

	lomsservice "route256.ozon.ru/project/loms/internal/app/loms_service"
	"route256.ozon.ru/project/loms/internal/pkg/logger"
)

func main() {
	ctx := context.Background()
	app := lomsservice.NewApp(ctx)
	if err := app.Run(); err != nil {
		logger.Fatalf(ctx, "can't run app: %w", err)
	}
}
