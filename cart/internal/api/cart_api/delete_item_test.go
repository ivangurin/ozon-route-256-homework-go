package cartapi

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"route256.ozon.ru/project/cart/internal/pkg/suite"
)

func TestDeleteItem(t *testing.T) {

	type test struct {
		Name       string
		UserID     int64
		SkuID      int64
		StatusCode int
	}

	tests := []*test{
		{
			Name:       "Не заполнен пользователь",
			UserID:     0,
			SkuID:      1,
			StatusCode: http.StatusBadRequest,
		},
		{
			Name:       "Не заполнен продукт",
			UserID:     1,
			SkuID:      0,
			StatusCode: http.StatusBadRequest,
		},
		{
			Name:       "Заполнен пользователь и продкут",
			UserID:     1,
			SkuID:      1,
			StatusCode: http.StatusNoContent,
		},
		{
			Name:       "Внутренняя ошибка",
			UserID:     1,
			SkuID:      2,
			StatusCode: http.StatusInternalServerError,
		},
	}

	ctx := context.Background()
	sp := suite.NewSuiteProvider(t)
	api := &api{
		cartService: sp.GetCartServiceMock(),
	}

	sp.GetCartServiceMock().DeleteItemMock.
		When(ctx, 1, 1).
		Then(nil)
	sp.GetCartServiceMock().DeleteItemMock.
		When(ctx, 1, 2).
		Then(errors.New("internal error"))

	for _, test := range tests {
		test := test

		t.Run(test.Name, func(t *testing.T) {

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/user/%d/delete/%d", test.UserID, test.SkuID), nil)
			r.SetPathValue(paramUserID, fmt.Sprintf("%d", test.UserID))
			r.SetPathValue(paramSkuID, fmt.Sprintf("%d", test.SkuID))
			api.DeleteItem()(w, r)

			assert.Equal(t, test.StatusCode, w.Result().StatusCode)

		})

	}
}
