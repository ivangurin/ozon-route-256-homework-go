package orderservice

import "route256.ozon.ru/project/loms/internal/model"

func (s *service) Cancel(orderID int64) error {
	orderStorage, err := s.orderStorage.GetByID(orderID)
	if err != nil {
		return err
	}

	order := toModelOrder(orderStorage)

	err = s.stockStorage.CancelReserve(ToStockItems(order.Items))
	if err != nil {
		return err
	}

	err = s.orderStorage.SetStatus(order.ID, model.OrederStatusCanceled)
	if err != nil {
		return err
	}

	return nil
}
