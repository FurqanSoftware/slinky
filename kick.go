package slinky

import (
	"fmt"
	"net/url"
	"strings"
)

// Kick Channel: ^https://(www\.)?kick\.com/[A-Za-z0-9_]{4,25}/?$

func decodeKickURL(url *url.URL) (*URL, error) {
	if url.Scheme == "http" {
		url.Scheme = "https"
	}
	if url.Scheme != "https" {
		return nil, fmt.Errorf("%w: invalid Kick scheme", ErrInvalidURL)
	}

	if url.Host != "kick.com" && url.Host != "www.kick.com" {
		return nil, fmt.Errorf("%w: invalid Kick host", ErrInvalidURL)
	}

	path := strings.TrimSuffix(url.Path, "/")
	if len(path) < 1 || path[0] != '/' {
		return nil, fmt.Errorf("%w: invalid Kick path", ErrInvalidURL)
	}

	username := strings.TrimPrefix(path, "/")
	if len(username) < 4 || len(username) > 25 {
		return nil, fmt.Errorf("%w: invalid Kick username length", ErrInvalidURL)
	}
	if strings.ContainsFunc(username, isNotKickHandleRune) {
		return nil, fmt.Errorf("%w: invalid Kick username", ErrInvalidURL)
	}

	return &URL{
		Service: Kick,
		Type:    "Channel",
		ID:      username,
		Data: map[string]string{
			"username": username,
		},
		URL: url,
	}, nil
}

const kickHandleAlpha = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_"

func isNotKickHandleRune(r rune) bool {
	return !strings.ContainsRune(kickHandleAlpha, r)
}
