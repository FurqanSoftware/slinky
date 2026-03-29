package slinky

import (
	"fmt"
	"net/url"
	"strings"
)

// Substack Publication: ^https://[A-Za-z0-9-]+\.substack\.com/?$

func decodeSubstackURL(url *url.URL) (*URL, error) {
	if url.Scheme == "http" {
		url.Scheme = "https"
	}
	if url.Scheme != "https" {
		return nil, fmt.Errorf("%w: invalid Substack scheme", ErrInvalidURL)
	}

	if !strings.HasSuffix(url.Host, ".substack.com") {
		return nil, fmt.Errorf("%w: invalid Substack host", ErrInvalidURL)
	}

	username := strings.TrimSuffix(url.Host, ".substack.com")
	if len(username) < 1 || len(username) > 30 {
		return nil, fmt.Errorf("%w: invalid Substack username length", ErrInvalidURL)
	}
	if strings.ContainsFunc(username, isNotSubstackHandleRune) {
		return nil, fmt.Errorf("%w: invalid Substack username", ErrInvalidURL)
	}

	return &URL{
		Service: Substack,
		Type:    "Publication",
		ID:      username,
		Data: map[string]string{
			"username": username,
		},
		URL: url,
	}, nil
}

const substackHandleAlpha = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-"

func isNotSubstackHandleRune(r rune) bool {
	return !strings.ContainsRune(substackHandleAlpha, r)
}
