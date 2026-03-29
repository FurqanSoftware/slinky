package slinky

import "net/url"

type decodeURLFunc func(*url.URL) (*URL, error)

var (
	decodeURLFuncs = map[string]decodeURLFunc{
		// Facebook
		"facebook.com":       decodeFacebookURL,
		"www.facebook.com":   decodeFacebookURL,
		"web.facebook.com":   decodeFacebookURL,
		"m.facebook.com":     decodeFacebookURL,
		"fb.me":              decodeFacebookURL,

		// Bluesky
		"bsky.app": decodeBlueskyURL,

		// FLOSS.social
		"floss.social": newMastodonURLDecoder(FLOSSSocial, "floss.social"),

		// Fostodon
		"fosstodon.org": newMastodonURLDecoder(Fosstodon, "fosstodon.org"),

		// GitHub
		"github.com":  decodeGitHubURL,
		"*.github.io": decodeGitHubURL,

		// Instagram
		"instagram.com":       decodeInstagramURL,
		"www.instagram.com":   decodeInstagramURL,
		"m.instagram.com":     decodeInstagramURL,

		// LinkedIn
		"linkedin.com":     decodeLinkedInURL,
		"www.linkedin.com": decodeLinkedInURL,

		// Twitch
		"twitch.tv":      decodeTwitchURL,
		"www.twitch.tv":  decodeTwitchURL,
		"twitch.com":     decodeTwitchURL,
		"www.twitch.com": decodeTwitchURL,

		// Messenger
		"m.me":     decodeMessengerURL,
		"www.m.me": decodeMessengerURL,

		// Pinterest
		"pinterest.com":     decodePinterestURL,
		"www.pinterest.com": decodePinterestURL,

		// Signal
		"signal.me": decodeSignalURL,

		// Snapchat
		"snapchat.com":     decodeSnapchatURL,
		"www.snapchat.com": decodeSnapchatURL,

		// Steam
		"steamcommunity.com":     decodeSteamURL,
		"www.steamcommunity.com": decodeSteamURL,

		// Reddit
		"reddit.com":     decodeRedditURL,
		"www.reddit.com": decodeRedditURL,
		"old.reddit.com": decodeRedditURL,

		// TikTok
		"tiktok.com":     decodeTikTokURL,
		"www.tiktok.com": decodeTikTokURL,

		// Threads
		"threads.net":     decodeThreadsURL,
		"www.threads.net": decodeThreadsURL,

		// Telegram
		"t.me":          decodeTelegramURL,
		"telegram.me":   decodeTelegramURL,

		// Twitter
		"x.com":           decodeTwitterURL,
		"www.x.com":       decodeTwitterURL,
		"twitter.com":     decodeTwitterURL,
		"www.twitter.com": decodeTwitterURL,

		// WhatsApp
		"wa.me":     decodeWhatsAppURL,
		"www.wa.me": decodeWhatsAppURL,

		// YouTube
		"youtube.com":       decodeYouTubeURL,
		"www.youtube.com":   decodeYouTubeURL,
		"m.youtube.com":     decodeYouTubeURL,
	}
)

// Service identifies a social media service.
type Service string

// Supported social media services.
const (
	Bluesky     Service = "Bluesky"
	Facebook    Service = "Facebook"
	FLOSSSocial Service = "FLOSSSocial"
	Fosstodon   Service = "Fosstodon"
	GitHub      Service = "GitHub"
	Instagram   Service = "Instagram"
	LinkedIn    Service = "LinkedIn"
	Twitch      Service = "Twitch"
	Messenger   Service = "Messenger"
	Pinterest   Service = "Pinterest"
	Reddit      Service = "Reddit"
	Signal      Service = "Signal"
	Snapchat    Service = "Snapchat"
	Steam       Service = "Steam"
	Telegram    Service = "Telegram"
	Threads     Service = "Threads"
	TikTok      Service = "TikTok"
	Twitter     Service = "Twitter"
	WhatsApp    Service = "WhatsApp"
	YouTube     Service = "YouTube"
)
