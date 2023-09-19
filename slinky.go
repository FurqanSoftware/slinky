package slinky

import (
	"net/url"
)

// A URL represents a parsed social media URL.
type URL struct {
	Service Service
	Type    string
	ID      string
	Data    map[string]string
	URL     *url.URL
}

// Parse parses a raw url into a URL structure.
//
// The url must be absolute (starting with a scheme).
func Parse(rawURL string) (*URL, error) {
	url, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}
	if !url.IsAbs() {
		return nil, ErrNotAbsolute
	}

	decodeFunc, ok := decodeURLFuncs[url.Host]
	if !ok {
		return nil, ErrUnknownService
	}
	return decodeFunc(url)
}
