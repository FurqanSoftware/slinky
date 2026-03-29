package slinky

import (
	"fmt"
	"net/url"
	"strings"
)

// Ko-fi Profile: ^https://ko-fi\.com/[A-Za-z0-9_]{3,40}/?$

func decodeKofiURL(url *url.URL) (*URL, error) {
	if url.Scheme == "http" {
		url.Scheme = "https"
	}
	if url.Scheme != "https" {
		return nil, fmt.Errorf("%w: invalid Ko-fi scheme", ErrInvalidURL)
	}

	if url.Host != "ko-fi.com" {
		return nil, fmt.Errorf("%w: invalid Ko-fi host", ErrInvalidURL)
	}

	path := strings.TrimSuffix(url.Path, "/")
	if len(path) < 1 || path[0] != '/' {
		return nil, fmt.Errorf("%w: invalid Ko-fi path", ErrInvalidURL)
	}

	username := strings.TrimPrefix(path, "/")
	if len(username) < 3 || len(username) > 40 {
		return nil, fmt.Errorf("%w: invalid Ko-fi username length", ErrInvalidURL)
	}
	if strings.ContainsFunc(username, isNotKofiHandleRune) {
		return nil, fmt.Errorf("%w: invalid Ko-fi username", ErrInvalidURL)
	}

	return &URL{
		Service: Kofi,
		Type:    "Profile",
		ID:      username,
		Data: map[string]string{
			"username": username,
		},
		URL: url,
	}, nil
}

const kofiHandleAlpha = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_"

func isNotKofiHandleRune(r rune) bool {
	return !strings.ContainsRune(kofiHandleAlpha, r)
}
