// Package slinky parses social media URLs and extracts structured information
// such as the service name, profile type, and username or identifier.
//
// It supports a wide range of platforms including Facebook, GitHub, GitLab,
// Instagram, LinkedIn, Mastodon, Reddit, Snapchat, Spotify, TikTok, Twitch,
// Twitter/X, YouTube, and many more.
//
// Usage:
//
//	u, err := slinky.Parse("https://github.com/hjr265")
//	// u.Service == slinky.GitHub
//	// u.Type == "User"
//	// u.ID == "hjr265"
//	// u.Data["username"] == "hjr265"
package slinky
