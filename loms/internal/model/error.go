package model

import "errors"

var (
	ErrNotFound     = errors.New("not found")
	ErrUnknownError = errors.New("unknown error")
)
