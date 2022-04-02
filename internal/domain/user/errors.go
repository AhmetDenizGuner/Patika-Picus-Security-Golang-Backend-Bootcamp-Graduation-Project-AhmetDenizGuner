package user

import "errors"

var ErrUserNotAuthorized = errors.New("You are not authorized for this operation!")
