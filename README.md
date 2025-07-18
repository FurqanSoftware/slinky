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
// Facebook
"facebook.com"
"www.facebook.com"
"web.facebook.com"
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

// LinkedIn
"linkedin.com"
"www.linkedin.com"

// Messenger
"m.me"
"www.m.me"

// Reddit
"reddit.com"
"www.reddit.com"

// Pinterest
"pinterest.com"
"www.pinterest.com"

// Telegram
"t.me"

// Twitter
"x.com"
"www.x.com"
"twitter.com"
"www.twitter.com"

// YouTube
"youtube.com"
"www.youtube.com"
```

## Contributing

Contributions are welcome.

## License

Slinky is available under the [BSD (3-Clause) License](https://opensource.org/licenses/BSD-3-Clause).
