package main

import (
	"context"

	lomsservice "route256.ozon.ru/project/loms/internal/app/loms_service"
)

func main() {
	ctx := context.Background()
	app := lomsservice.NewApp(ctx)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
