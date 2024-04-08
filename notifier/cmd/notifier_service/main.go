package main

import (
	"context"

	notifierservice "route256.ozon.ru/project/notifier/internal/app/notifier_service"
	"route256.ozon.ru/project/notifier/internal/pkg/logger"
)

func main() {
	ctx := context.Background()
	app := notifierservice.NewApp(ctx)
	if err := app.Run(); err != nil {
		logger.Fatalf("can't run app: %w", err)
	}
}
