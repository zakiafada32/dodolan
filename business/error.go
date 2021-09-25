package business

import "errors"

var (
	ErrInternalServerError = errors.New("internal_server_error")
	ErrNotFound            = errors.New("data_not_found")
	ErrBadRequest          = errors.New("bad_request")
)
