package slinky

import (
	"fmt"
	"net/url"
	"strings"
)

// Messenger Username: ^https://m.me/[A-Za-z0-9.]{1,50}$

func decodeMessengerURL(url *url.URL) (*URL, error) {
	if url.Scheme == "http" {
		url.Scheme = "https"
	}
	if url.Scheme != "https" {
		return nil, fmt.Errorf("%w: invalid Messenger Scheme", ErrInvalidURL)
	}

	if url.Host != "m.me" && url.Host != "www.m.me" {
		return nil, fmt.Errorf("%w: invalid Messenger host", ErrInvalidURL)
	}

	path := strings.TrimSuffix(url.Path, "/")
	if len(path) < 1 || path[0] != '/' {
		return nil, fmt.Errorf("%w: invalid Messenger path", ErrInvalidURL)
	}

	username := strings.TrimPrefix(path, "/")
	if len(username) < 1 || len(username) > 50 {
		return nil, fmt.Errorf("%w: invalid Messenger username length", ErrInvalidURL)
	}
	if strings.ContainsFunc(username, isNotMessengerHandleRune) {
		return nil, fmt.Errorf("%w: invalid Messenger username", ErrInvalidURL)
	}

	return &URL{
		Service: Messenger,
		Type:    "User",
		ID:      username,
		Data: map[string]string{
			"username": username,
		},
		URL: url,
	}, nil
}

const messengerHandleAlpha = "ABCDEFGHIJKLMONPQRSTUVWXYZabcdefghijklmonpqrstuvwxyz0123456789."

func isNotMessengerHandleRune(r rune) bool {
	return !strings.ContainsRune(messengerHandleAlpha, r)
}
