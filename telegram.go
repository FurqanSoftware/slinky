package slinky

import (
	"fmt"
	"net/url"
	"strings"
)

// Telegram Account: ^https://t\.me/[A-Za-z0-9_]{5,32}$
// Telegram Account with Phone Number: ^https://t\.me/+[0-9]{1,15}$

func decodeTelegramURL(url *url.URL) (*URL, error) {
	if url.Scheme == "http" {
		url.Scheme = "https"
	}
	if url.Scheme != "https" {
		return nil, fmt.Errorf("%w: invalid Telegram scheme", ErrInvalidURL)
	}

	if url.Host != "t.me" {
		return nil, fmt.Errorf("%w: invalid Telegram host", ErrInvalidURL)
	}

	path := strings.TrimSuffix(url.Path, "/")
	switch {
	case strings.HasPrefix(path, "/+"):
		if len(path) < 1 || path[0] != '/' {
			return nil, fmt.Errorf("%w: invalid Telegram path", ErrInvalidURL)
		}

		phoneNumber := strings.TrimPrefix(path, "/")
		if strings.ContainsFunc(phoneNumber, isNotTelegramPhoneNumberRune) {
			return nil, fmt.Errorf("%w: invalid Telegram phone number", ErrInvalidURL)
		}

		return &URL{
			Service: Telegram,
			Type:    "Account",
			ID:      phoneNumber,
			Data: map[string]string{
				"phoneNumber": phoneNumber,
			},
			URL: url,
		}, nil

	default:
		if len(path) < 1 || path[0] != '/' {
			return nil, fmt.Errorf("%w: invalid Telegram path", ErrInvalidURL)
		}

		username := strings.TrimPrefix(path, "/")
		if strings.ContainsFunc(username, isNotTelegramHandleRune) {
			return nil, fmt.Errorf("%w: invalid Telegram username", ErrInvalidURL)
		}

		return &URL{
			Service: Telegram,
			Type:    "Account",
			ID:      username,
			Data: map[string]string{
				"username": username,
			},
			URL: url,
		}, nil
	}
}

const telegramHandleAlpha = "ABCDEFGHIJKLMONPQRSTUVWXYZabcdefghijklmonpqrstuvwxyz0123456789_"

func isNotTelegramHandleRune(r rune) bool {
	return !strings.ContainsRune(telegramHandleAlpha, r)
}

const telegramPhoneNumberAlpha = "0123456789+"

func isNotTelegramPhoneNumberRune(r rune) bool {
	return !strings.ContainsRune(telegramPhoneNumberAlpha, r)
}
