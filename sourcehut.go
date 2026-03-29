package slinky

import (
	"fmt"
	"net/url"
	"strings"
)

// Sourcehut User: ^https://sr\.ht/~[A-Za-z0-9_-]{2,30}/?$

func decodeSourcehutURL(url *url.URL) (*URL, error) {
	if url.Scheme == "http" {
		url.Scheme = "https"
	}
	if url.Scheme != "https" {
		return nil, fmt.Errorf("%w: invalid Sourcehut scheme", ErrInvalidURL)
	}

	if url.Host != "sr.ht" {
		return nil, fmt.Errorf("%w: invalid Sourcehut host", ErrInvalidURL)
	}

	path := strings.TrimSuffix(url.Path, "/")
	if !strings.HasPrefix(path, "/~") {
		return nil, fmt.Errorf("%w: invalid Sourcehut path", ErrInvalidURL)
	}

	username := strings.TrimPrefix(path, "/~")
	if len(username) < 2 || len(username) > 30 {
		return nil, fmt.Errorf("%w: invalid Sourcehut username length", ErrInvalidURL)
	}
	if strings.ContainsFunc(username, isNotSourcehutHandleRune) {
		return nil, fmt.Errorf("%w: invalid Sourcehut username", ErrInvalidURL)
	}

	return &URL{
		Service: Sourcehut,
		Type:    "User",
		ID:      username,
		Data: map[string]string{
			"username": username,
		},
		URL: url,
	}, nil
}

const sourcehutHandleAlpha = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_-"

func isNotSourcehutHandleRune(r rune) bool {
	return !strings.ContainsRune(sourcehutHandleAlpha, r)
}
