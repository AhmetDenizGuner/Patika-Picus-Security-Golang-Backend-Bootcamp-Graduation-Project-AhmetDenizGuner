package product

import "errors"

var ErrProductShortSearchKeyword = errors.New("Please use the search keyword that is longer than 3 charachter!")
var ErrProductStockCodeMustBeUnique = errors.New("Stock code field is unique identifier,please try with another stock code")
var ErrProductFieldsMustBePositive = errors.New("Price and stock quantity and categoryID fields must be positive")
var ErrProductStockIsNotEnough = errors.New("There is no enough stock for this product, please update your basket!")
