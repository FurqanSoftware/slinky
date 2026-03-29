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
// 		Data:    map[string]string{,
// 			"username": "hjr265",
// 		},
//	}
```

## URLs Supported

``` go
// Bluesky
"bsky.app"

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

// GitHub
"github.com"
"*.github.io"

// Instagram
"instagram.com"
"www.instagram.com"
"m.instagram.com"

// LinkedIn
"linkedin.com"
"www.linkedin.com"

// Messenger
"m.me"
"www.m.me"

// Signal
"signal.me"

// Snapchat
"snapchat.com"
"www.snapchat.com"

// Reddit
"reddit.com"
"www.reddit.com"
"old.reddit.com"

// Pinterest
"pinterest.com"
"www.pinterest.com"

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
