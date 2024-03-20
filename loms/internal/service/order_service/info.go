package orderservice

import "route256.ozon.ru/project/loms/internal/model"

func (s *service) Info(orderID int64) (*model.Order, error) {
	order, err := s.orderStorage.GetByID(orderID)
	if err != nil {
		return nil, err
	}

	return ToModelOrder(order), nil
}
