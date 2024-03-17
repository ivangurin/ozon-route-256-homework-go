package cartapi

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-test/deep"
	"github.com/stretchr/testify/require"
	"route256.ozon.ru/project/cart/internal/model"
	"route256.ozon.ru/project/cart/internal/pkg/suite"
	cartservice "route256.ozon.ru/project/cart/internal/service/cart_service"
)

func TestGetItemsByUserID(t *testing.T) {

	type test struct {
		Name       string
		UserID     int64
		Cart       *cartservice.Cart
		StatusCode int
		Error      error
	}

	tests := []*test{
		{
			Name:       "Не заполнен пользователь",
			UserID:     0,
			StatusCode: http.StatusBadRequest,
		},
		{
			Name:       "Корзина не найдена",
			UserID:     1,
			StatusCode: http.StatusNotFound,
			Error:      model.ErrNotFound,
		},
		{
			Name:       "Корзина найдена",
			UserID:     2,
			StatusCode: http.StatusOK,
			Cart: &cartservice.Cart{
				Items: []*cartservice.CartItem{
					{
						SkuID:    1,
						Name:     "Продукт 1",
						Quantity: 1,
						Price:    1,
					},
					{
						SkuID:    2,
						Name:     "Продукт 2",
						Quantity: 2,
						Price:    2,
					},
				},
				TotalPrice: 5,
			},
		},
		{
			Name:       "Внутренняя ошибка",
			UserID:     3,
			StatusCode: http.StatusInternalServerError,
			Error:      errors.New("internal error"),
		},
	}

	ctx := context.Background()
	sp := suite.NewSuiteProvider(t)
	api := &api{
		cartService: sp.GetCartServiceMock(),
	}

	for _, test := range tests {

		sp.GetCartServiceMock().GetItemsByUserIDMock.
			When(ctx, test.UserID).
			Then(test.Cart, test.Error)

		t.Run(test.Name, func(t *testing.T) {

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/user/%d/cart", test.UserID), nil)
			r.SetPathValue(paramUserID, fmt.Sprintf("%d", test.UserID))
			api.GetItemsByUserID()(w, r)

			require.Equal(t, test.StatusCode, w.Result().StatusCode)

			if w.Result().StatusCode != http.StatusOK {
				return
			}

			body, err := io.ReadAll(w.Body)
			require.NoError(t, err, "Ошибка при получение body ответа")
			require.Greater(t, len(body), 0, "Вернулся пустое тело ответа")

			cart := &cartservice.Cart{}
			err = json.Unmarshal(body, cart)
			require.NoError(t, err, "Ошибка при unmarshal тела ответа")

			diff := deep.Equal(test.Cart, cart)
			if diff != nil {
				t.Errorf("Корзины должны совпадать: %+v", diff)
			}

		})

	}
}
