package slinky

import (
	"fmt"
	"net/url"
	"strings"
)

// Pinterest Profile: ^https://www\.pinterest.com/u/[A-Za-z0-9_]{3,30}/?$

func decodePinterestURL(url *url.URL) (*URL, error) {
	if url.Scheme == "http" {
		url.Scheme = "https"
	}
	if url.Scheme != "https" {
		return nil, fmt.Errorf("%w: invald Pinterest scheme", ErrInvalidURL)
	}

	if url.Host != "pinterest.com" && url.Host != "www.pinterest.com" {
		return nil, fmt.Errorf("%w: invalid Pinterest host", ErrInvalidURL)
	}

	path := strings.TrimSuffix(url.Path, "/")
	if len(path) < 1 || path[0] != '/' {
		return nil, fmt.Errorf("%w: invalid pinterest path", ErrInvalidURL)
	}

	username := strings.TrimPrefix(path, "/")
	if len(username) < 3 || len(username) > 30 {
		return nil, fmt.Errorf("%w: invalid Pinterest username length", ErrInvalidURL)
	}
	if strings.ContainsFunc(username, isNotPinterestHandleRune) {
		return nil, fmt.Errorf("%w: invalid Pinterest username", ErrInvalidURL)
	}

	return &URL{
		Service: Pinterest,
		Type:    "Profile",
		ID:      username,
		Data: map[string]string{
			"username": username,
		},
		URL: url,
	}, nil
}

const pinterestHandleAlpha = "ABCDEFGHIJKLMONPQRSTUVWXYZabcdefghijklmonpqrstuvwxyz0123456789."

func isNotPinterestHandleRune(r rune) bool {
	return !strings.ContainsRune(pinterestHandleAlpha, r)
}
