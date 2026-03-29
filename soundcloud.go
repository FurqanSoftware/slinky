package slinky

import (
	"fmt"
	"net/url"
	"strings"
)

// SoundCloud Profile: ^https://(www\.)?soundcloud\.com/[A-Za-z0-9_-]{3,25}/?$

func decodeSoundCloudURL(url *url.URL) (*URL, error) {
	if url.Scheme == "http" {
		url.Scheme = "https"
	}
	if url.Scheme != "https" {
		return nil, fmt.Errorf("%w: invalid SoundCloud scheme", ErrInvalidURL)
	}

	if url.Host != "soundcloud.com" && url.Host != "www.soundcloud.com" {
		return nil, fmt.Errorf("%w: invalid SoundCloud host", ErrInvalidURL)
	}

	path := strings.TrimSuffix(url.Path, "/")
	if len(path) < 1 || path[0] != '/' {
		return nil, fmt.Errorf("%w: invalid SoundCloud path", ErrInvalidURL)
	}

	username := strings.TrimPrefix(path, "/")
	if len(username) < 3 || len(username) > 25 {
		return nil, fmt.Errorf("%w: invalid SoundCloud username length", ErrInvalidURL)
	}
	if strings.ContainsFunc(username, isNotSoundCloudHandleRune) {
		return nil, fmt.Errorf("%w: invalid SoundCloud username", ErrInvalidURL)
	}

	return &URL{
		Service: SoundCloud,
		Type:    "Profile",
		ID:      username,
		Data: map[string]string{
			"username": username,
		},
		URL: url,
	}, nil
}

const soundcloudHandleAlpha = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_-"

func isNotSoundCloudHandleRune(r rune) bool {
	return !strings.ContainsRune(soundcloudHandleAlpha, r)
}
