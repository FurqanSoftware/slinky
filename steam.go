package slinky

import (
	"fmt"
	"net/url"
	"strings"
)

// Steam Profile: ^https://steamcommunity\.com/id/[A-Za-z0-9_-]{2,32}/?$

func decodeSteamURL(url *url.URL) (*URL, error) {
	if url.Scheme == "http" {
		url.Scheme = "https"
	}
	if url.Scheme != "https" {
		return nil, fmt.Errorf("%w: invalid Steam scheme", ErrInvalidURL)
	}

	if url.Host != "steamcommunity.com" && url.Host != "www.steamcommunity.com" {
		return nil, fmt.Errorf("%w: invalid Steam host", ErrInvalidURL)
	}

	path := strings.TrimSuffix(url.Path, "/")
	if !strings.HasPrefix(path, "/id/") {
		return nil, fmt.Errorf("%w: invalid Steam path", ErrInvalidURL)
	}

	username := strings.TrimPrefix(path, "/id/")
	if len(username) < 2 || len(username) > 32 {
		return nil, fmt.Errorf("%w: invalid Steam username length", ErrInvalidURL)
	}
	if strings.ContainsFunc(username, isNotSteamHandleRune) {
		return nil, fmt.Errorf("%w: invalid Steam username", ErrInvalidURL)
	}

	return &URL{
		Service: Steam,
		Type:    "Profile",
		ID:      username,
		Data: map[string]string{
			"username": username,
		},
		URL: url,
	}, nil
}

const steamHandleAlpha = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_-"

func isNotSteamHandleRune(r rune) bool {
	return !strings.ContainsRune(steamHandleAlpha, r)
}
