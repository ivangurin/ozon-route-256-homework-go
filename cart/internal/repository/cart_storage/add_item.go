package cartstorage

import "context"

func (cs *cartStorage) AddItem(
	ctx context.Context,
	userID int64,
	skuID int64,
	quantity uint16,
) error {
	cs.Lock()
	defer cs.Unlock()

	cart, exists := storage[userID]
	if !exists {
		cart = &Cart{
			Items: CartItems{},
		}
		storage[userID] = cart
	}

	cartItem, exists := cart.Items[skuID]
	if !exists {
		cartItem = &CartItem{}
		cart.Items[skuID] = cartItem
	}

	cartItem.Quantity += quantity

	return nil
}
