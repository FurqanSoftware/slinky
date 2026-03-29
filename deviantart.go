package slinky

import (
	"fmt"
	"net/url"
	"strings"
)

// DeviantArt Profile: ^https://(www\.)?deviantart\.com/[A-Za-z0-9-]{1,20}/?$

func decodeDeviantArtURL(url *url.URL) (*URL, error) {
	if url.Scheme == "http" {
		url.Scheme = "https"
	}
	if url.Scheme != "https" {
		return nil, fmt.Errorf("%w: invalid DeviantArt scheme", ErrInvalidURL)
	}

	if url.Host != "deviantart.com" && url.Host != "www.deviantart.com" {
		return nil, fmt.Errorf("%w: invalid DeviantArt host", ErrInvalidURL)
	}

	path := strings.TrimSuffix(url.Path, "/")
	if len(path) < 1 || path[0] != '/' {
		return nil, fmt.Errorf("%w: invalid DeviantArt path", ErrInvalidURL)
	}

	username := strings.TrimPrefix(path, "/")
	if len(username) < 1 || len(username) > 20 {
		return nil, fmt.Errorf("%w: invalid DeviantArt username length", ErrInvalidURL)
	}
	if strings.ContainsFunc(username, isNotDeviantArtHandleRune) {
		return nil, fmt.Errorf("%w: invalid DeviantArt username", ErrInvalidURL)
	}

	return &URL{
		Service: DeviantArt,
		Type:    "Profile",
		ID:      username,
		Data: map[string]string{
			"username": username,
		},
		URL: url,
	}, nil
}

const deviantartHandleAlpha = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-"

func isNotDeviantArtHandleRune(r rune) bool {
	return !strings.ContainsRune(deviantartHandleAlpha, r)
}
