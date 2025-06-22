package slinky

import (
	"fmt"
	"net/url"
	"strings"
)

// Reddit Profile: ^https://www\.reddit.com/u/[A-Za-z0-9_-]{3,20}/?$
// Reddit Profile: ^https://www\.reddit.com/user/[A-Za-z0-9_-]{3,20}/?$
// Reddit Subreddit: ^https://www\.reddit.com/r/[A-Za-z0-9_]{3,20}/?$

func decodeRedditURL(url *url.URL) (*URL, error) {
	if url.Scheme == "http" {
		url.Scheme = "https"
	}
	if url.Scheme != "https" {
		return nil, fmt.Errorf("%w: invald Reddit scheme", ErrInvalidURL)
	}

	if url.Host != "reddit.com" && url.Host != "www.reddit.com" {
		return nil, fmt.Errorf("%w: invalid Reddit host", ErrInvalidURL)
	}

	path := strings.TrimSuffix(url.Path, "/")
	var username string
	var typ string
	switch {
	case strings.HasPrefix(path, "/u/"):
		username = strings.TrimPrefix(path, "/u/")
		typ = "User"
	case strings.HasPrefix(path, "/user/"):
		username = strings.TrimPrefix(path, "/user/")
		typ = "User"
	case strings.HasPrefix(path, "/r/"):
		username = strings.TrimPrefix(path, "/r/")
		typ = "Subreddit"
	default:
		return nil, fmt.Errorf("%w: invalid Reddit path", ErrInvalidURL)
	}

	if strings.ContainsFunc(username, isNotRedditHandleRune) {
		return nil, fmt.Errorf("%w: invalid Reddit username", ErrInvalidURL)
	}

	return &URL{
		Service: Reddit,
		Type:    typ,
		ID:      username,
		Data: map[string]string{
			"username": username,
		},
		URL: url,
	}, nil
}

const redditHandleAlpha = "ABCDEFGHIJKLMONPQRSTUVWXYZabcdefghijklmonpqrstuvwxyz0123456789_-"

func isNotRedditHandleRune(r rune) bool {
	return !strings.ContainsRune(redditHandleAlpha, r)
}
