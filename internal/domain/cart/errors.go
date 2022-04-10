package cart

import "errors"

var (
	ErrCartItemQuantityNegative = errors.New("Quantity cannot be negative!")
)
