package shared

import "errors"

var GeneralErrorRequestBodyNotCorrect = errors.New("Check your request body.")
var GeneralErrorRequestParamsNotCorrect = errors.New("Check your params.")
var GeneralServerError = errors.New("We are sorry, we can't process your request")
