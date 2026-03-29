package slinky

import (
	"fmt"
	"net/url"
	"strings"
)

// Tumblr Blog: ^https://(www\.)?tumblr\.com/[A-Za-z0-9-]{1,32}/?$
// Tumblr Blog: ^https://[A-Za-z0-9-]{1,32}\.tumblr\.com/?$

func decodeTumblrURL(url *url.URL) (*URL, error) {
	if url.Scheme == "http" {
		url.Scheme = "https"
	}
	if url.Scheme != "https" {
		return nil, fmt.Errorf("%w: invalid Tumblr scheme", ErrInvalidURL)
	}

	switch {
	case url.Host == "tumblr.com" || url.Host == "www.tumblr.com":
		path := strings.TrimSuffix(url.Path, "/")
		if len(path) < 1 || path[0] != '/' {
			return nil, fmt.Errorf("%w: invalid Tumblr path", ErrInvalidURL)
		}

		username := strings.TrimPrefix(path, "/")
		if len(username) < 1 || len(username) > 32 {
			return nil, fmt.Errorf("%w: invalid Tumblr username length", ErrInvalidURL)
		}
		if strings.ContainsFunc(username, isNotTumblrHandleRune) {
			return nil, fmt.Errorf("%w: invalid Tumblr username", ErrInvalidURL)
		}

		return &URL{
			Service: Tumblr,
			Type:    "Blog",
			ID:      username,
			Data: map[string]string{
				"username": username,
			},
			URL: url,
		}, nil

	case strings.HasSuffix(url.Host, ".tumblr.com"):
		username := strings.TrimSuffix(url.Host, ".tumblr.com")
		if len(username) < 1 || len(username) > 32 {
			return nil, fmt.Errorf("%w: invalid Tumblr username length", ErrInvalidURL)
		}
		if strings.ContainsFunc(username, isNotTumblrHandleRune) {
			return nil, fmt.Errorf("%w: invalid Tumblr username", ErrInvalidURL)
		}

		return &URL{
			Service: Tumblr,
			Type:    "Blog",
			ID:      username,
			Data: map[string]string{
				"username": username,
			},
			URL: url,
		}, nil

	default:
		return nil, fmt.Errorf("%w: invalid Tumblr host", ErrInvalidURL)
	}
}

const tumblrHandleAlpha = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-"

func isNotTumblrHandleRune(r rune) bool {
	return !strings.ContainsRune(tumblrHandleAlpha, r)
}
