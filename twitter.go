package slinky

import (
	"fmt"
	"net/url"
	"strings"
)

// Twitter Account: ^https://twitter\.com/[A-Za-z0-9_]{1,15}$

func decodeTwitterURL(url *url.URL) (*URL, error) {
	if url.Scheme == "http" {
		url.Scheme = "https"
	}
	if url.Scheme != "https" {
		return nil, fmt.Errorf("%w: invalid Twitter scheme", ErrInvalidURL)
	}

	if url.Host != "twitter.com" {
		return nil, fmt.Errorf("%w: invalid Twitter host", ErrInvalidURL)
	}

	path := strings.TrimSuffix(url.Path, "/")
	if len(path) < 1 || len(path) > 15 || path[0] != '/' {
		return nil, fmt.Errorf("%w: invalid Twitter path", ErrInvalidURL)
	}

	username := strings.TrimPrefix(path, "/")
	if strings.ContainsFunc(username, isNotTwitterHandleRune) {
		return nil, fmt.Errorf("%w: invalid Twitter username", ErrInvalidURL)
	}

	return &URL{
		Service: Twitter,
		Type:    "Account",
		ID:      username,
		Data: map[string]string{
			"username": username,
		},
		URL: url,
	}, nil
}

const twitterHandleAlpha = "ABCDEFGHIJKLMONPQRSTUVWXYZabcdefghijklmonpqrstuvwxyz0123456789_"

func isNotTwitterHandleRune(r rune) bool {
	return !strings.ContainsRune(twitterHandleAlpha, r)
}
