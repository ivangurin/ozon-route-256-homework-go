package cartstorage

import "context"

func (s *storage) AddItem(
	ctx context.Context,
	userID int64,
	skuID int64,
	quantity uint16,
) error {
	s.Lock()
	defer s.Unlock()

	cart, exists := cartStorage[userID]
	if !exists {
		cart = &Cart{
			Items: CartItems{},
		}
		cartStorage[userID] = cart
	}

	cartItem, exists := cart.Items[skuID]
	if !exists {
		cartItem = &CartItem{}
		cart.Items[skuID] = cartItem
	}

	cartItem.Quantity += quantity

	return nil
}
