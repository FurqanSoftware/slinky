package slinky

import (
	"fmt"
	"net/url"
	"strings"
)

// Bandcamp Profile: ^https://[A-Za-z0-9-]+\.bandcamp\.com/?$

func decodeBandcampURL(url *url.URL) (*URL, error) {
	if url.Scheme == "http" {
		url.Scheme = "https"
	}
	if url.Scheme != "https" {
		return nil, fmt.Errorf("%w: invalid Bandcamp scheme", ErrInvalidURL)
	}

	if !strings.HasSuffix(url.Host, ".bandcamp.com") {
		return nil, fmt.Errorf("%w: invalid Bandcamp host", ErrInvalidURL)
	}

	username := strings.TrimSuffix(url.Host, ".bandcamp.com")
	if len(username) < 1 || len(username) > 30 {
		return nil, fmt.Errorf("%w: invalid Bandcamp username length", ErrInvalidURL)
	}
	if strings.ContainsFunc(username, isNotBandcampHandleRune) {
		return nil, fmt.Errorf("%w: invalid Bandcamp username", ErrInvalidURL)
	}

	return &URL{
		Service: Bandcamp,
		Type:    "Profile",
		ID:      username,
		Data: map[string]string{
			"username": username,
		},
		URL: url,
	}, nil
}

const bandcampHandleAlpha = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-"

func isNotBandcampHandleRune(r rune) bool {
	return !strings.ContainsRune(bandcampHandleAlpha, r)
}
