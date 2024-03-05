package app

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"route256.ozon.ru/project/cart/internal/config"
	"route256.ozon.ru/project/cart/internal/pkg/logger"
)

const (
	paramUserID string = "user_id"
	paramSkuID  string = "sku_id"
	paramCount  string = "count"
)

type IApp interface {
	Run() error
}

type app struct {
	ctx context.Context
	sp  *tServiceProvider
}

func NewApp(ctx context.Context) IApp {
	return &app{
		ctx: ctx,
		sp:  GetServiceProvider(),
	}
}

func (a *app) Run() error {
	logger.Info("starting app...")

	logger.Info("listner is createing...")
	conn, err := net.Listen(config.AppProtocol, fmt.Sprintf(":%s", config.AppAddressPort))
	if err != nil {
		return fmt.Errorf("faild to create listner: %w", err)
	}
	defer conn.Close()
	logger.Info("listner is created")

	http.HandleFunc(fmt.Sprintf("POST /user/{%s}/cart/{%s}", paramUserID, paramSkuID), a.handleAddItem(a.ctx))
	http.HandleFunc(fmt.Sprintf("DELETE /user/{%s}/cart/{%s}", paramUserID, paramSkuID), a.handleDeleteItem(a.ctx))
	http.HandleFunc(fmt.Sprintf("GET /user/{%s}/cart", paramUserID), a.handleGetItemsByUserID(a.ctx))
	http.HandleFunc(fmt.Sprintf("DELETE /user/{%s}/cart", paramUserID), a.handleDeleteItemsByUserID(a.ctx))

	logger.Info("srtarting http server...")
	if err := http.Serve(conn, nil); err != nil {
		return fmt.Errorf("faild to start http server: %w", err)
	}

	return nil
}
