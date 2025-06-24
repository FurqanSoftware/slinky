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
			in:      "https://facebook.com/mahmud.rayed.152111111111111212121312121212123414444/",
			wantErr: ErrInvalidURL,
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
			in:      "https://fosstodon.org/@rayed15211111111",
			wantErr: ErrInvalidURL,
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
			in:   "https://t.me/hjr265",
			want: wantWithURL(wantTelegramHjr265, must(url.Parse("https://t.me/hjr265"))),
		},
		{
			in:      "https://t.me/rayed152152111111111111212121312121212123414444/",
			wantErr: ErrInvalidURL,
		},
		{
			in:   "https://t.me/+100000000000001",
			want: wantWithURL(wantTelegramKeyboardCatPhoneNumber, must(url.Parse("https://t.me/+100000000000001"))),
		},
		{
			in:      "https://t.me/+1000000000000011",
			wantErr: ErrInvalidURL,
		},
		{
			in:   "https://www.x.com/hjr265",
			want: wantWithURL(wantTwitterHjr265, must(url.Parse("https://www.x.com/hjr265"))),
		},
		{
			in:   "https://x.com/hjr265",
			want: wantWithURL(wantTwitterHjr265, must(url.Parse("https://x.com/hjr265"))),
		},
		{
			in:   "https://www.twitter.com/hjr265",
			want: wantWithURL(wantTwitterHjr265, must(url.Parse("https://www.twitter.com/hjr265"))),
		},
		{
			in:   "https://twitter.com/hjr265",
			want: wantWithURL(wantTwitterHjr265, must(url.Parse("https://twitter.com/hjr265"))),
		},
		{
			in:      "https://www.twitter.com/rayed15211111111111111111111111",
			wantErr: ErrInvalidURL,
		},
		{
			in:      "https://twitter.com/rayed15211111111111111111111111",
			wantErr: ErrInvalidURL,
		},
		{
			in:      "https://x.com/rayed15211111111111111111111111",
			wantErr: ErrInvalidURL,
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
			in:      "https://www.instagram.com/rayed15211111111111111111111111/",
			wantErr: ErrInvalidURL,
		},
		{
			in:   "https://youtube.com/I-AM_KEYBOARDCAT",
			want: wantWithURL(wantYouTubeIAmKeyboardCat, must(url.Parse("https://youtube.com/I-AM_KEYBOARDCAT"))),
		},
		{
			in:   "https://www.youtube.com/I-AM_KEYBOARDCAT",
			want: wantWithURL(wantYouTubeIAmKeyboardCat, must(url.Parse("https://www.youtube.com/I-AM_KEYBOARDCAT"))),
		},
		{
			in:   "https://www.reddit.com/user/Acceptable-Mix8356/",
			want: wantWithURL(wantRedditAcceptableMix8356, must(url.Parse("https://www.reddit.com/user/Acceptable-Mix8356/"))),
		},
		{
			in:      "https://www.reddit.com/user/Acceptable-Mix8356111/",
			wantErr: ErrInvalidURL,
		},
		{
			in:   "https://www.reddit.com/r/idk_1_52/",
			want: wantWithURL(wantSubRedditIdk152, must(url.Parse("https://www.reddit.com/r/idk_1_52/"))),
		},
		{
			in:   "https://www.pinterest.com/rayed152/",
			want: wantWithURL(wantPinterestRayed152, must(url.Parse("https://www.pinterest.com/rayed152/"))),
		},
		{
			in:      "https://www.pinterest.com/rayed15211111111111111111111111/",
			wantErr: ErrInvalidURL,
		},
		{
			in:   "https://www.youtube.com/@MahmudRayed",
			want: wantWithURL(wantYouTubeMahmudRayed, must(url.Parse("https://www.youtube.com/@MahmudRayed"))),
		},
		{
			in:      "https://www.youtube.com/@rayed15211111111111111111111111111111111111111111111111",
			wantErr: ErrInvalidURL,
		},
		{
			in:      "https://www.linkedin.com/in/rayed152111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111/",
			wantErr: ErrInvalidURL,
		},
	} {
		t.Run(c.in, func(t *testing.T) {
			got, err := Parse(c.in)
			if c.wantErr != nil {
				if !errors.Is(err, c.wantErr) {
					t.Fatalf("want error %q, got %q", c.wantErr, err)
				}
			} else if err != nil {
				t.Fatal(err)
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
	wantFacebookMahmudrayed152 = &URL{
		Service: Facebook,
		Type:    "Profile",
		ID:      "mahmud.rayed.152111111111111212121312121212123414444",
		Data: map[string]string{
			"username": "mahmud.rayed.152111111111111212121312121212123414444",
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
	wantTelegramHjr265 = &URL{
		Service: Telegram,
		Type:    "Account",
		ID:      "hjr265",
		Data: map[string]string{
			"username": "hjr265",
		},
	}
	wantTelegramKeyboardCatPhoneNumber = &URL{
		Service: Telegram,
		Type:    "Account",
		ID:      "+100000000000001",
		Data: map[string]string{
			"phoneNumber": "+100000000000001",
		},
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
	wantInstagramRayed152 = &URL{
		Service: Instagram,
		Type:    "Profile",
		ID:      "rayed15211111111111111111111111",
		Data: map[string]string{
			"username": "rayed15211111111111111111111111",
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
	wantRedditAcceptableMix8356 = &URL{
		Service: Reddit,
		Type:    "User",
		ID:      "Acceptable-Mix8356",
		Data: map[string]string{
			"username": "Acceptable-Mix8356",
		},
	}
	wantRedditAcceptableMix8356111 = &URL{
		Service: Reddit,
		Type:    "User",
		ID:      "Acceptable-Mix8356111",
		Data: map[string]string{
			"username": "Acceptable-Mix8356111",
		},
	}
	wantSubRedditIdk152 = &URL{
		Service: Reddit,
		Type:    "Subreddit",
		ID:      "idk_1_52",
		Data: map[string]string{
			"username": "idk_1_52",
		},
	}
	wantPinterestRayed152 = &URL{
		Service: Pinterest,
		Type:    "Profile",
		ID:      "rayed152",
		Data: map[string]string{
			"username": "rayed152",
		},
	}
	wantYouTubeMahmudRayed = &URL{
		Service: YouTube,
		Type:    "Channel",
		ID:      "MahmudRayed",
		Data: map[string]string{
			"channelID": "MahmudRayed",
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
