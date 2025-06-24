package slinky

import "net/url"

type decodeURLFunc func(*url.URL) (*URL, error)

var (
	decodeURLFuncs = map[string]decodeURLFunc{
		// Facebook
		"facebook.com":     decodeFacebookURL,
		"www.facebook.com": decodeFacebookURL,
		"web.facebook.com": decodeFacebookURL,
		"fb.me":            decodeFacebookURL,

		// FLOSS.social
		"floss.social": newMastodonURLDecoder(FLOSSSocial, "floss.social"),

		// Fostodon
		"fosstodon.org": newMastodonURLDecoder(Fosstodon, "fosstodon.org"),

		// GitHub
		"github.com":  decodeGitHubURL,
		"*.github.io": decodeGitHubURL,

		// Instagram
		"instagram.com":     decodeInstagramURL,
		"www.instagram.com": decodeInstagramURL,

		// LinkedIn
		"linkedin.com":     decodeLinkedInURL,
		"www.linkedin.com": decodeLinkedInURL,

		// Messenger
		"m.me":     decodeMessengerURL,
		"www.m.me": decodeMessengerURL,

		// Pinterest
		"pinterest.com":     decodePinterestURL,
		"www.pinterest.com": decodePinterestURL,

		// Reddit
		"reddit.com":     decodeRedditURL,
		"www.reddit.com": decodeRedditURL,

		// Telegram
		"t.me": decodeTelegramURL,

		// Twitter
		"x.com":           decodeTwitterURL,
		"www.x.com":       decodeTwitterURL,
		"twitter.com":     decodeTwitterURL,
		"www.twitter.com": decodeTwitterURL,

		// YouTube
		"youtube.com":     decodeYouTubeURL,
		"www.youtube.com": decodeYouTubeURL,
	}
)

type Service string

const (
	Facebook    Service = "Facebook"
	FLOSSSocial Service = "FLOSSSocial"
	Fosstodon   Service = "Fosstodon"
	GitHub      Service = "GitHub"
	Instagram   Service = "Instagram"
	LinkedIn    Service = "LinkedIn"
	Messenger   Service = "Messenger"
	Pinterest   Service = "Pinterest"
	Reddit      Service = "Reddit"
	Telegram    Service = "Telegram"
	Twitter     Service = "Twitter"
	YouTube     Service = "YouTube"
)
