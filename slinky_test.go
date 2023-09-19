package slinky

import (
	"net/url"
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
			in:   "https://github.com/hjr265",
			want: wantWithURL(wantGitHubHjr265, must(url.Parse("https://github.com/hjr265"))),
		},
		{
			in:   "https://www.linkedin.com/in/hjr265/",
			want: wantWithURL(wantLinkedInHjr265, must(url.Parse("https://www.linkedin.com/in/hjr265/"))),
		},
		{
			in:   "https://twitter.com/hjr265",
			want: wantWithURL(wantTwitterHjr265, must(url.Parse("https://twitter.com/hjr265"))),
		},
	} {
		t.Run(c.in, func(t *testing.T) {
			got, err := Parse(c.in)
			if !cmp.Equal(c.wantErr, err) {
				t.Fatal(cmp.Diff(c.wantErr, err))
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
)

func wantWithURL(want *URL, url *url.URL) *URL {
	copy := *want
	copy.URL = url
	return &copy
}

func must(url *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	return url
}
