package orderservice

import "route256.ozon.ru/project/loms/internal/model"

func (s *service) Pay(orderID int64) error {
	orderStorage, err := s.orderStorage.GetByID(orderID)
	if err != nil {
		return err
	}

	order := ToModelOrder(orderStorage)

	err = s.stockStorage.RemoveReserve(ToStockItems(order.Items))
	if err != nil {
		return err
	}

	err = s.orderStorage.SetStatus(order.ID, model.OrederStatusPayed)
	if err != nil {
		return err
	}

	return nil
}
