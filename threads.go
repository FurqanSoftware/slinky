package slinky

import (
	"fmt"
	"net/url"
	"strings"
)

// Threads Profile: ^https://(www\.)?threads\.net/@[A-Za-z0-9._]{1,30}/?$

func decodeThreadsURL(url *url.URL) (*URL, error) {
	if url.Scheme == "http" {
		url.Scheme = "https"
	}
	if url.Scheme != "https" {
		return nil, fmt.Errorf("%w: invalid Threads scheme", ErrInvalidURL)
	}

	if url.Host != "threads.net" && url.Host != "www.threads.net" {
		return nil, fmt.Errorf("%w: invalid Threads host", ErrInvalidURL)
	}

	path := strings.TrimSuffix(url.Path, "/")
	if !strings.HasPrefix(path, "/@") {
		return nil, fmt.Errorf("%w: invalid Threads path", ErrInvalidURL)
	}

	username := strings.TrimPrefix(path, "/@")
	if len(username) < 1 || len(username) > 30 {
		return nil, fmt.Errorf("%w: invalid Threads username length", ErrInvalidURL)
	}
	if strings.ContainsFunc(username, isNotThreadsHandleRune) {
		return nil, fmt.Errorf("%w: invalid Threads username", ErrInvalidURL)
	}

	return &URL{
		Service: Threads,
		Type:    "Profile",
		ID:      username,
		Data: map[string]string{
			"username": username,
		},
		URL: url,
	}, nil
}

const threadsHandleAlpha = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789._"

func isNotThreadsHandleRune(r rune) bool {
	return !strings.ContainsRune(threadsHandleAlpha, r)
}
