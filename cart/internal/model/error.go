package model

import "errors"

var (
	ErrNotFound         = errors.New("not found")
	ErrTooManyRequests  = errors.New("too many requests")
	ErrInsufficientSock = errors.New("insufficient stock")
	ErrUnknownError     = errors.New("unknown error")
)
