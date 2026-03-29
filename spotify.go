package slinky

import (
	"fmt"
	"net/url"
	"strings"
)

// Spotify User: ^https://open\.spotify\.com/user/[A-Za-z0-9._-]{1,30}/?$

func decodeSpotifyURL(url *url.URL) (*URL, error) {
	if url.Scheme == "http" {
		url.Scheme = "https"
	}
	if url.Scheme != "https" {
		return nil, fmt.Errorf("%w: invalid Spotify scheme", ErrInvalidURL)
	}

	if url.Host != "open.spotify.com" {
		return nil, fmt.Errorf("%w: invalid Spotify host", ErrInvalidURL)
	}

	path := strings.TrimSuffix(url.Path, "/")
	if !strings.HasPrefix(path, "/user/") {
		return nil, fmt.Errorf("%w: invalid Spotify path", ErrInvalidURL)
	}

	username := strings.TrimPrefix(path, "/user/")
	if len(username) < 1 || len(username) > 30 {
		return nil, fmt.Errorf("%w: invalid Spotify username length", ErrInvalidURL)
	}
	if strings.ContainsFunc(username, isNotSpotifyHandleRune) {
		return nil, fmt.Errorf("%w: invalid Spotify username", ErrInvalidURL)
	}

	return &URL{
		Service: Spotify,
		Type:    "User",
		ID:      username,
		Data: map[string]string{
			"username": username,
		},
		URL: url,
	}, nil
}

const spotifyHandleAlpha = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789._-"

func isNotSpotifyHandleRune(r rune) bool {
	return !strings.ContainsRune(spotifyHandleAlpha, r)
}
