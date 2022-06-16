package utility

import "errors"

var (
	ErrDatabase = errors.New("ERR_DATABASE")
	ErrNoDataFound = errors.New("ERR_NO_DATA_FOUND")
	ErrBadRequest = errors.New("ERR_BAD_REQUEST")
)