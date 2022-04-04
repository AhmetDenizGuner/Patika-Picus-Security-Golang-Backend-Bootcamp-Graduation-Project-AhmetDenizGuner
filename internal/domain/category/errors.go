package category

import "errors"

var ErrCategoryDataNotFound = errors.New("Category data cannot be fetched!")
var ErrUploadDataNotFoundOrNotSupported = errors.New("Please check your uploaded media type!")
