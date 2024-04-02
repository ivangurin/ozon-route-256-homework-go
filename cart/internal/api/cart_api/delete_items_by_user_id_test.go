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

func TestDeleteItemsByUserID(t *testing.T) {

	type test struct {
		Name       string
		UserID     int64
		StatusCode int
		Error      error
	}

	t.Parallel()

	tests := []*test{
		{
			Name:       "Не заполнен пользователь",
			UserID:     0,
			StatusCode: http.StatusBadRequest,
		},
		{
			Name:       "Заполнен пользователь",
			UserID:     1,
			StatusCode: http.StatusNoContent,
		},
		{
			Name:       "Внутренняя ошибка",
			UserID:     2,
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
			DeleteItemsByUserID(mock.Anything, test.UserID).
			Return(test.Error).
			Once()

		t.Run(test.Name, func(t *testing.T) {

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/user/%d", test.UserID), nil)
			r.SetPathValue(paramUserID, fmt.Sprintf("%d", test.UserID))
			api.DeleteItemsByUserID()(w, r)

			assert.Equal(t, test.StatusCode, w.Result().StatusCode)

		})

	}
}
