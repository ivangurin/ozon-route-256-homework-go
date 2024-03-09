package model

import "errors"

var (
	ErrNotFound        = errors.New("not found")
	ErrTooManyRequests = errors.New("too many requests")
	ErrUnknownError    = errors.New("unknown error")
)
