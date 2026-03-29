package slinky

import (
	"fmt"
	"net/url"
	"strings"
)

// WhatsApp Account: ^https://wa\.me/\+?[0-9]{7,15}/?$

func decodeWhatsAppURL(url *url.URL) (*URL, error) {
	if url.Scheme == "http" {
		url.Scheme = "https"
	}
	if url.Scheme != "https" {
		return nil, fmt.Errorf("%w: invalid WhatsApp scheme", ErrInvalidURL)
	}

	if url.Host != "wa.me" && url.Host != "www.wa.me" {
		return nil, fmt.Errorf("%w: invalid WhatsApp host", ErrInvalidURL)
	}

	path := strings.TrimSuffix(url.Path, "/")
	if len(path) < 1 || path[0] != '/' {
		return nil, fmt.Errorf("%w: invalid WhatsApp path", ErrInvalidURL)
	}

	phoneNumber := strings.TrimPrefix(path, "/")
	if len(phoneNumber) < 7 || len(phoneNumber) > 16 {
		return nil, fmt.Errorf("%w: invalid WhatsApp phone number length", ErrInvalidURL)
	}
	if strings.ContainsFunc(phoneNumber, isNotWhatsAppPhoneNumberRune) {
		return nil, fmt.Errorf("%w: invalid WhatsApp phone number", ErrInvalidURL)
	}

	return &URL{
		Service: WhatsApp,
		Type:    "Account",
		ID:      phoneNumber,
		Data: map[string]string{
			"phoneNumber": phoneNumber,
		},
		URL: url,
	}, nil
}

const whatsAppPhoneNumberAlpha = "0123456789+"

func isNotWhatsAppPhoneNumberRune(r rune) bool {
	return !strings.ContainsRune(whatsAppPhoneNumberAlpha, r)
}
