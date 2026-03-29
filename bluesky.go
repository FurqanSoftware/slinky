package slinky

import (
	"fmt"
	"net/url"
	"strings"
)

// Bluesky Profile: ^https://bsky\.app/profile/[A-Za-z0-9._-]+/?$

func decodeBlueskyURL(url *url.URL) (*URL, error) {
	if url.Scheme == "http" {
		url.Scheme = "https"
	}
	if url.Scheme != "https" {
		return nil, fmt.Errorf("%w: invalid Bluesky scheme", ErrInvalidURL)
	}

	if url.Host != "bsky.app" {
		return nil, fmt.Errorf("%w: invalid Bluesky host", ErrInvalidURL)
	}

	path := strings.TrimSuffix(url.Path, "/")
	if !strings.HasPrefix(path, "/profile/") {
		return nil, fmt.Errorf("%w: invalid Bluesky path", ErrInvalidURL)
	}

	handle := strings.TrimPrefix(path, "/profile/")
	if len(handle) < 3 || len(handle) > 253 {
		return nil, fmt.Errorf("%w: invalid Bluesky handle length", ErrInvalidURL)
	}
	if strings.ContainsFunc(handle, isNotBlueskyHandleRune) {
		return nil, fmt.Errorf("%w: invalid Bluesky handle", ErrInvalidURL)
	}

	return &URL{
		Service: Bluesky,
		Type:    "Profile",
		ID:      handle,
		Data: map[string]string{
			"handle": handle,
		},
		URL: url,
	}, nil
}

const blueskyHandleAlpha = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789.-"

func isNotBlueskyHandleRune(r rune) bool {
	return !strings.ContainsRune(blueskyHandleAlpha, r)
}
