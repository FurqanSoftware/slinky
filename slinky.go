package slinky

import (
	"net/url"
	"strings"
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

	for _, pattern := range hostPatterns(url.Host, 1) {
		decodeFunc, ok := decodeURLFuncs[pattern]
		if ok {
			return decodeFunc(url)
		}
	}
	return nil, ErrUnknownService
}

func hostPatterns(host string, maxWildcards int) []string {
	parts := strings.Split(host, ".")
	patterns := make([]string, 0, len(parts)-1)
	for i := 0; i <= maxWildcards && i < len(parts)-1; i++ {
		patterns = append(patterns, strings.Repeat("*.", i)+strings.Join(parts[i:], "."))
	}
	return patterns
}
