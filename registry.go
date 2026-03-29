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

		// Bandcamp
		"*.bandcamp.com": decodeBandcampURL,

		// Behance
		"behance.net":     decodeBehanceURL,
		"www.behance.net": decodeBehanceURL,

		// Bitbucket
		"bitbucket.org": decodeBitbucketURL,

		// Bluesky
		"bsky.app": decodeBlueskyURL,

		// Codeberg
		"codeberg.org": decodeCodebergURL,

		// DeviantArt
		"deviantart.com":     decodeDeviantArtURL,
		"www.deviantart.com": decodeDeviantArtURL,

		// Dribbble
		"dribbble.com":     decodeDribbbleURL,
		"www.dribbble.com": decodeDribbbleURL,

		// FLOSS.social
		"floss.social": newMastodonURLDecoder(FLOSSSocial, "floss.social"),

		// Fostodon
		"fosstodon.org": newMastodonURLDecoder(Fosstodon, "fosstodon.org"),

		// Mastodon
		"mastodon.social": newMastodonURLDecoder(Mastodon, "mastodon.social"),

		// GitHub
		"github.com":  decodeGitHubURL,
		"*.github.io": decodeGitHubURL,

		// GitLab
		"gitlab.com":     decodeGitLabURL,
		"www.gitlab.com": decodeGitLabURL,

		// Goodreads
		"goodreads.com":     decodeGoodreadsURL,
		"www.goodreads.com": decodeGoodreadsURL,

		// Kick
		"kick.com":     decodeKickURL,
		"www.kick.com": decodeKickURL,

		// Ko-fi
		"ko-fi.com": decodeKofiURL,

		// Instagram
		"instagram.com":       decodeInstagramURL,
		"www.instagram.com":   decodeInstagramURL,
		"m.instagram.com":     decodeInstagramURL,

		// Letterboxd
		"letterboxd.com":     decodeLetterboxdURL,
		"www.letterboxd.com": decodeLetterboxdURL,

		// LinkedIn
		"linkedin.com":     decodeLinkedInURL,
		"www.linkedin.com": decodeLinkedInURL,

		// Medium
		"medium.com":     decodeMediumURL,
		"www.medium.com": decodeMediumURL,

		// Toph
		"toph.co": decodeTophURL,

		// Tumblr
		"tumblr.com":      decodeTumblrURL,
		"www.tumblr.com":  decodeTumblrURL,
		"*.tumblr.com":    decodeTumblrURL,

		// Twitch
		"twitch.tv":      decodeTwitchURL,
		"www.twitch.tv":  decodeTwitchURL,
		"twitch.com":     decodeTwitchURL,
		"www.twitch.com": decodeTwitchURL,

		// Messenger
		"m.me":     decodeMessengerURL,
		"www.m.me": decodeMessengerURL,

		// Patreon
		"patreon.com":     decodePatreonURL,
		"www.patreon.com": decodePatreonURL,

		// Pinterest
		"pinterest.com":     decodePinterestURL,
		"www.pinterest.com": decodePinterestURL,

		// Signal
		"signal.me": decodeSignalURL,

		// Snapchat
		"snapchat.com":     decodeSnapchatURL,
		"www.snapchat.com": decodeSnapchatURL,

		// Sourcehut
		"sr.ht": decodeSourcehutURL,

		// SoundCloud
		"soundcloud.com":     decodeSoundCloudURL,
		"www.soundcloud.com": decodeSoundCloudURL,

		// Steam
		"steamcommunity.com":     decodeSteamURL,
		"www.steamcommunity.com": decodeSteamURL,

		// Spotify
		"open.spotify.com": decodeSpotifyURL,

		// Substack
		"*.substack.com": decodeSubstackURL,

		// Reddit
		"reddit.com":       decodeRedditURL,
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

		// Vimeo
		"vimeo.com":     decodeVimeoURL,
		"www.vimeo.com": decodeVimeoURL,

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
	Bandcamp    Service = "Bandcamp"
	Behance     Service = "Behance"
	Bitbucket   Service = "Bitbucket"
	Bluesky     Service = "Bluesky"
	Codeberg    Service = "Codeberg"
	DeviantArt  Service = "DeviantArt"
	Dribbble    Service = "Dribbble"
	Facebook    Service = "Facebook"
	FLOSSSocial Service = "FLOSSSocial"
	Fosstodon   Service = "Fosstodon"
	GitHub      Service = "GitHub"
	GitLab      Service = "GitLab"
	Goodreads   Service = "Goodreads"
	Instagram   Service = "Instagram"
	Kick        Service = "Kick"
	Kofi        Service = "Kofi"
	Letterboxd  Service = "Letterboxd"
	LinkedIn    Service = "LinkedIn"
	Mastodon    Service = "Mastodon"
	Medium      Service = "Medium"
	Messenger   Service = "Messenger"
	Patreon     Service = "Patreon"
	Pinterest   Service = "Pinterest"
	Reddit      Service = "Reddit"
	Signal      Service = "Signal"
	Snapchat    Service = "Snapchat"
	Sourcehut   Service = "Sourcehut"
	SoundCloud  Service = "SoundCloud"
	Spotify     Service = "Spotify"
	Steam       Service = "Steam"
	Substack    Service = "Substack"
	Telegram    Service = "Telegram"
	Threads     Service = "Threads"
	TikTok      Service = "TikTok"
	Toph        Service = "Toph"
	Tumblr      Service = "Tumblr"
	Twitch      Service = "Twitch"
	Twitter     Service = "Twitter"
	Vimeo       Service = "Vimeo"
	WhatsApp    Service = "WhatsApp"
	YouTube     Service = "YouTube"
)
