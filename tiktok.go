package slinky

import (
	"fmt"
	"net/url"
	"strings"
)

// TikTok Profile: ^https://(www\.)?tiktok\.com/@[A-Za-z0-9._]{1,24}/?$

func decodeTikTokURL(url *url.URL) (*URL, error) {
	if url.Scheme == "http" {
		url.Scheme = "https"
	}
	if url.Scheme != "https" {
		return nil, fmt.Errorf("%w: invalid TikTok scheme", ErrInvalidURL)
	}

	if url.Host != "tiktok.com" && url.Host != "www.tiktok.com" {
		return nil, fmt.Errorf("%w: invalid TikTok host", ErrInvalidURL)
	}

	path := strings.TrimSuffix(url.Path, "/")
	if !strings.HasPrefix(path, "/@") {
		return nil, fmt.Errorf("%w: invalid TikTok path", ErrInvalidURL)
	}

	username := strings.TrimPrefix(path, "/@")
	if len(username) < 1 || len(username) > 24 {
		return nil, fmt.Errorf("%w: invalid TikTok username length", ErrInvalidURL)
	}
	if strings.ContainsFunc(username, isNotTikTokHandleRune) {
		return nil, fmt.Errorf("%w: invalid TikTok username", ErrInvalidURL)
	}

	return &URL{
		Service: TikTok,
		Type:    "Profile",
		ID:      username,
		Data: map[string]string{
			"username": username,
		},
		URL: url,
	}, nil
}

const tiktokHandleAlpha = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789._"

func isNotTikTokHandleRune(r rune) bool {
	return !strings.ContainsRune(tiktokHandleAlpha, r)
}
