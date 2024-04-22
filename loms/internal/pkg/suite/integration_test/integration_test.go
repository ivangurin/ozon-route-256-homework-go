//go:build integration

package integrationtest

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	"route256.ozon.ru/project/loms/internal/config"
	"route256.ozon.ru/project/loms/internal/db"
	"route256.ozon.ru/project/loms/internal/model"
	"route256.ozon.ru/project/loms/internal/pkg/logger"
	orderstorage "route256.ozon.ru/project/loms/internal/repository/order_storage"
	stockstorage "route256.ozon.ru/project/loms/internal/repository/stock_storage"
)

const (
	user = int64(1)
	sku  = int64(773297411)
)

type Suite struct {
	suite.Suite
	ctx          context.Context
	dbClient     db.Client
	stockStorage stockstorage.Repository
	orderStorage orderstorage.Repository
}

func TestIntegrationTest(t *testing.T) {
	suite.Run(t, &Suite{})

}

func (s *Suite) SetupSuite() {
	s.ctx = context.Background()

	s.dbClient = db.NewClient(s.ctx)

	var err error
	err = s.dbClient.AddShard(config.PostgresTestUrl, config.PostgresTestUrl)
	if err != nil {
		logger.Errorf(s.ctx, "failed to create db client: %v", err)
		return
	}

	s.stockStorage = stockstorage.NewRepository(s.ctx, s.dbClient)
	if err != nil {
		logger.Fatalf(s.ctx, "failed to create stock storage: %v", err)
	}

	s.orderStorage = orderstorage.NewRepository(s.ctx, s.dbClient)
	if err != nil {
		logger.Fatalf(s.ctx, "failed to create order storage: %v", err)
	}

}

func (s *Suite) TearDownSuite() {
	s.dbClient.Close()
}

func (s *Suite) TestStockStorage() {

	qty, err := s.stockStorage.GetBySku(s.ctx, 123)
	s.Require().Error(err, "Нет ошибки при запросе несуществующего запаса")

	// Получаем количество на начало
	qty, err = s.stockStorage.GetBySku(s.ctx, sku)
	s.Require().NoError(err, "Ошибка при получении количества")
	s.Require().Equal(uint16(140), qty, "Не совпало количество на начало")

	// Резервируем 10 штук
	err = s.stockStorage.Reserve(s.ctx, stockstorage.ReserveItems{
		{
			Sku:      sku,
			Quantity: 10,
		},
	})
	s.Require().NoError(err, "Ошибка при резервировании")

	// Получаем количество после резервирования
	qty, err = s.stockStorage.GetBySku(s.ctx, sku)
	s.Require().NoError(err, "Ошибка при получении количества")
	s.Require().Equal(uint16(130), qty, "Не совпало количество после резервирования")

	// Резервируем 20 штук
	err = s.stockStorage.Reserve(s.ctx, stockstorage.ReserveItems{
		{
			Sku:      sku,
			Quantity: 20,
		},
	})
	s.Require().NoError(err, "Ошибка при резервировании")

	// Получаем количество после резервирования
	qty, err = s.stockStorage.GetBySku(s.ctx, sku)
	s.Require().NoError(err, "Ошибка при получении количества")
	s.Require().Equal(uint16(110), qty, "Не совпало количество после резервирования")

	// Отменяем резерв
	err = s.stockStorage.CancelReserve(s.ctx, stockstorage.ReserveItems{
		{
			Sku:      sku,
			Quantity: 20,
		},
	})

	// Получаем количество после отмены резервирования
	qty, err = s.stockStorage.GetBySku(s.ctx, sku)
	s.Require().NoError(err, "Ошибка при получении количества")
	s.Require().Equal(uint16(130), qty, "Не совпало количество после отмены резервирования")

	// Резервируем 30 штук
	err = s.stockStorage.Reserve(s.ctx, stockstorage.ReserveItems{
		{
			Sku:      sku,
			Quantity: 30,
		},
	})
	s.Require().NoError(err, "Ошибка при резервировании")

	// Получаем количество после резервирования
	qty, err = s.stockStorage.GetBySku(s.ctx, sku)
	s.Require().NoError(err, "Ошибка при получении количества")
	s.Require().Equal(uint16(100), qty, "Не совпало количество после резервирования")

	// Выдача резерва
	err = s.stockStorage.RemoveReserve(s.ctx, stockstorage.ReserveItems{
		{
			Sku:      sku,
			Quantity: 30,
		},
	})

	// Получаем количество после выдачи резервирования
	qty, err = s.stockStorage.GetBySku(s.ctx, sku)
	s.Require().NoError(err, "Ошибка при получении количества")
	s.Require().Equal(uint16(100), qty, "Не совпало количество после выдачи резервирования")

}

func (s *Suite) TestOrderStorage() {

	// Создаем заказ
	orderID, err := s.orderStorage.Create(s.ctx, user, []*orderstorage.OrderItem{{
		Sku:      sku,
		Quantity: 10,
	}})
	s.Require().NoError(err, "Ошибка при создании заказа")
	s.Require().Equal(int64(1), orderID, "Не совпал ID созданного заказа")

	// Получаем созданный заказ
	order, err := s.orderStorage.GetByID(s.ctx, orderID)
	s.Require().NoError(err, "Ошибка при получении созданного заказа")
	s.Require().Equal(user, order.User, "Не совпал ID пользователя")
	s.Require().Equal(model.OrderStatusNew, order.Status, "Не совпал статус заказа")

	items := order.GetItemsMap()
	s.Require().Len(items, 1, "Не совпало количество позиций в заказе")

	item, exists := items[sku]
	s.Require().True(exists, "Не найдена позиция в заказе")
	s.Require().NotNil(item, "Не найдена позиция в заказе")
	s.Require().Equal(sku, item.Sku, "Не совпал код товара в заказе")
	s.Require().Equal(uint16(10), item.Quantity, "Не совпало количество в заказе")

	// Изменяем статусу заказа
	err = s.orderStorage.SetStatus(s.ctx, orderID, model.OrderStatusAwaitingPayment)
	s.Require().NoError(err, "Ошибка при изменении статуса заказа")

	// Получаем измененный заказ
	order, err = s.orderStorage.GetByID(s.ctx, orderID)
	s.Require().NoError(err, "Ошибка при получении измененного заказа")
	s.Require().Equal(model.OrderStatusAwaitingPayment, order.Status, "Не совпал статус заказа")

}
