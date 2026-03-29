package slinky

import (
	"fmt"
	"net/url"
	"strings"
)

// GitLab User: ^https://(www\.)?gitlab\.com/[A-Za-z0-9._-]{2,255}/?$

func decodeGitLabURL(url *url.URL) (*URL, error) {
	if url.Scheme == "http" {
		url.Scheme = "https"
	}
	if url.Scheme != "https" {
		return nil, fmt.Errorf("%w: invalid GitLab scheme", ErrInvalidURL)
	}

	if url.Host != "gitlab.com" && url.Host != "www.gitlab.com" {
		return nil, fmt.Errorf("%w: invalid GitLab host", ErrInvalidURL)
	}

	path := strings.TrimSuffix(url.Path, "/")
	if len(path) < 1 || path[0] != '/' {
		return nil, fmt.Errorf("%w: invalid GitLab path", ErrInvalidURL)
	}

	username := strings.TrimPrefix(path, "/")
	if len(username) < 2 || len(username) > 255 {
		return nil, fmt.Errorf("%w: invalid GitLab username length", ErrInvalidURL)
	}
	if strings.ContainsFunc(username, isNotGitLabHandleRune) {
		return nil, fmt.Errorf("%w: invalid GitLab username", ErrInvalidURL)
	}

	return &URL{
		Service: GitLab,
		Type:    "User",
		ID:      username,
		Data: map[string]string{
			"username": username,
		},
		URL: url,
	}, nil
}

const gitlabHandleAlpha = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789._-"

func isNotGitLabHandleRune(r rune) bool {
	return !strings.ContainsRune(gitlabHandleAlpha, r)
}
