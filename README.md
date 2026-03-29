# Slinky

[![Go Reference](https://pkg.go.dev/badge/github.com/FurqanSoftware/slinky.svg)](https://pkg.go.dev/github.com/FurqanSoftware/slinky)

Parse social media URLs in Go.

## Usage

``` go
slinky.Parse("https://github.com/hjr265")
// Output:
// 	&URL{
// 		Service: slinky.GitHub,
// 		Type:    "User",
// 		ID:      "hjr265",
// 		Data:    map[string]string{
// 			"username": "hjr265",
// 		},
//	}
```

## URLs Supported

``` go
// Bandcamp
"*.bandcamp.com"

// Behance
"behance.net"
"www.behance.net"

// Bitbucket
"bitbucket.org"

// Bluesky
"bsky.app"

// Codeberg
"codeberg.org"

// DeviantArt
"deviantart.com"
"www.deviantart.com"

// Dribbble
"dribbble.com"
"www.dribbble.com"

// Facebook
"facebook.com"
"www.facebook.com"
"web.facebook.com"
"m.facebook.com"
"fb.me"

// FLOSS.social
"floss.social"

// Fostodon
"fosstodon.org"

// Mastodon
"mastodon.social"

// GitHub
"github.com"
"*.github.io"

// GitLab
"gitlab.com"
"www.gitlab.com"

// Goodreads
"goodreads.com"
"www.goodreads.com"

// Kick
"kick.com"
"www.kick.com"

// Ko-fi
"ko-fi.com"

// Instagram
"instagram.com"
"www.instagram.com"
"m.instagram.com"

// Letterboxd
"letterboxd.com"
"www.letterboxd.com"

// LinkedIn
"linkedin.com"
"www.linkedin.com"

// Medium
"medium.com"
"www.medium.com"

// Messenger
"m.me"
"www.m.me"

// Patreon
"patreon.com"
"www.patreon.com"

// Signal
"signal.me"

// Snapchat
"snapchat.com"
"www.snapchat.com"

// Sourcehut
"sr.ht"

// SoundCloud
"soundcloud.com"
"www.soundcloud.com"

// Spotify
"open.spotify.com"

// Substack
"*.substack.com"

// Reddit
"reddit.com"
"www.reddit.com"
"old.reddit.com"

// Pinterest
"pinterest.com"
"www.pinterest.com"

// Vimeo
"vimeo.com"
"www.vimeo.com"

// TikTok
"tiktok.com"
"www.tiktok.com"

// Threads
"threads.net"
"www.threads.net"

// Steam
"steamcommunity.com"
"www.steamcommunity.com"

// Telegram
"t.me"
"telegram.me"

// Toph
"toph.co"

// Tumblr
"tumblr.com"
"www.tumblr.com"
"*.tumblr.com"

// Twitch
"twitch.tv"
"www.twitch.tv"
"twitch.com"
"www.twitch.com"

// Twitter
"x.com"
"www.x.com"
"twitter.com"
"www.twitter.com"

// WhatsApp
"wa.me"
"www.wa.me"

// YouTube
"youtube.com"
"www.youtube.com"
"m.youtube.com"
```

## Contributing

Contributions are welcome.

## License

Slinky is available under the [BSD (3-Clause) License](https://opensource.org/licenses/BSD-3-Clause).
