package slinky

import (
	"fmt"
	"net/url"
	"strings"
)

// Medium Profile: ^https://(www\.)?medium\.com/@[A-Za-z0-9._]{1,30}/?$

func decodeMediumURL(url *url.URL) (*URL, error) {
	if url.Scheme == "http" {
		url.Scheme = "https"
	}
	if url.Scheme != "https" {
		return nil, fmt.Errorf("%w: invalid Medium scheme", ErrInvalidURL)
	}

	if url.Host != "medium.com" && url.Host != "www.medium.com" {
		return nil, fmt.Errorf("%w: invalid Medium host", ErrInvalidURL)
	}

	path := strings.TrimSuffix(url.Path, "/")
	if !strings.HasPrefix(path, "/@") {
		return nil, fmt.Errorf("%w: invalid Medium path", ErrInvalidURL)
	}

	username := strings.TrimPrefix(path, "/@")
	if len(username) < 1 || len(username) > 30 {
		return nil, fmt.Errorf("%w: invalid Medium username length", ErrInvalidURL)
	}
	if strings.ContainsFunc(username, isNotMediumHandleRune) {
		return nil, fmt.Errorf("%w: invalid Medium username", ErrInvalidURL)
	}

	return &URL{
		Service: Medium,
		Type:    "Profile",
		ID:      username,
		Data: map[string]string{
			"username": username,
		},
		URL: url,
	}, nil
}

const mediumHandleAlpha = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789._"

func isNotMediumHandleRune(r rune) bool {
	return !strings.ContainsRune(mediumHandleAlpha, r)
}
