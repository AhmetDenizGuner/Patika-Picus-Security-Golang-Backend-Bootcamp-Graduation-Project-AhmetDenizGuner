package user

import "errors"

var ErrUserNotAuthorized = errors.New("You are not authorized for this operation!")
var ErrUserCredentialsNotCorrect = errors.New("User not found or password is not correct!")
var ErrUserCheckFormInputs = errors.New("Check your inputs on form")
