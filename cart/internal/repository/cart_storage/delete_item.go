package cartstorage

import "context"

func (s *storage) DeleteItem(
	ctx context.Context,
	userID int64,
	skuID int64,
) error {
	s.Lock()
	defer s.Unlock()

	cart, exists := cartStorage[userID]
	if !exists {
		return nil
	}

	_, exists = cart.Items[skuID]
	if !exists {
		return nil
	}

	delete(cart.Items, skuID)

	if len(cart.Items) == 0 {
		delete(cartStorage, userID)
	}

	return nil
}
