package business

import "errors"

var (
	ErrInternalServerError = errors.New("internal server error")
	ErrHasBeenModified     = errors.New("data has been modified")
	ErrNotFound            = errors.New("data not found")
	ErrInvalidSpec         = errors.New("given spec is not valid")
)
