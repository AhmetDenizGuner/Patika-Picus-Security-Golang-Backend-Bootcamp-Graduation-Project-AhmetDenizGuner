package product

import "errors"

var ErrProductShortSearchKeyword = errors.New("Please use the search keyword that is longer than 3 charachter!")
var ErrProductStockCodeMustBeUnique = errors.New("Stock code field is unique identifier,please try with another stock code")
var ErrProductFieldsMustBePositive = errors.New("Price and stock quantity fields must be positive")
