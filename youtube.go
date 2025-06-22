package slinky

import (
	"fmt"
	"net/url"
	"strings"
)

// YouTube Profile/Page: ^https://www\.youtube.com/[A-Za-z0-9\-_]{1,50}/?$

func decodeYouTubeURL(url *url.URL) (*URL, error) {
	if url.Scheme == "http" {
		url.Scheme = "https"
	}
	if url.Scheme != "https" {
		return nil, fmt.Errorf("%w: invalid YouTube scheme", ErrInvalidURL)
	}

	if url.Host != "youtube.com" && url.Host != "www.youtube.com" {
		return nil, fmt.Errorf("%w: invalid YouTube host", ErrInvalidURL)
	}

	path := strings.TrimSuffix(url.Path, "/")
	if len(path) < 1 || path[0] != '/' {
		return nil, fmt.Errorf("%w: invalid YouTube path", ErrInvalidURL)
	}

	channel := strings.TrimPrefix(path, "/")
	channel = strings.TrimPrefix(channel, "@")

	if strings.ContainsFunc(channel, isNotYouTubeHandleRune) {
		return nil, fmt.Errorf("%w: invalid YouTube channel ID", ErrInvalidURL)
	}

	return &URL{
		Service: YouTube,
		Type:    "Channel",
		ID:      channel,
		Data: map[string]string{
			"channelID": channel,
		},
		URL: url,
	}, nil
}

const youTubeHandleAlpha = "ABCDEFGHIJKLMONPQRSTUVWXYZabcdefghijklmonpqrstuvwxyz0123456789-_"

func isNotYouTubeHandleRune(r rune) bool {
	return !strings.ContainsRune(youTubeHandleAlpha, r)
}
