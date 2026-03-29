package slinky

import (
	"fmt"
	"net/url"
	"strings"
)

// Letterboxd Profile: ^https://(www\.)?letterboxd\.com/[A-Za-z0-9_]{2,15}/?$

func decodeLetterboxdURL(url *url.URL) (*URL, error) {
	if url.Scheme == "http" {
		url.Scheme = "https"
	}
	if url.Scheme != "https" {
		return nil, fmt.Errorf("%w: invalid Letterboxd scheme", ErrInvalidURL)
	}

	if url.Host != "letterboxd.com" && url.Host != "www.letterboxd.com" {
		return nil, fmt.Errorf("%w: invalid Letterboxd host", ErrInvalidURL)
	}

	path := strings.TrimSuffix(url.Path, "/")
	if len(path) < 1 || path[0] != '/' {
		return nil, fmt.Errorf("%w: invalid Letterboxd path", ErrInvalidURL)
	}

	username := strings.TrimPrefix(path, "/")
	if len(username) < 2 || len(username) > 15 {
		return nil, fmt.Errorf("%w: invalid Letterboxd username length", ErrInvalidURL)
	}
	if strings.ContainsFunc(username, isNotLetterboxdHandleRune) {
		return nil, fmt.Errorf("%w: invalid Letterboxd username", ErrInvalidURL)
	}

	return &URL{
		Service: Letterboxd,
		Type:    "Profile",
		ID:      username,
		Data: map[string]string{
			"username": username,
		},
		URL: url,
	}, nil
}

const letterboxdHandleAlpha = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_"

func isNotLetterboxdHandleRune(r rune) bool {
	return !strings.ContainsRune(letterboxdHandleAlpha, r)
}
