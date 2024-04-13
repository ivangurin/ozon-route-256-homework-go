package main

import (
	"context"

	cartservice "route256.ozon.ru/project/cart/internal/app/cart_service"
)

func main() {
	ctx := context.Background()
	app := cartservice.NewApp(ctx)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
