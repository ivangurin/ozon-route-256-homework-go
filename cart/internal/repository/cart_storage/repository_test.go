package cartstorage

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"route256.ozon.ru/project/cart/internal/model"
)

func TestCartStorage(t *testing.T) {
	ctx := context.Background()
	cartStorage := NewCartStorage()

	cart, err := cartStorage.GetItemsByUserID(ctx, 1)
	require.ErrorIs(t, err, model.ErrNotFound, "Корзина еще не существует")
	require.Nil(t, cart, "Данные по корзине не должны вернуться")

	err = cartStorage.AddItem(ctx, 1, 1, 1)
	require.NoError(t, err, "При добавлении ошибки быть не должно")

	cart, err = cartStorage.GetItemsByUserID(ctx, 1)
	require.NoError(t, err, "Корзина должна быть, ошибки быть не должно")
	require.NotNil(t, cart, "Корзина должна быть")
	require.Len(t, cart.Items, 1, "Должна быть одна позиция")

	_, exists := cart.Items[1]
	require.True(t, exists, "Должна быть позиция 1")

	err = cartStorage.AddItem(ctx, 1, 2, 2)
	require.NoError(t, err, "При добавлении ошибки быть не должно")

	cart, err = cartStorage.GetItemsByUserID(ctx, 1)
	require.NoError(t, err, "Корзина должна быть, ошибки быть не должно")
	require.NotNil(t, cart, "Корзина должна быть")
	require.Len(t, cart.Items, 2, "Должно быть две позиция")

	_, exists = cart.Items[2]
	require.True(t, exists, "Должна быть позиция 2")

	err = cartStorage.DeleteItem(ctx, 1, 1)
	require.NoError(t, err, "При удалении ошибки быть не должно")

	cart, err = cartStorage.GetItemsByUserID(ctx, 1)
	require.NoError(t, err, "Корзина должна быть, ошибки быть не должно")
	require.NotNil(t, cart, "Корзина должна быть")
	require.Len(t, cart.Items, 1, "Должна быть одна позиция")

	_, exists = cart.Items[2]
	require.True(t, exists, "Должна быть позиция 2")

	err = cartStorage.DeleteItemsByUserID(ctx, 1)
	require.NoError(t, err, "При удалении все корзины ошибки быть не должно")

	cart, err = cartStorage.GetItemsByUserID(ctx, 1)
	require.ErrorIs(t, err, model.ErrNotFound, "Корзина еще не существует")
	require.Nil(t, cart, "Данные по корзине не должны вернуться")

	err = cartStorage.AddItem(ctx, 1, 3, 1)
	require.NoError(t, err, "При добавлении ошибки быть не должно")

	err = cartStorage.DeleteItem(ctx, 1, 4)
	require.NoError(t, err, "При удалении несуществующей позиции ошибки быть не должно")

	err = cartStorage.DeleteItem(ctx, 1, 3)
	require.NoError(t, err, "При удалении ошибки быть не должно")

	cart, err = cartStorage.GetItemsByUserID(ctx, 1)
	require.ErrorIs(t, err, model.ErrNotFound, "Корзина еще не существует")
	require.Nil(t, cart, "Данные по корзине не должны вернуться")

	err = cartStorage.DeleteItem(ctx, 2, 1)
	require.NoError(t, err, "При удалении из несуществующей коризны ошибки быть не должно")

	err = cartStorage.AddItem(ctx, 1, 4, 1)
	require.NoError(t, err, "При добавлении ошибки быть не должно")

	cartStorage.Reset()
	cart, err = cartStorage.GetItemsByUserID(ctx, 1)
	require.ErrorIs(t, err, model.ErrNotFound, "Корзина еще не существует")
	require.Nil(t, cart, "Данные по корзине не должны вернуться")
}
