package cartapi

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"route256.ozon.ru/project/cart/internal/pkg/suite"
)

func TestDeleteItem(t *testing.T) {

	type test struct {
		Name       string
		UserID     int64
		SkuID      int64
		StatusCode int
		Error      error
	}

	t.Parallel()

	tests := []*test{
		{
			Name:       "Не заполнен пользователь",
			UserID:     0,
			SkuID:      1,
			StatusCode: http.StatusBadRequest,
			Error:      nil,
		},
		{
			Name:       "Не заполнен продукт",
			UserID:     1,
			SkuID:      0,
			StatusCode: http.StatusBadRequest,
			Error:      nil,
		},
		{
			Name:       "Заполнен пользователь и продкут",
			UserID:     1,
			SkuID:      1,
			StatusCode: http.StatusNoContent,
			Error:      nil,
		},
		{
			Name:       "Внутренняя ошибка",
			UserID:     1,
			SkuID:      2,
			StatusCode: http.StatusInternalServerError,
			Error:      errors.New("internal error"),
		},
	}

	sp := suite.NewSuiteProvider()
	api := &api{
		cartService: sp.GetCartServiceMock(),
	}

	for _, test := range tests {

		sp.GetCartServiceMock().EXPECT().
			DeleteItem(mock.Anything, test.UserID, test.SkuID).
			Return(test.Error).
			Once()

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
