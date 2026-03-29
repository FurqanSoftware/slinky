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
			in:   "https://facebook.com/mahmud.rayed.152",
			want: wantWithURL(wantFacebookMahmudrayed152, must(url.Parse("https://facebook.com/mahmud.rayed.152"))),
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
			in:   "https://www.instagram.com/rayed152",
			want: wantWithURL(wantInstagramRayed152, must(url.Parse("https://www.instagram.com/rayed152"))),
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
			in:   "https://m.me/6585231744937052",
			want: wantWithURL(wantMessengerMahmudRayed, must(url.Parse("https://m.me/6585231744937052"))),
		},
		{
			in:      "https://www.youtube.com/@rayed15211111111111111111111111111111111111111111111111",
			wantErr: ErrInvalidURL,
		},
		{
			in:      "https://www.linkedin.com/in/rayed152111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111/",
			wantErr: ErrInvalidURL,
		},
		{
			in:      "https://www.twitch.com/rayed152111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111/",
			wantErr: ErrInvalidURL,
		},
		{
			in:   "https://www.twitch.com/rayed152/",
			want: wantWithURL(wantTwitchRayed152, must(url.Parse("https://www.twitch.com/rayed152/"))),
		},
		{
			in:   "https://www.snapchat.com/add/hjr265",
			want: wantWithURL(wantSnapchatHjr265, must(url.Parse("https://www.snapchat.com/add/hjr265"))),
		},
		{
			in:   "https://snapchat.com/add/hjr265/",
			want: wantWithURL(wantSnapchatHjr265, must(url.Parse("https://snapchat.com/add/hjr265/"))),
		},
		{
			in:      "https://www.snapchat.com/add/ab",
			wantErr: ErrInvalidURL,
		},
		{
			in:      "https://www.snapchat.com/add/abcdefghijklmnop",
			wantErr: ErrInvalidURL,
		},
		{
			in:      "https://www.snapchat.com/hjr265",
			wantErr: ErrInvalidURL,
		},
		{
			in:   "https://wa.me/1234567890",
			want: wantWithURL(wantWhatsApp1234567890, must(url.Parse("https://wa.me/1234567890"))),
		},
		{
			in:   "https://wa.me/+1234567890/",
			want: wantWithURL(wantWhatsAppPlus1234567890, must(url.Parse("https://wa.me/+1234567890/"))),
		},
		{
			in:      "https://wa.me/123456",
			wantErr: ErrInvalidURL,
		},
		{
			in:      "https://wa.me/12345678901234567",
			wantErr: ErrInvalidURL,
		},
		{
			in:   "https://bsky.app/profile/hjr265.bsky.social",
			want: wantWithURL(wantBlueskyHjr265, must(url.Parse("https://bsky.app/profile/hjr265.bsky.social"))),
		},
		{
			in:   "https://bsky.app/profile/hjr265.bsky.social/",
			want: wantWithURL(wantBlueskyHjr265, must(url.Parse("https://bsky.app/profile/hjr265.bsky.social/"))),
		},
		{
			in:      "https://bsky.app/profile/ab",
			wantErr: ErrInvalidURL,
		},
		{
			in:      "https://bsky.app/hjr265.bsky.social",
			wantErr: ErrInvalidURL,
		},
		{
			in:   "https://www.threads.net/@hjr265",
			want: wantWithURL(wantThreadsHjr265, must(url.Parse("https://www.threads.net/@hjr265"))),
		},
		{
			in:   "https://threads.net/@hjr265/",
			want: wantWithURL(wantThreadsHjr265, must(url.Parse("https://threads.net/@hjr265/"))),
		},
		{
			in:      "https://www.threads.net/hjr265",
			wantErr: ErrInvalidURL,
		},
		{
			in:      "https://www.threads.net/@abcdefghijklmnopqrstuvwxyz12345",
			wantErr: ErrInvalidURL,
		},
		{
			in:   "https://signal.me/#p/+1234567890",
			want: wantWithURL(wantSignal1234567890, must(url.Parse("https://signal.me/#p/+1234567890"))),
		},
		{
			in:      "https://signal.me/#p/123456",
			wantErr: ErrInvalidURL,
		},
		{
			in:      "https://signal.me/#x/+1234567890",
			wantErr: ErrInvalidURL,
		},
		{
			in:   "https://steamcommunity.com/id/hjr265",
			want: wantWithURL(wantSteamHjr265, must(url.Parse("https://steamcommunity.com/id/hjr265"))),
		},
		{
			in:   "https://steamcommunity.com/id/hjr265/",
			want: wantWithURL(wantSteamHjr265, must(url.Parse("https://steamcommunity.com/id/hjr265/"))),
		},
		{
			in:      "https://steamcommunity.com/id/a",
			wantErr: ErrInvalidURL,
		},
		{
			in:      "https://steamcommunity.com/hjr265",
			wantErr: ErrInvalidURL,
		},
		{
			in:   "https://www.tiktok.com/@hjr265",
			want: wantWithURL(wantTikTokHjr265, must(url.Parse("https://www.tiktok.com/@hjr265"))),
		},
		{
			in:   "https://tiktok.com/@hjr265/",
			want: wantWithURL(wantTikTokHjr265, must(url.Parse("https://tiktok.com/@hjr265/"))),
		},
		{
			in:      "https://www.tiktok.com/hjr265",
			wantErr: ErrInvalidURL,
		},
		{
			in:      "https://www.tiktok.com/@abcdefghijklmnopqrstuvwxy",
			wantErr: ErrInvalidURL,
		},
		{
			in:   "https://kick.com/hjr265",
			want: wantWithURL(wantKickHjr265, must(url.Parse("https://kick.com/hjr265"))),
		},
		{
			in:   "https://www.kick.com/hjr265/",
			want: wantWithURL(wantKickHjr265, must(url.Parse("https://www.kick.com/hjr265/"))),
		},
		{
			in:      "https://kick.com/abc",
			wantErr: ErrInvalidURL,
		},
		{
			in:   "https://mastodon.social/@hjr265",
			want: wantWithURL(wantMastodonHjr265, must(url.Parse("https://mastodon.social/@hjr265"))),
		},
		{
			in:   "https://open.spotify.com/user/hjr265",
			want: wantWithURL(wantSpotifyHjr265, must(url.Parse("https://open.spotify.com/user/hjr265"))),
		},
		{
			in:   "https://open.spotify.com/user/hjr265/",
			want: wantWithURL(wantSpotifyHjr265, must(url.Parse("https://open.spotify.com/user/hjr265/"))),
		},
		{
			in:      "https://open.spotify.com/hjr265",
			wantErr: ErrInvalidURL,
		},
		{
			in:   "https://tumblr.com/hjr265",
			want: wantWithURL(wantTumblrHjr265, must(url.Parse("https://tumblr.com/hjr265"))),
		},
		{
			in:   "https://www.tumblr.com/hjr265/",
			want: wantWithURL(wantTumblrHjr265, must(url.Parse("https://www.tumblr.com/hjr265/"))),
		},
		{
			in:   "https://hjr265.tumblr.com",
			want: wantWithURL(wantTumblrHjr265, must(url.Parse("https://hjr265.tumblr.com"))),
		},
		{
			in:   "https://hjr265.tumblr.com/",
			want: wantWithURL(wantTumblrHjr265, must(url.Parse("https://hjr265.tumblr.com/"))),
		},
		{
			in:   "https://gitlab.com/hjr265",
			want: wantWithURL(wantGitLabHjr265, must(url.Parse("https://gitlab.com/hjr265"))),
		},
		{
			in:   "https://www.gitlab.com/hjr265/",
			want: wantWithURL(wantGitLabHjr265, must(url.Parse("https://www.gitlab.com/hjr265/"))),
		},
		{
			in:   "https://bitbucket.org/hjr265",
			want: wantWithURL(wantBitbucketHjr265, must(url.Parse("https://bitbucket.org/hjr265"))),
		},
		{
			in:   "https://bitbucket.org/hjr265/",
			want: wantWithURL(wantBitbucketHjr265, must(url.Parse("https://bitbucket.org/hjr265/"))),
		},
		{
			in:   "https://codeberg.org/hjr265",
			want: wantWithURL(wantCodebergHjr265, must(url.Parse("https://codeberg.org/hjr265"))),
		},
		{
			in:   "https://medium.com/@hjr265",
			want: wantWithURL(wantMediumHjr265, must(url.Parse("https://medium.com/@hjr265"))),
		},
		{
			in:   "https://www.medium.com/@hjr265/",
			want: wantWithURL(wantMediumHjr265, must(url.Parse("https://www.medium.com/@hjr265/"))),
		},
		{
			in:      "https://medium.com/hjr265",
			wantErr: ErrInvalidURL,
		},
		{
			in:   "https://hjr265.substack.com",
			want: wantWithURL(wantSubstackHjr265, must(url.Parse("https://hjr265.substack.com"))),
		},
		{
			in:   "https://hjr265.substack.com/",
			want: wantWithURL(wantSubstackHjr265, must(url.Parse("https://hjr265.substack.com/"))),
		},
		{
			in:   "https://www.patreon.com/hjr265",
			want: wantWithURL(wantPatreonHjr265, must(url.Parse("https://www.patreon.com/hjr265"))),
		},
		{
			in:   "https://patreon.com/hjr265/",
			want: wantWithURL(wantPatreonHjr265, must(url.Parse("https://patreon.com/hjr265/"))),
		},
		{
			in:   "https://ko-fi.com/hjr265",
			want: wantWithURL(wantKofiHjr265, must(url.Parse("https://ko-fi.com/hjr265"))),
		},
		{
			in:      "https://ko-fi.com/ab",
			wantErr: ErrInvalidURL,
		},
		{
			in:   "https://dribbble.com/hjr265",
			want: wantWithURL(wantDribbbleHjr265, must(url.Parse("https://dribbble.com/hjr265"))),
		},
		{
			in:   "https://www.dribbble.com/hjr265/",
			want: wantWithURL(wantDribbbleHjr265, must(url.Parse("https://www.dribbble.com/hjr265/"))),
		},
		{
			in:   "https://www.behance.net/hjr265",
			want: wantWithURL(wantBehanceHjr265, must(url.Parse("https://www.behance.net/hjr265"))),
		},
		{
			in:   "https://behance.net/hjr265/",
			want: wantWithURL(wantBehanceHjr265, must(url.Parse("https://behance.net/hjr265/"))),
		},
		{
			in:   "https://www.deviantart.com/hjr265",
			want: wantWithURL(wantDeviantArtHjr265, must(url.Parse("https://www.deviantart.com/hjr265"))),
		},
		{
			in:   "https://deviantart.com/hjr265/",
			want: wantWithURL(wantDeviantArtHjr265, must(url.Parse("https://deviantart.com/hjr265/"))),
		},
		{
			in:   "https://vimeo.com/hjr265",
			want: wantWithURL(wantVimeoHjr265, must(url.Parse("https://vimeo.com/hjr265"))),
		},
		{
			in:   "https://www.vimeo.com/hjr265/",
			want: wantWithURL(wantVimeoHjr265, must(url.Parse("https://www.vimeo.com/hjr265/"))),
		},
		{
			in:   "https://soundcloud.com/hjr265",
			want: wantWithURL(wantSoundCloudHjr265, must(url.Parse("https://soundcloud.com/hjr265"))),
		},
		{
			in:   "https://www.soundcloud.com/hjr265/",
			want: wantWithURL(wantSoundCloudHjr265, must(url.Parse("https://www.soundcloud.com/hjr265/"))),
		},
		{
			in:   "https://hjr265.bandcamp.com",
			want: wantWithURL(wantBandcampHjr265, must(url.Parse("https://hjr265.bandcamp.com"))),
		},
		{
			in:   "https://hjr265.bandcamp.com/",
			want: wantWithURL(wantBandcampHjr265, must(url.Parse("https://hjr265.bandcamp.com/"))),
		},
		{
			in:   "https://letterboxd.com/hjr265",
			want: wantWithURL(wantLetterboxdHjr265, must(url.Parse("https://letterboxd.com/hjr265"))),
		},
		{
			in:   "https://www.letterboxd.com/hjr265/",
			want: wantWithURL(wantLetterboxdHjr265, must(url.Parse("https://www.letterboxd.com/hjr265/"))),
		},
		{
			in:   "https://www.goodreads.com/user/show/12345678",
			want: wantWithURL(wantGoodreads12345678, must(url.Parse("https://www.goodreads.com/user/show/12345678"))),
		},
		{
			in:   "https://goodreads.com/user/show/12345678/",
			want: wantWithURL(wantGoodreads12345678, must(url.Parse("https://goodreads.com/user/show/12345678/"))),
		},
		{
			in:      "https://goodreads.com/user/12345678",
			wantErr: ErrInvalidURL,
		},
		{
			in:   "https://sr.ht/~hjr265",
			want: wantWithURL(wantSourcehutHjr265, must(url.Parse("https://sr.ht/~hjr265"))),
		},
		{
			in:   "https://sr.ht/~hjr265/",
			want: wantWithURL(wantSourcehutHjr265, must(url.Parse("https://sr.ht/~hjr265/"))),
		},
		{
			in:      "https://sr.ht/~a",
			wantErr: ErrInvalidURL,
		},
		{
			in:      "https://sr.ht/hjr265",
			wantErr: ErrInvalidURL,
		},
		{
			in:   "https://toph.co/u/hjr265",
			want: wantWithURL(wantTophHjr265, must(url.Parse("https://toph.co/u/hjr265"))),
		},
		{
			in:   "https://toph.co/u/hjr265/",
			want: wantWithURL(wantTophHjr265, must(url.Parse("https://toph.co/u/hjr265/"))),
		},
		{
			in:      "https://toph.co/u/abc",
			wantErr: ErrInvalidURL,
		},
		{
			in:      "https://toph.co/u/1hjr265",
			wantErr: ErrInvalidURL,
		},
		{
			in:      "https://toph.co/u/hjr265_",
			wantErr: ErrInvalidURL,
		},
		{
			in:      "https://toph.co/u/hjr.26.5",
			wantErr: ErrInvalidURL,
		},
		{
			in:      "https://toph.co/hjr265",
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
		ID:      "mahmud.rayed.152",
		Data: map[string]string{
			"username": "mahmud.rayed.152",
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
		ID:      "rayed152",
		Data: map[string]string{
			"username": "rayed152",
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
	wantMessengerMahmudRayed = &URL{
		Service: Messenger,
		Type:    "User",
		ID:      "6585231744937052",
		Data: map[string]string{
			"username": "6585231744937052",
		},
	}
	wantSnapchatHjr265 = &URL{
		Service: Snapchat,
		Type:    "Profile",
		ID:      "hjr265",
		Data: map[string]string{
			"username": "hjr265",
		},
	}
	wantSignal1234567890 = &URL{
		Service: Signal,
		Type:    "Account",
		ID:      "+1234567890",
		Data: map[string]string{
			"phoneNumber": "+1234567890",
		},
	}
	wantGitLabHjr265 = &URL{
		Service: GitLab,
		Type:    "User",
		ID:      "hjr265",
		Data: map[string]string{
			"username": "hjr265",
		},
	}
	wantBitbucketHjr265 = &URL{
		Service: Bitbucket,
		Type:    "User",
		ID:      "hjr265",
		Data: map[string]string{
			"username": "hjr265",
		},
	}
	wantCodebergHjr265 = &URL{
		Service: Codeberg,
		Type:    "User",
		ID:      "hjr265",
		Data: map[string]string{
			"username": "hjr265",
		},
	}
	wantMediumHjr265 = &URL{
		Service: Medium,
		Type:    "Profile",
		ID:      "hjr265",
		Data: map[string]string{
			"username": "hjr265",
		},
	}
	wantSubstackHjr265 = &URL{
		Service: Substack,
		Type:    "Publication",
		ID:      "hjr265",
		Data: map[string]string{
			"username": "hjr265",
		},
	}
	wantPatreonHjr265 = &URL{
		Service: Patreon,
		Type:    "Profile",
		ID:      "hjr265",
		Data: map[string]string{
			"username": "hjr265",
		},
	}
	wantKofiHjr265 = &URL{
		Service: Kofi,
		Type:    "Profile",
		ID:      "hjr265",
		Data: map[string]string{
			"username": "hjr265",
		},
	}
	wantDribbbleHjr265 = &URL{
		Service: Dribbble,
		Type:    "Profile",
		ID:      "hjr265",
		Data: map[string]string{
			"username": "hjr265",
		},
	}
	wantBehanceHjr265 = &URL{
		Service: Behance,
		Type:    "Profile",
		ID:      "hjr265",
		Data: map[string]string{
			"username": "hjr265",
		},
	}
	wantDeviantArtHjr265 = &URL{
		Service: DeviantArt,
		Type:    "Profile",
		ID:      "hjr265",
		Data: map[string]string{
			"username": "hjr265",
		},
	}
	wantVimeoHjr265 = &URL{
		Service: Vimeo,
		Type:    "Profile",
		ID:      "hjr265",
		Data: map[string]string{
			"username": "hjr265",
		},
	}
	wantSoundCloudHjr265 = &URL{
		Service: SoundCloud,
		Type:    "Profile",
		ID:      "hjr265",
		Data: map[string]string{
			"username": "hjr265",
		},
	}
	wantBandcampHjr265 = &URL{
		Service: Bandcamp,
		Type:    "Profile",
		ID:      "hjr265",
		Data: map[string]string{
			"username": "hjr265",
		},
	}
	wantLetterboxdHjr265 = &URL{
		Service: Letterboxd,
		Type:    "Profile",
		ID:      "hjr265",
		Data: map[string]string{
			"username": "hjr265",
		},
	}
	wantSourcehutHjr265 = &URL{
		Service: Sourcehut,
		Type:    "User",
		ID:      "hjr265",
		Data: map[string]string{
			"username": "hjr265",
		},
	}
	wantGoodreads12345678 = &URL{
		Service: Goodreads,
		Type:    "Profile",
		ID:      "12345678",
		Data: map[string]string{
			"userID": "12345678",
		},
	}
	wantKickHjr265 = &URL{
		Service: Kick,
		Type:    "Channel",
		ID:      "hjr265",
		Data: map[string]string{
			"username": "hjr265",
		},
	}
	wantMastodonHjr265 = &URL{
		Service: Mastodon,
		Type:    "Profile",
		ID:      "hjr265",
		Data: map[string]string{
			"username": "hjr265",
			"platform": "Mastodon",
		},
	}
	wantSpotifyHjr265 = &URL{
		Service: Spotify,
		Type:    "User",
		ID:      "hjr265",
		Data: map[string]string{
			"username": "hjr265",
		},
	}
	wantTophHjr265 = &URL{
		Service: Toph,
		Type:    "Profile",
		ID:      "hjr265",
		Data: map[string]string{
			"handle": "hjr265",
		},
	}
	wantTumblrHjr265 = &URL{
		Service: Tumblr,
		Type:    "Blog",
		ID:      "hjr265",
		Data: map[string]string{
			"username": "hjr265",
		},
	}
	wantTikTokHjr265 = &URL{
		Service: TikTok,
		Type:    "Profile",
		ID:      "hjr265",
		Data: map[string]string{
			"username": "hjr265",
		},
	}
	wantSteamHjr265 = &URL{
		Service: Steam,
		Type:    "Profile",
		ID:      "hjr265",
		Data: map[string]string{
			"username": "hjr265",
		},
	}
	wantThreadsHjr265 = &URL{
		Service: Threads,
		Type:    "Profile",
		ID:      "hjr265",
		Data: map[string]string{
			"username": "hjr265",
		},
	}
	wantBlueskyHjr265 = &URL{
		Service: Bluesky,
		Type:    "Profile",
		ID:      "hjr265.bsky.social",
		Data: map[string]string{
			"handle": "hjr265.bsky.social",
		},
	}
	wantWhatsApp1234567890 = &URL{
		Service: WhatsApp,
		Type:    "Account",
		ID:      "1234567890",
		Data: map[string]string{
			"phoneNumber": "1234567890",
		},
	}
	wantWhatsAppPlus1234567890 = &URL{
		Service: WhatsApp,
		Type:    "Account",
		ID:      "+1234567890",
		Data: map[string]string{
			"phoneNumber": "+1234567890",
		},
	}
	wantTwitchRayed152 = &URL{
		Service: Twitch,
		Type:    "Channel",
		ID:      "rayed152",
		Data: map[string]string{
			"username": "rayed152",
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
