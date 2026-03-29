package slinky

import (
	"fmt"
	"net/url"
	"strings"
)

// Patreon Profile: ^https://(www\.)?patreon\.com/[A-Za-z0-9_]{1,64}/?$

func decodePatreonURL(url *url.URL) (*URL, error) {
	if url.Scheme == "http" {
		url.Scheme = "https"
	}
	if url.Scheme != "https" {
		return nil, fmt.Errorf("%w: invalid Patreon scheme", ErrInvalidURL)
	}

	if url.Host != "patreon.com" && url.Host != "www.patreon.com" {
		return nil, fmt.Errorf("%w: invalid Patreon host", ErrInvalidURL)
	}

	path := strings.TrimSuffix(url.Path, "/")
	if len(path) < 1 || path[0] != '/' {
		return nil, fmt.Errorf("%w: invalid Patreon path", ErrInvalidURL)
	}

	username := strings.TrimPrefix(path, "/")
	if len(username) < 1 || len(username) > 64 {
		return nil, fmt.Errorf("%w: invalid Patreon username length", ErrInvalidURL)
	}
	if strings.ContainsFunc(username, isNotPatreonHandleRune) {
		return nil, fmt.Errorf("%w: invalid Patreon username", ErrInvalidURL)
	}

	return &URL{
		Service: Patreon,
		Type:    "Profile",
		ID:      username,
		Data: map[string]string{
			"username": username,
		},
		URL: url,
	}, nil
}

const patreonHandleAlpha = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_"

func isNotPatreonHandleRune(r rune) bool {
	return !strings.ContainsRune(patreonHandleAlpha, r)
}
