package slinky

import (
	"fmt"
	"net/url"
	"strings"
)

// Behance Profile: ^https://(www\.)?behance\.net/[A-Za-z0-9-]{3,50}/?$

func decodeBehanceURL(url *url.URL) (*URL, error) {
	if url.Scheme == "http" {
		url.Scheme = "https"
	}
	if url.Scheme != "https" {
		return nil, fmt.Errorf("%w: invalid Behance scheme", ErrInvalidURL)
	}

	if url.Host != "behance.net" && url.Host != "www.behance.net" {
		return nil, fmt.Errorf("%w: invalid Behance host", ErrInvalidURL)
	}

	path := strings.TrimSuffix(url.Path, "/")
	if len(path) < 1 || path[0] != '/' {
		return nil, fmt.Errorf("%w: invalid Behance path", ErrInvalidURL)
	}

	username := strings.TrimPrefix(path, "/")
	if len(username) < 3 || len(username) > 50 {
		return nil, fmt.Errorf("%w: invalid Behance username length", ErrInvalidURL)
	}
	if strings.ContainsFunc(username, isNotBehanceHandleRune) {
		return nil, fmt.Errorf("%w: invalid Behance username", ErrInvalidURL)
	}

	return &URL{
		Service: Behance,
		Type:    "Profile",
		ID:      username,
		Data: map[string]string{
			"username": username,
		},
		URL: url,
	}, nil
}

const behanceHandleAlpha = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-"

func isNotBehanceHandleRune(r rune) bool {
	return !strings.ContainsRune(behanceHandleAlpha, r)
}
