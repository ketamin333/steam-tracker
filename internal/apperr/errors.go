package apperr

import "errors"

var (
	ErrInvalidBody   = errors.New("invalid request body")
	ErrInvalidParams = errors.New("invalid params")
	ErrNotFound      = errors.New("not found")
	ErrUnauthorized  = errors.New("unauthorized")
	ErrAlreadyExists = errors.New("already exists")
	ErrInternal      = errors.New("server error")
)
