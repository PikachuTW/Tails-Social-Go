package scraper

import (
	"errors"
	"regexp"
	"tails-social-go/internal/util"
)

type YoutubeScraper struct {
	BaseScraper
}

func NewYoutubeScraper() Scraper {
	return &YoutubeScraper{
		BaseScraper{
			sourceName: "Youtube",
			regex:      regexp.MustCompile("https://www\\.youtube\\.com/post/\\S+"),
			embedColor: 0xFF0000,
		},
	}
}

const (
	contentRegex   = `"urlCanonical":"(.*?)","title":"(.*?)","description":"(.*?)"`
	thumbnailRegex = `"thumbnail":{"thumbnails":\[\{"url":"(.*?)"`
)

func (scraper *YoutubeScraper) FetchData(link string) (*FetchedData, error) {
	html, err := util.GetHtml(link)
	if err != nil {
		return nil, err
	}

	contentMatch := regexp.MustCompile(contentRegex).FindStringSubmatch(html)
	if len(contentMatch) < 2 {
		return nil, errors.New("can't find the content matching the regex")
	}
	thumbnailMatch := regexp.MustCompile(thumbnailRegex).FindStringSubmatch(html)
	thumbnail := ""
	if len(thumbnailMatch) >= 2 {
		thumbnail = thumbnailMatch[1]
	}

	return &FetchedData{
		Description: contentMatch[3],
		Title:       contentMatch[2],
		URL:         contentMatch[1],
		Image:       thumbnail,
	}, nil
}
