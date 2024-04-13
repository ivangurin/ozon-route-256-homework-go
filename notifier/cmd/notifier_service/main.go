package main

import (
	"context"

	notifierservice "route256.ozon.ru/project/notifier/internal/app/notifier_service"
)

func main() {
	ctx := context.Background()
	app := notifierservice.NewApp(ctx)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
