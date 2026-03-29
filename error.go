package slinky

import "errors"

var (
	// ErrNotAbsolute is returned when the URL does not have a scheme.
	ErrNotAbsolute = errors.New("url is not absolute")

	// ErrUnknownService is returned when the URL does not match any known social media service.
	ErrUnknownService = errors.New("url belongs to an unknown service")

	// ErrInvalidURL is returned when the URL matches a known service but has an invalid format.
	ErrInvalidURL = errors.New("invalid URL")
)
