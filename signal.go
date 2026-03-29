package slinky

import (
	"fmt"
	"net/url"
	"strings"
)

// Signal Account: ^https://signal\.me/#p/\+[0-9]{7,15}$

func decodeSignalURL(url *url.URL) (*URL, error) {
	if url.Scheme == "http" {
		url.Scheme = "https"
	}
	if url.Scheme != "https" {
		return nil, fmt.Errorf("%w: invalid Signal scheme", ErrInvalidURL)
	}

	if url.Host != "signal.me" {
		return nil, fmt.Errorf("%w: invalid Signal host", ErrInvalidURL)
	}

	fragment := url.Fragment
	if !strings.HasPrefix(fragment, "p/") {
		return nil, fmt.Errorf("%w: invalid Signal path", ErrInvalidURL)
	}

	phoneNumber := strings.TrimPrefix(fragment, "p/")
	if len(phoneNumber) < 8 || len(phoneNumber) > 16 {
		return nil, fmt.Errorf("%w: invalid Signal phone number length", ErrInvalidURL)
	}
	if strings.ContainsFunc(phoneNumber, isNotSignalPhoneNumberRune) {
		return nil, fmt.Errorf("%w: invalid Signal phone number", ErrInvalidURL)
	}

	return &URL{
		Service: Signal,
		Type:    "Account",
		ID:      phoneNumber,
		Data: map[string]string{
			"phoneNumber": phoneNumber,
		},
		URL: url,
	}, nil
}

const signalPhoneNumberAlpha = "0123456789+"

func isNotSignalPhoneNumberRune(r rune) bool {
	return !strings.ContainsRune(signalPhoneNumberAlpha, r)
}
