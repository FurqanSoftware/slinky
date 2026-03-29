package slinky

import (
	"fmt"
	"net/url"
	"strings"
)

// Snapchat Profile: ^https://(www\.)?snapchat\.com/add/[A-Za-z0-9._-]{3,15}/?$

func decodeSnapchatURL(url *url.URL) (*URL, error) {
	if url.Scheme == "http" {
		url.Scheme = "https"
	}
	if url.Scheme != "https" {
		return nil, fmt.Errorf("%w: invalid Snapchat scheme", ErrInvalidURL)
	}

	if url.Host != "snapchat.com" && url.Host != "www.snapchat.com" {
		return nil, fmt.Errorf("%w: invalid Snapchat host", ErrInvalidURL)
	}

	path := strings.TrimSuffix(url.Path, "/")
	if !strings.HasPrefix(path, "/add/") {
		return nil, fmt.Errorf("%w: invalid Snapchat path", ErrInvalidURL)
	}

	username := strings.TrimPrefix(path, "/add/")
	if len(username) < 3 || len(username) > 15 {
		return nil, fmt.Errorf("%w: invalid Snapchat username length", ErrInvalidURL)
	}
	if strings.ContainsFunc(username, isNotSnapchatHandleRune) {
		return nil, fmt.Errorf("%w: invalid Snapchat username", ErrInvalidURL)
	}

	return &URL{
		Service: Snapchat,
		Type:    "Profile",
		ID:      username,
		Data: map[string]string{
			"username": username,
		},
		URL: url,
	}, nil
}

const snapchatHandleAlpha = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789._-"

func isNotSnapchatHandleRune(r rune) bool {
	return !strings.ContainsRune(snapchatHandleAlpha, r)
}
