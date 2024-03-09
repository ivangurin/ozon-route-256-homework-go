package cartstorage

import "context"

func (cs *cartStorage) DeleteItem(
	ctx context.Context,
	userID int64,
	skuID int64,
) error {

	cart, exists := storage[userID]
	if !exists {
		return nil
	}

	_, exists = cart.Items[skuID]
	if !exists {
		return nil
	}

	delete(cart.Items, skuID)

	if len(cart.Items) == 0 {
		delete(storage, userID)
	}

	return nil

}
