//go:build integration

package integrationtest

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"route256.ozon.ru/project/cart/internal/model"
	cartservice "route256.ozon.ru/project/cart/internal/pkg/client/cart_service"
)

type Suite struct {
	suite.Suite
	ctx               context.Context
	cartServiceClient cartservice.IClient
}

func TestSuite(t *testing.T) {
	suite.Run(t, &Suite{})
}

func (s *Suite) SetupSuite() {
	s.ctx = context.Background()
	s.cartServiceClient = cartservice.NewClient("127.0.0.1:8082")
}

func (s *Suite) TestCartList() {

	// Чистим корзину перед тестом
	err := s.cartServiceClient.DeleteItemsByUserID(s.ctx, 1)
	require.NoError(s.T(), err, "Ошибка при удалении позиций корзины")

	// Добавялем первую позицию
	err = s.cartServiceClient.AddItem(s.ctx, 2, 1076963, 1)
	require.NoError(s.T(), err, "Ошибка при добавлении позиции 1076963")

	// Добавляем вторую позицию
	err = s.cartServiceClient.AddItem(s.ctx, 2, 1148162, 1)
	require.NoError(s.T(), err, "Ошибка при добавлении позиции 1148162")

	// Добавляем третью позицию
	err = s.cartServiceClient.AddItem(s.ctx, 2, 1625903, 1)
	require.NoError(s.T(), err, "Ошибка при добавлении позиции 1625903")

	// Получаем корзину и проверяем содержимое корзины
	cart, err := s.cartServiceClient.GetItemsByUserID(s.ctx, 2)
	require.NoError(s.T(), err, "Ошибка при получении корзины")
	require.NotNil(s.T(), cart, "Данные по корзине должны были вернуться")
	require.Len(s.T(), cart.Items, 3, "Должно было вернуться две позиции")

}

func (s *Suite) TestCartItemDelete() {

	// Чистим корзину перед тестом
	err := s.cartServiceClient.DeleteItemsByUserID(s.ctx, 2)
	require.NoError(s.T(), err, "Ошибка при удалении позиций корзины")

	// Добавялем первую позицию
	err = s.cartServiceClient.AddItem(s.ctx, 2, 1076963, 1)
	require.NoError(s.T(), err, "Ошибка при добавлении позиции 1076963")

	// Добавляем вторую позицию
	err = s.cartServiceClient.AddItem(s.ctx, 2, 1148162, 1)
	require.NoError(s.T(), err, "Ошибка при добавлении позиции 1148162")

	// Получаем корзину и проверяем содержимое корзины
	cart, err := s.cartServiceClient.GetItemsByUserID(s.ctx, 2)
	require.NoError(s.T(), err, "Ошибка при получении корзины")
	require.NotNil(s.T(), cart, "Данные по корзине должны были вернуться")
	require.Len(s.T(), cart.Items, 2, "Должно было вернуться две позиции")

	// Удаляем первую позицию
	err = s.cartServiceClient.DeleteItem(s.ctx, 2, 1076963)
	require.NoError(s.T(), err, "Ошибка при удалении позиции 1076963")

	// Получаем корзину и проверяем содержимое корзины
	cart, err = s.cartServiceClient.GetItemsByUserID(s.ctx, 2)
	require.NoError(s.T(), err, "Ошибка при получении корзины")
	require.NotNil(s.T(), cart, "Данные по корзине должны были вернуться")
	require.Len(s.T(), cart.Items, 1, "Должно была вернуться одна позиции")

	// Удаляем вторую позицию
	err = s.cartServiceClient.DeleteItem(s.ctx, 2, 1148162)
	require.NoError(s.T(), err, "Ошибка при удалении позиции 1148162")

	// Получаем корзину и проверяем содержимое корзины
	cart, err = s.cartServiceClient.GetItemsByUserID(s.ctx, 2)
	require.Error(s.T(), err, "Должна вернеуться ошибка что корзина не найдена")
	require.ErrorIs(s.T(), err, model.ErrNotFound, "Должна вернеуться ошибка что корзина не найдена")
	require.Nil(s.T(), cart, "Данные по корзине не должны были вернуться")

}
