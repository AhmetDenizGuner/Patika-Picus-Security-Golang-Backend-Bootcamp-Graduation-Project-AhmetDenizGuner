package cart

import "errors"

var (
	ErrNotFound                 = errors.New("Item not found")
	ErrCustomerCannotBeNil      = errors.New("Customer cannot be nil")
	ErrCartItemQuantityNegative = errors.New("Quantity cannot be negative!")
)
