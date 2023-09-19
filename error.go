package slinky

import "errors"

var (
	ErrNotAbsolute    = errors.New("url is not absolute")
	ErrUnknownService = errors.New("url belongs to an unknown service")
	ErrInvalidURL     = errors.New("invalid URL")
)
