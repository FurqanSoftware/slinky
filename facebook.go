package slinky

import (
	"fmt"
	"net/url"
	"strings"
)

// Facebook Profile/Page: ^https://www\.facebook.com/[A-Za-z0-9.]{1,50}/?$
// Facebook Group: ^https://www\.facebook\.com/groups/[A-Za-z0-9.]{1,50}/?$
// Facebook Profile: ^https://facebook.com/profile.php?id=$
// Facebook Profile: ^https://fb.me/[A-Za-z0-9.]{1,50}$

func decodeFacebookURL(url *url.URL) (*URL, error) {
	if url.Scheme == "http" {
		url.Scheme = "https"
	}
	if url.Scheme != "https" {
		return nil, fmt.Errorf("%w: invalid Facebook scheme", ErrInvalidURL)
	}

	if url.Host != "facebook.com" && url.Host != "www.facebook.com" && url.Host != "web.facebook.com" && url.Host != "fb.me" {
		return nil, fmt.Errorf("%w: invalid Facebook host", ErrInvalidURL)
	}

	path := strings.TrimSuffix(url.Path, "/")
	switch {
	case path == "/profile.php":
		profileID := url.Query().Get("id")

		if strings.ContainsFunc(profileID, isNotFacebookProfileIDRune) {
			return nil, fmt.Errorf("%w: invalid Facebook profile ID", ErrInvalidURL)
		}

		return &URL{
			Service: Facebook,
			Type:    "Profile", // This could be a page as well.
			ID:      profileID,
			Data: map[string]string{
				"profileID": profileID,
			},
			URL: url,
		}, nil

	default:
		if len(path) < 1 || path[0] != '/' {
			return nil, fmt.Errorf("%w: invalid Facebook path", ErrInvalidURL)
		}

		username := strings.TrimPrefix(path, "/")
		if len(username) < 1 || len(username) > 50 {
			return nil, fmt.Errorf("%w: invalid Facebook username length", ErrInvalidURL)
		}
		if strings.ContainsFunc(username, isNotFacebookHandleRune) {
			return nil, fmt.Errorf("%w: invalid Facebook username", ErrInvalidURL)
		}

		return &URL{
			Service: Facebook,
			Type:    "Profile", // This could be a page as well.
			ID:      username,
			Data: map[string]string{
				"username": username,
			},
			URL: url,
		}, nil
	}
}

const facebookHandleAlpha = "ABCDEFGHIJKLMONPQRSTUVWXYZabcdefghijklmonpqrstuvwxyz0123456789."

func isNotFacebookHandleRune(r rune) bool {
	return !strings.ContainsRune(facebookHandleAlpha, r)
}

const facebookProfileIDAlpha = "0123456789"

func isNotFacebookProfileIDRune(r rune) bool {
	return !strings.ContainsRune(facebookProfileIDAlpha, r)
}
