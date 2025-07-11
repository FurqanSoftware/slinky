package slinky

import (
	"fmt"
	"net/url"
	"strings"
)

// LinkedIn Profile: ^https://www\.linkedin\.com/in/[^/]{1,100}/?$

func decodeLinkedInURL(url *url.URL) (*URL, error) {
	if url.Scheme == "http" {
		url.Scheme = "https"
	}
	if url.Scheme != "https" {
		return nil, fmt.Errorf("%w: invalid LinkedIn scheme", ErrInvalidURL)
	}

	if url.Host != "linkedin.com" && url.Host != "www.linkedin.com" {
		return nil, fmt.Errorf("%w: invalid LinkedIn host", ErrInvalidURL)
	}

	path := strings.TrimSuffix(url.Path, "/")
	if len(path) < 1 || !strings.HasPrefix(path, "/in/") {
		return nil, fmt.Errorf("%w: invalid LinkedIn path", ErrInvalidURL)
	}

	username := strings.TrimPrefix(path, "/in/")
	if len(username) < 3 || len(username) > 100 {
		return nil, fmt.Errorf("%w: invalid Linkedin username length", ErrInvalidURL)
	}
	if strings.ContainsFunc(username, isNotLinkedInHandleRune) {
		return nil, fmt.Errorf("%w: invalid Linkedin username", ErrInvalidURL)
	}

	return &URL{
		Service: LinkedIn,
		Type:    "Profile",
		ID:      username,
		Data: map[string]string{
			"username": username,
		},
		URL: url,
	}, nil
}

const linkedInHandleAlpha = "ABCDEFGHIJKLMONPQRSTUVWXYZabcdefghijklmonpqrstuvwxyz0123456789."

func isNotLinkedInHandleRune(r rune) bool {
	return !strings.ContainsRune(linkedInHandleAlpha, r)
}
