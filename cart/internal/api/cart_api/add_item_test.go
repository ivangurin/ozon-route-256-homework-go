package cartapi

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"route256.ozon.ru/project/cart/internal/model"
	"route256.ozon.ru/project/cart/internal/pkg/suite"
)

func TestAddItem(t *testing.T) {

	type test struct {
		Name       string
		UserID     int64
		SkuID      int64
		Quantity   uint16
		Json       string
		StatusCode int
		Error      error
	}

	t.Parallel()

	tests := []*test{
		{
			Name:       "Не заполнен пользователь",
			UserID:     0,
			SkuID:      1,
			Quantity:   1,
			StatusCode: http.StatusBadRequest,
		},
		{
			Name:       "Не заполнен продукт",
			UserID:     1,
			SkuID:      0,
			Quantity:   1,
			StatusCode: http.StatusBadRequest,
		},
		{
			Name:       "Не заполнено количество",
			UserID:     1,
			SkuID:      1,
			Quantity:   0,
			StatusCode: http.StatusBadRequest,
		},
		{
			Name:       "Передан не правильный JSON",
			UserID:     1,
			SkuID:      1,
			Json:       "test",
			StatusCode: http.StatusBadRequest,
		},
		{
			Name:       "Передан не существующий продукт",
			UserID:     1,
			SkuID:      1,
			Quantity:   1,
			StatusCode: http.StatusNotFound,
			Error:      model.ErrNotFound,
		},
		{
			Name:       "Передан существующий продукт",
			UserID:     1,
			SkuID:      2,
			Quantity:   1,
			StatusCode: http.StatusOK,
		},
		{
			Name:       "Внутренняя ошибка",
			UserID:     1,
			SkuID:      3,
			Quantity:   1,
			StatusCode: http.StatusInternalServerError,
			Error:      errors.New("internal error"),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			sp := suite.NewSuiteProvider()

			api := &api{
				cartService: sp.GetCartServiceMock(),
			}

			sp.GetCartServiceMock().EXPECT().
				AddItem(mock.Anything, test.UserID, test.SkuID, test.Quantity).
				Return(test.Error).
				Once()

			var jsonRequest []byte
			var err error
			if test.Json == "" {
				bodyRequest := &AddItemRequestBody{Count: test.Quantity}
				jsonRequest, err = json.Marshal(bodyRequest)
				require.NoError(t, err)
			} else {
				jsonRequest = []byte(test.Json)
			}

			var body bytes.Buffer
			bodyWriter := bufio.NewWriter(&body)
			_, err = bodyWriter.Write(jsonRequest)
			require.NoError(t, err)
			err = bodyWriter.Flush()
			require.NoError(t, err)

			reader := bufio.NewReader(&body)

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/user/%d/cart/%d", test.UserID, test.SkuID), reader)
			r.SetPathValue(paramUserID, fmt.Sprintf("%d", test.UserID))
			r.SetPathValue(paramSkuID, fmt.Sprintf("%d", test.SkuID))
			api.AddItem()(w, r)

			assert.Equal(t, test.StatusCode, w.Result().StatusCode)

		})
	}
}
