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

	reserved := false
	reserveErr := s.stockStorage.Reserve(ToStockItems(items))
	if reserveErr != nil {
		logger.Error("failed to reserve quantity for items", reserveErr)
	} else {
		reserved = true
	}

	if reserved {
		err = s.orderStorage.SetStatus(orderID, model.OrederStatusAwatingPayment)
	} else {
		err = s.orderStorage.SetStatus(orderID, model.OrederStatusFailed)
	}
	if err != nil {
		logger.Error("failed to change status", err)
		return 0, fmt.Errorf("failed to change status: %w", err)
	}

	if !reserved {
		return 0, fmt.Errorf("failed to reserve quantity for items, %w", reserveErr)
	}

	return orderID, nil
}
