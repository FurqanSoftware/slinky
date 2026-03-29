package slinky

import (
	"fmt"
	"net/url"
	"strings"
)

// Vimeo Profile: ^https://(www\.)?vimeo\.com/[A-Za-z0-9_]{1,30}/?$

func decodeVimeoURL(url *url.URL) (*URL, error) {
	if url.Scheme == "http" {
		url.Scheme = "https"
	}
	if url.Scheme != "https" {
		return nil, fmt.Errorf("%w: invalid Vimeo scheme", ErrInvalidURL)
	}

	if url.Host != "vimeo.com" && url.Host != "www.vimeo.com" {
		return nil, fmt.Errorf("%w: invalid Vimeo host", ErrInvalidURL)
	}

	path := strings.TrimSuffix(url.Path, "/")
	if len(path) < 1 || path[0] != '/' {
		return nil, fmt.Errorf("%w: invalid Vimeo path", ErrInvalidURL)
	}

	username := strings.TrimPrefix(path, "/")
	if len(username) < 1 || len(username) > 30 {
		return nil, fmt.Errorf("%w: invalid Vimeo username length", ErrInvalidURL)
	}
	if strings.ContainsFunc(username, isNotVimeoHandleRune) {
		return nil, fmt.Errorf("%w: invalid Vimeo username", ErrInvalidURL)
	}

	return &URL{
		Service: Vimeo,
		Type:    "Profile",
		ID:      username,
		Data: map[string]string{
			"username": username,
		},
		URL: url,
	}, nil
}

const vimeoHandleAlpha = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_"

func isNotVimeoHandleRune(r rune) bool {
	return !strings.ContainsRune(vimeoHandleAlpha, r)
}
