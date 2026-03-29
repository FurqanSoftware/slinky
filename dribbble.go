package slinky

import (
	"fmt"
	"net/url"
	"strings"
)

// Dribbble Profile: ^https://(www\.)?dribbble\.com/[A-Za-z0-9_-]{1,30}/?$

func decodeDribbbleURL(url *url.URL) (*URL, error) {
	if url.Scheme == "http" {
		url.Scheme = "https"
	}
	if url.Scheme != "https" {
		return nil, fmt.Errorf("%w: invalid Dribbble scheme", ErrInvalidURL)
	}

	if url.Host != "dribbble.com" && url.Host != "www.dribbble.com" {
		return nil, fmt.Errorf("%w: invalid Dribbble host", ErrInvalidURL)
	}

	path := strings.TrimSuffix(url.Path, "/")
	if len(path) < 1 || path[0] != '/' {
		return nil, fmt.Errorf("%w: invalid Dribbble path", ErrInvalidURL)
	}

	username := strings.TrimPrefix(path, "/")
	if len(username) < 1 || len(username) > 30 {
		return nil, fmt.Errorf("%w: invalid Dribbble username length", ErrInvalidURL)
	}
	if strings.ContainsFunc(username, isNotDribbbleHandleRune) {
		return nil, fmt.Errorf("%w: invalid Dribbble username", ErrInvalidURL)
	}

	return &URL{
		Service: Dribbble,
		Type:    "Profile",
		ID:      username,
		Data: map[string]string{
			"username": username,
		},
		URL: url,
	}, nil
}

const dribbbleHandleAlpha = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_-"

func isNotDribbbleHandleRune(r rune) bool {
	return !strings.ContainsRune(dribbbleHandleAlpha, r)
}
