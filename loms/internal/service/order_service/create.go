package orderservice

import (
	"fmt"

	"route256.ozon.ru/project/loms/internal/model"
	"route256.ozon.ru/project/loms/internal/pkg/logger"
)

func (s *service) Create(user int64, items model.OrderItems) (int64, error) {

	orderID, err := s.orderStorage.Create(user, ToOrderStorageItems(items))
	if err != nil {
		logger.Error("faild to create order", err)
		return 0, fmt.Errorf("faild to create order: %w", err)
	}

	reserveFailed := false
	reservErr := s.stockStorage.Reserve(ToStockItems(items))
	if reservErr != nil {
		logger.Error("failed to reserve quantity for items", reservErr)
		reserveFailed = true
	}

	if reserveFailed {
		err = s.orderStorage.SetStatus(orderID, model.OrederStatusFailed)
		if err == nil {
			return 0, fmt.Errorf("failed to reserve quantity for items: %w", reservErr)
		}
	} else {
		err = s.orderStorage.SetStatus(orderID, model.OrederStatusAwatingPayment)
	}
	if err != nil {
		logger.Error("failed to change status", err)
		return 0, fmt.Errorf("failed to change status: %w", err)
	}

	return orderID, nil
}
