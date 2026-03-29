package slinky

import (
	"fmt"
	"net/url"
	"strings"
)

// Bitbucket User: ^https://bitbucket\.org/[A-Za-z0-9_-]{1,30}/?$

func decodeBitbucketURL(url *url.URL) (*URL, error) {
	if url.Scheme == "http" {
		url.Scheme = "https"
	}
	if url.Scheme != "https" {
		return nil, fmt.Errorf("%w: invalid Bitbucket scheme", ErrInvalidURL)
	}

	if url.Host != "bitbucket.org" {
		return nil, fmt.Errorf("%w: invalid Bitbucket host", ErrInvalidURL)
	}

	path := strings.TrimSuffix(url.Path, "/")
	if len(path) < 1 || path[0] != '/' {
		return nil, fmt.Errorf("%w: invalid Bitbucket path", ErrInvalidURL)
	}

	username := strings.TrimPrefix(path, "/")
	if len(username) < 1 || len(username) > 30 {
		return nil, fmt.Errorf("%w: invalid Bitbucket username length", ErrInvalidURL)
	}
	if strings.ContainsFunc(username, isNotBitbucketHandleRune) {
		return nil, fmt.Errorf("%w: invalid Bitbucket username", ErrInvalidURL)
	}

	return &URL{
		Service: Bitbucket,
		Type:    "User",
		ID:      username,
		Data: map[string]string{
			"username": username,
		},
		URL: url,
	}, nil
}

const bitbucketHandleAlpha = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_-"

func isNotBitbucketHandleRune(r rune) bool {
	return !strings.ContainsRune(bitbucketHandleAlpha, r)
}
