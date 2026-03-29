package slinky

import (
	"fmt"
	"net/url"
	"strings"
)

// Toph Profile: ^https://toph\.co/u/[a-zA-Z][a-zA-Z0-9._]+[a-zA-Z0-9]$

func decodeTophURL(url *url.URL) (*URL, error) {
	if url.Scheme == "http" {
		url.Scheme = "https"
	}
	if url.Scheme != "https" {
		return nil, fmt.Errorf("%w: invalid Toph scheme", ErrInvalidURL)
	}

	if url.Host != "toph.co" {
		return nil, fmt.Errorf("%w: invalid Toph host", ErrInvalidURL)
	}

	path := strings.TrimSuffix(url.Path, "/")
	if !strings.HasPrefix(path, "/u/") {
		return nil, fmt.Errorf("%w: invalid Toph path", ErrInvalidURL)
	}

	handle := strings.TrimPrefix(path, "/u/")
	if !isTophHandleValid(handle) {
		return nil, fmt.Errorf("%w: invalid Toph handle", ErrInvalidURL)
	}

	return &URL{
		Service: Toph,
		Type:    "Profile",
		ID:      handle,
		Data: map[string]string{
			"handle": handle,
		},
		URL: url,
	}, nil
}

func isTophHandleValid(handle string) bool {
	if len(handle) < 6 || len(handle) > 20 {
		return false
	}
	r := rune(handle[0])
	if !((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')) {
		return false
	}
	separators := 0
	for _, r := range handle {
		switch {
		case (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9'):
		case r == '.' || r == '_':
			separators++
		default:
			return false
		}
	}
	last := rune(handle[len(handle)-1])
	if last == '.' || last == '_' {
		return false
	}
	return separators <= 1
}
