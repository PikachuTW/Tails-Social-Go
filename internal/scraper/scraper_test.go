package scraper

import (
	"testing"
)

func Test_Match(t *testing.T) {
	facebookScraper := NewFacebookScraper()
	tests := []struct {
		name      string
		scraper   Scraper
		content   string
		wantMatch string
	}{
		{
			name:      "valid Facebook URL",
			scraper:   facebookScraper,
			content:   "https://www.facebook.com/example/posts/123456",
			wantMatch: "https://www.facebook.com/example/posts/123456",
		},
		{
			name:      "Facebook URL with query parameters",
			scraper:   facebookScraper,
			content:   "Look at https://www.facebook.com/watch/?v=123456789",
			wantMatch: "https://www.facebook.com/watch/?v=123456789",
		},
		{
			name:      "Facebook URL in middle of text",
			scraper:   facebookScraper,
			content:   "Before text https://www.facebook.com/page/post after text",
			wantMatch: "https://www.facebook.com/page/post",
		},
		{
			name:      "Facebook URL after text",
			scraper:   facebookScraper,
			content:   "Check this out https://www.facebook.com/example/posts/123456",
			wantMatch: "https://www.facebook.com/example/posts/123456",
		},
		{
			name:      "no match - http instead of https",
			scraper:   facebookScraper,
			content:   "This won't match http://www.facebook.com/example",
			wantMatch: "",
		},
		{
			name:      "no match - different domain",
			scraper:   facebookScraper,
			content:   "Wrong site https://www.twitter.com/example",
			wantMatch: "",
		},
		{
			name:      "no match - empty string",
			scraper:   facebookScraper,
			content:   "",
			wantMatch: "",
		},
		{
			name:      "no match - no URL",
			scraper:   facebookScraper,
			content:   "Just some random text without any URL",
			wantMatch: "",
		},
		{
			name:      "Facebook URL with special characters",
			scraper:   facebookScraper,
			content:   "Check https://www.facebook.com/groups/123/posts/456?comment_id=789&reply_comment_id=999",
			wantMatch: "https://www.facebook.com/groups/123/posts/456?comment_id=789&reply_comment_id=999",
		},
		{
			name:      "multiple URLs - returns first match",
			scraper:   facebookScraper,
			content:   "First https://www.facebook.com/first and second https://www.facebook.com/second",
			wantMatch: "https://www.facebook.com/first",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			gotMatch := test.scraper.Match(test.content)

			if gotMatch != test.wantMatch {
				t.Errorf("Match() = %q, want %q", gotMatch, test.wantMatch)
			}
		})
	}
}

func TestBaseScraper_EmbedColor(t *testing.T) {
	scraper := &BaseScraper{
		embedColor: 0x1877F2,
	}

	if got := scraper.EmbedColor(); got != 0x1877F2 {
		t.Errorf("EmbedColor() = %d, want %d", got, 0x1877F2)
	}
}

func TestBaseScraper_SourceName(t *testing.T) {
	scraper := &BaseScraper{
		sourceName: "TestSource",
	}

	if got := scraper.SourceName(); got != "TestSource" {
		t.Errorf("SourceName() = %q, want %q", got, "TestSource")
	}
}
