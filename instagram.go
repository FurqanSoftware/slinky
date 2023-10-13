package slinky

import (
	"fmt"
	"net/url"
	"strings"
)

// Instagram Profile/Page: ^https://www\.instagram.com/[A-Za-z0-9._]{1,30}/?$

func decodeInstagramURL(url *url.URL) (*URL, error) {
	if url.Scheme == "http" {
		url.Scheme = "https"
	}
	if url.Scheme != "https" {
		return nil, fmt.Errorf("%w: invalid Instagram scheme", ErrInvalidURL)
	}

	if url.Host != "instagram.com" && url.Host != "www.instagram.com" {
		return nil, fmt.Errorf("%w: invalid Instagram host", ErrInvalidURL)
	}

	path := strings.TrimSuffix(url.Path, "/")
	if len(path) < 1 || path[0] != '/' {
		return nil, fmt.Errorf("%w: invalid Instagram path", ErrInvalidURL)
	}

	username := strings.TrimPrefix(path, "/")
	if strings.ContainsFunc(username, isNotInstagramHandleRune) {
		return nil, fmt.Errorf("%w: invalid Instagram username", ErrInvalidURL)
	}

	return &URL{
		Service: Instagram,
		Type:    "Profile",
		ID:      username,
		Data: map[string]string{
			"username": username,
		},
		URL: url,
	}, nil
}

const instagramHandleAlpha = "ABCDEFGHIJKLMONPQRSTUVWXYZabcdefghijklmonpqrstuvwxyz0123456789._"

func isNotInstagramHandleRune(r rune) bool {
	return !strings.ContainsRune(instagramHandleAlpha, r)
}
