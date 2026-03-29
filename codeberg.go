package slinky

import (
	"fmt"
	"net/url"
	"strings"
)

// Codeberg User: ^https://codeberg\.org/[A-Za-z0-9._-]{1,40}/?$

func decodeCodebergURL(url *url.URL) (*URL, error) {
	if url.Scheme == "http" {
		url.Scheme = "https"
	}
	if url.Scheme != "https" {
		return nil, fmt.Errorf("%w: invalid Codeberg scheme", ErrInvalidURL)
	}

	if url.Host != "codeberg.org" {
		return nil, fmt.Errorf("%w: invalid Codeberg host", ErrInvalidURL)
	}

	path := strings.TrimSuffix(url.Path, "/")
	if len(path) < 1 || path[0] != '/' {
		return nil, fmt.Errorf("%w: invalid Codeberg path", ErrInvalidURL)
	}

	username := strings.TrimPrefix(path, "/")
	if len(username) < 1 || len(username) > 40 {
		return nil, fmt.Errorf("%w: invalid Codeberg username length", ErrInvalidURL)
	}
	if strings.ContainsFunc(username, isNotCodebergHandleRune) {
		return nil, fmt.Errorf("%w: invalid Codeberg username", ErrInvalidURL)
	}

	return &URL{
		Service: Codeberg,
		Type:    "User",
		ID:      username,
		Data: map[string]string{
			"username": username,
		},
		URL: url,
	}, nil
}

const codebergHandleAlpha = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789._-"

func isNotCodebergHandleRune(r rune) bool {
	return !strings.ContainsRune(codebergHandleAlpha, r)
}
