package stockservice_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"route256.ozon.ru/project/loms/internal/model"
	"route256.ozon.ru/project/loms/internal/pkg/suite"
	stockservice "route256.ozon.ru/project/loms/internal/service/stock_service"
)

func TestStockInfo(t *testing.T) {

	type test struct {
		Name     string
		Sku      int64
		Qunatity uint16
		Error    error
	}

	tests := []*test{
		{
			Name:  "Стока нет",
			Sku:   1,
			Error: model.ErrNotFound,
		},
		{
			Name:     "Сток есть",
			Sku:      2,
			Qunatity: 10,
		},
	}

	t.Parallel()

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			sp := suite.NewSuiteProvider()

			stockService := stockservice.NewService(
				sp.GetStockStorage(),
			)

			sp.GetStockStorageMock().EXPECT().
				GetBySku(mock.Anything, test.Sku).
				Return(uint16(test.Qunatity), test.Error)

			quantity, err := stockService.Info(context.Background(), test.Sku)
			if test.Error != nil {
				require.NotNil(t, err, "Должна быть ошибка")
				require.ErrorIs(t, err, test.Error, "Не та ошибка")
			} else {
				require.Equal(t, test.Qunatity, quantity, "Не совпало количество")
			}
		})
	}

}
