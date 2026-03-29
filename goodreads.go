package slinky

import (
	"fmt"
	"net/url"
	"strings"
)

// Goodreads Profile: ^https://(www\.)?goodreads\.com/user/show/[0-9]+/?$

func decodeGoodreadsURL(url *url.URL) (*URL, error) {
	if url.Scheme == "http" {
		url.Scheme = "https"
	}
	if url.Scheme != "https" {
		return nil, fmt.Errorf("%w: invalid Goodreads scheme", ErrInvalidURL)
	}

	if url.Host != "goodreads.com" && url.Host != "www.goodreads.com" {
		return nil, fmt.Errorf("%w: invalid Goodreads host", ErrInvalidURL)
	}

	path := strings.TrimSuffix(url.Path, "/")
	if !strings.HasPrefix(path, "/user/show/") {
		return nil, fmt.Errorf("%w: invalid Goodreads path", ErrInvalidURL)
	}

	userID := strings.TrimPrefix(path, "/user/show/")
	if len(userID) < 1 || len(userID) > 20 {
		return nil, fmt.Errorf("%w: invalid Goodreads user ID length", ErrInvalidURL)
	}
	if strings.ContainsFunc(userID, isNotGoodreadsIDRune) {
		return nil, fmt.Errorf("%w: invalid Goodreads user ID", ErrInvalidURL)
	}

	return &URL{
		Service: Goodreads,
		Type:    "Profile",
		ID:      userID,
		Data: map[string]string{
			"userID": userID,
		},
		URL: url,
	}, nil
}

const goodreadsIDAlpha = "0123456789"

func isNotGoodreadsIDRune(r rune) bool {
	return !strings.ContainsRune(goodreadsIDAlpha, r)
}
