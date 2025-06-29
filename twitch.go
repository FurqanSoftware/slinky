package slinky

import (
	"fmt"
	"net/url"
	"strings"
)

// twtich Channel: ^https://www\.twitch.com/u/[A-Za-z0-9_]{4,25}/?$

func decodeTwitchURL(url *url.URL) (*URL, error) {
	if url.Scheme == "http" {
		url.Scheme = "https"
	}
	if url.Scheme != "https" {
		return nil, fmt.Errorf("%w: invalid Twitch scheme", ErrInvalidURL)
	}

	if url.Host != "twitch.com" && url.Host != "www.twitch.com" {
		return nil, fmt.Errorf("%w: invalid Twitch host", ErrInvalidURL)
	}

	path := strings.TrimSuffix(url.Path, "/")
	if len(path) < 1 || path[0] != '/' {
		return nil, fmt.Errorf("%w: invalid Twitch path", ErrInvalidURL)
	}

	username := strings.TrimPrefix(path, "/")
	if len(username) < 4 || len(username) > 25 {
		return nil, fmt.Errorf("%w: invalid Twitch username length", ErrInvalidURL)
	}
	if strings.ContainsFunc(username, isNotTwitchHandleRune) {
		return nil, fmt.Errorf("%w: invalid Twitch username", ErrInvalidURL)
	}

	return &URL{
		Service: Twitch,
		Type:    "Channel",
		ID:      username,
		Data: map[string]string{
			"username": username,
		},
		URL: url,
	}, nil
}

const twitchHandleAlpha = "ABCDEFGHIJKLMONPQRSTUVWXYZabcdefghijklmonpqrstuvwxyz0123456789_"

func isNotTwitchHandleRune(r rune) bool {
	return !strings.ContainsRune(twitchHandleAlpha, r)
}
