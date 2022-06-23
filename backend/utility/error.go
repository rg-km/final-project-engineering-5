package utility

import "errors"

var (
	ErrDatabase = errors.New("ERR_DATABASE")
	ErrNoDataFound = errors.New("ERR_NO_DATA_FOUND")
	ErrBadRequest = errors.New("ERR_BAD_REQUEST")
	ErrUnauthorized = errors.New("ERR_UNAUTHORIZED")
	ErrForbiddedn = errors.New("ERR_FORBIDDEN")
)