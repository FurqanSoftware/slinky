package slinky

import (
	"fmt"
	"net/url"
	"strings"
)

// GitHub User/Company: ^https://github\.com/[A-Za-z0-9\-]{1,39}$

func decodeGitHubURL(url *url.URL) (*URL, error) {
	if url.Scheme == "http" {
		url.Scheme = "https"
	}
	if url.Scheme != "https" {
		return nil, fmt.Errorf("%w: invalid GitHub scheme", ErrInvalidURL)
	}

	switch {
	case url.Host == "github.com":
		path := strings.TrimSuffix(url.Path, "/")
		if len(path) < 1 || path[0] != '/' {
			return nil, fmt.Errorf("%w: invalid GitHub path", ErrInvalidURL)
		}

		username := strings.TrimPrefix(path, "/")
		if strings.ContainsFunc(username, isNotGitHubHandleRune) {
			return nil, fmt.Errorf("%w: invalid GitHub username", ErrInvalidURL)
		}

		return &URL{
			Service: GitHub,
			Type:    "User",
			ID:      username,
			Data: map[string]string{
				"username": username,
			},
			URL: url,
		}, nil

	case strings.HasSuffix(url.Host, ".github.io"):
		username := strings.TrimSuffix(url.Host, ".github.io")
		if strings.ContainsFunc(username, isNotGitHubHandleRune) {
			return nil, fmt.Errorf("%w: invalid GitHub username", ErrInvalidURL)
		}

		return &URL{
			Service: GitHub,
			Type:    "User",
			ID:      username,
			Data: map[string]string{
				"username": username,
			},
			URL: url,
		}, nil

	default:
		return nil, fmt.Errorf("%w: invalid GitHub host", ErrInvalidURL)
	}
}

const githubHandleAlpha = "ABCDEFGHIJKLMONPQRSTUVWXYZabcdefghijklmonpqrstuvwxyz0123456789-"

func isNotGitHubHandleRune(r rune) bool {
	return !strings.ContainsRune(githubHandleAlpha, r)
}
