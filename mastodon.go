package slinky

import (
	"fmt"
	"net/url"
	"strings"
)

// Mastodon Profile: ^https://mastodon\.com/[A-Za-z0-9_]{1,15}$

func newMastodonURLDecoder(service Service, host string) decodeURLFunc {
	return func(url *url.URL) (*URL, error) {
		if url.Scheme == "http" {
			url.Scheme = "https"
		}
		if url.Scheme != "https" {
			return nil, fmt.Errorf("%w: invalid Mastodon scheme", ErrInvalidURL)
		}

		if url.Host != host {
			return nil, fmt.Errorf("%w: invalid Mastodon host", ErrInvalidURL)
		}

		path := strings.TrimSuffix(url.Path, "/")
		if len(path) < 1 || len(path) > 15 || path[0] != '/' {
			return nil, fmt.Errorf("%w: invalid Mastodon path", ErrInvalidURL)
		}

		username := strings.TrimPrefix(path, "/@")
		if strings.ContainsFunc(username, isNotMastodonHandleRune) {
			return nil, fmt.Errorf("%w: invalid Mastodon username", ErrInvalidURL)
		}

		return &URL{
			Service: service,
			Type:    "Profile",
			ID:      username,
			Data: map[string]string{
				"username": username,
				"platform": "Mastodon",
			},
			URL: url,
		}, nil
	}
}

const mastodonHandleAlpha = "ABCDEFGHIJKLMONPQRSTUVWXYZabcdefghijklmonpqrstuvwxyz0123456789_"

func isNotMastodonHandleRune(r rune) bool {
	return !strings.ContainsRune(mastodonHandleAlpha, r)
}
