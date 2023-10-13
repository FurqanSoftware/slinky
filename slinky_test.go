package slinky

import (
	"errors"
	"net/url"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParse(t *testing.T) {
	for _, c := range []struct {
		in      string
		want    *URL
		wantErr error
	}{
		{
			in:   "https://www.facebook.com/hjr265/",
			want: wantWithURL(wantFacebookHjr265, must(url.Parse("https://www.facebook.com/hjr265/"))),
		},
		{
			in:   "https://www.facebook.com/hjr265",
			want: wantWithURL(wantFacebookHjr265, must(url.Parse("https://www.facebook.com/hjr265"))),
		},
		{
			in:   "https://facebook.com/hjr265",
			want: wantWithURL(wantFacebookHjr265, must(url.Parse("https://facebook.com/hjr265"))),
		},
		{
			in:   "https://fb.me/hjr265",
			want: wantWithURL(wantFacebookHjr265, must(url.Parse("https://fb.me/hjr265"))),
		},
		{
			in:   "https://floss.social/@hjr265",
			want: wantWithURL(wantFLOSSSocialHjr265, must(url.Parse("https://floss.social/@hjr265"))),
		},
		{
			in:   "https://fosstodon.org/@hjr265",
			want: wantWithURL(wantFosstodonHjr265, must(url.Parse("https://fosstodon.org/@hjr265"))),
		},
		{
			in:   "https://github.com/hjr265",
			want: wantWithURL(wantGitHubHjr265, must(url.Parse("https://github.com/hjr265"))),
		},
		{
			in:   "https://hjr265.github.io",
			want: wantWithURL(wantGitHubHjr265, must(url.Parse("https://hjr265.github.io"))),
		},
		{
			in:   "https://www.linkedin.com/in/hjr265/",
			want: wantWithURL(wantLinkedInHjr265, must(url.Parse("https://www.linkedin.com/in/hjr265/"))),
		},
		{
			in:   "https://twitter.com/hjr265",
			want: wantWithURL(wantTwitterHjr265, must(url.Parse("https://twitter.com/hjr265"))),
		},
		{
			in:   "https://www.facebook.com/I.AM.KEYBOARDCAT/",
			want: wantWithURL(wantFacebookIAmKeyboardCat, must(url.Parse("https://www.facebook.com/I.AM.KEYBOARDCAT/"))),
		},
		{
			in:   "https://www.facebook.com/profile.php?id=100000000000001",
			want: wantWithURL(wantFacebookIAmKeyboardCatProfileID, must(url.Parse("https://www.facebook.com/profile.php?id=100000000000001"))),
		},
		{
			in:   "https://www.instagram.com/I.AM.KEYBOARDCAT/",
			want: wantWithURL(wantInstagramIAmKeyboardCat, must(url.Parse("https://www.instagram.com/I.AM.KEYBOARDCAT/"))),
		},
		{
			in:   "https://youtube.com/I-AM_KEYBOARDCAT",
			want: wantWithURL(wantYouTubeIAmKeyboardCat, must(url.Parse("https://youtube.com/I-AM_KEYBOARDCAT"))),
		},
		{
			in:   "https://www.youtube.com/I-AM_KEYBOARDCAT",
			want: wantWithURL(wantYouTubeIAmKeyboardCat, must(url.Parse("https://www.youtube.com/I-AM_KEYBOARDCAT"))),
		},
	} {
		t.Run(c.in, func(t *testing.T) {
			got, err := Parse(c.in)
			if c.wantErr != nil {
				if !errors.Is(err, c.wantErr) {
					t.Fatalf("want error %q, got %q", c.wantErr, err)
				}
			}
			if !cmp.Equal(c.want, got) {
				t.Fatal(cmp.Diff(c.want, got))
			}
		})
	}
}

var (
	wantFacebookHjr265 = &URL{
		Service: Facebook,
		Type:    "Profile",
		ID:      "hjr265",
		Data: map[string]string{
			"username": "hjr265",
		},
	}
	wantFLOSSSocialHjr265 = &URL{
		Service: FLOSSSocial,
		Type:    "Profile",
		ID:      "hjr265",
		Data: map[string]string{
			"username": "hjr265",
			"platform": "Mastodon",
		},
	}
	wantFosstodonHjr265 = &URL{
		Service: Fosstodon,
		Type:    "Profile",
		ID:      "hjr265",
		Data: map[string]string{
			"username": "hjr265",
			"platform": "Mastodon",
		},
	}
	wantGitHubHjr265 = &URL{
		Service: GitHub,
		Type:    "User",
		ID:      "hjr265",
		Data: map[string]string{
			"username": "hjr265",
		},
	}
	wantLinkedInHjr265 = &URL{
		Service: LinkedIn,
		Type:    "Profile",
		ID:      "hjr265",
		Data: map[string]string{
			"username": "hjr265",
		},
		URL: must(url.Parse("https://www.linkedin.com/in/hjr265/")),
	}
	wantTwitterHjr265 = &URL{
		Service: Twitter,
		Type:    "Account",
		ID:      "hjr265",
		Data: map[string]string{
			"username": "hjr265",
		},
	}
	wantFacebookIAmKeyboardCat = &URL{
		Service: Facebook,
		Type:    "Profile",
		ID:      "I.AM.KEYBOARDCAT",
		Data: map[string]string{
			"username": "I.AM.KEYBOARDCAT",
		},
	}
	wantFacebookIAmKeyboardCatProfileID = &URL{
		Service: Facebook,
		Type:    "Profile",
		ID:      "100000000000001",
		Data: map[string]string{
			"profileID": "100000000000001",
		},
	}
	wantInstagramIAmKeyboardCat = &URL{
		Service: Instagram,
		Type:    "Profile",
		ID:      "I.AM.KEYBOARDCAT",
		Data: map[string]string{
			"username": "I.AM.KEYBOARDCAT",
		},
	}
	wantYouTubeIAmKeyboardCat = &URL{
		Service: YouTube,
		Type:    "Channel",
		ID:      "I-AM_KEYBOARDCAT",
		Data: map[string]string{
			"channelID": "I-AM_KEYBOARDCAT",
		},
	}
)

func wantWithURL(want *URL, url *url.URL) *URL {
	copy := *want
	copy.URL = url
	return &copy
}

func TestHostPatterns(t *testing.T) {
	for _, c := range []struct {
		host         string
		maxWildcards int
		want         []string
	}{
		{
			host:         "hjr265.github.io",
			maxWildcards: 1,
			want:         []string{"hjr265.github.io", "*.github.io"},
		},
		{
			host:         "keyboard.cat.example.com",
			maxWildcards: 1,
			want:         []string{"keyboard.cat.example.com", "*.cat.example.com"},
		},
		{
			host:         "keyboard.cat.example.com",
			maxWildcards: 2,
			want:         []string{"keyboard.cat.example.com", "*.cat.example.com", "*.*.example.com"},
		},
	} {
		got := hostPatterns(c.host, c.maxWildcards)
		if !reflect.DeepEqual(got, c.want) {
			t.Fatalf("want %v, got %v", c.want, got)
		}
	}
}

func must(url *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	return url
}
