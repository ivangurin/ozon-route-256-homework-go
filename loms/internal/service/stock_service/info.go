package stockservice

func (s *service) Info(sku int64) (uint16, error) {
	quantity, err := s.stockStorage.GetBySku(sku)
	if err != nil {
		return 0, err
	}

	return quantity, nil
}
