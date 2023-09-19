package slinky

import "net/url"

type decodeURLFunc func(*url.URL) (*URL, error)

var (
	decodeURLFuncs = map[string]decodeURLFunc{
		// Facebook
		"facebook.com":     decodeFacebookURL,
		"www.facebook.com": decodeFacebookURL,
		"web.facebook.com": decodeFacebookURL,

		// FLOSS.social
		"floss.social": newMastodonURLDecoder(FLOSSSocial, "floss.social"),

		// Fostodon
		"fosstodon.org": newMastodonURLDecoder(Fosstodon, "fosstodon.org"),

		// GitHub
		"github.com": decodeGitHubURL,

		// LinkedIn
		"linkedin.com":     decodeLinkedInURL,
		"www.linkedin.com": decodeLinkedInURL,

		// Twitter
		"twitter.com": decodeTwitterURL,
	}
)

type Service string

const (
	Facebook    Service = "Facebook"
	FLOSSSocial Service = "FLOSSSocial"
	Fosstodon   Service = "Fosstodon"
	GitHub      Service = "GitHub"
	LinkedIn    Service = "LinkedIn"
	Twitter     Service = "Twitter"
)
