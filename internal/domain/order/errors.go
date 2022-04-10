package order

import "errors"

var ErrOrderBasketEmpty = errors.New("Your basket is empty, please add least 1 product!")
var ErrUserNotAuth = errors.New("Credintials are not matched!")
var ErrOrderCannotBeCanceled = errors.New("You cant cancel this order because it past 14 days")
