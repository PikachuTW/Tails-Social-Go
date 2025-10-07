package scraper

import (
	"regexp"
	"tails-social-go/internal/util"
)

type FacebookScraper struct {
	BaseScraper
}

func NewFacebookScraper() Scraper {
	return &FacebookScraper{
		BaseScraper{
			sourceName: "Facebook",
			regex:      regexp.MustCompile("https://www\\.facebook\\.com/\\S+"),
			embedColor: 0x1877F2,
		},
	}
}

func (scraper *FacebookScraper) FetchData(link string) (*FetchedData, error) {
	doc, err := util.GetDoc(link)
	if err != nil {
		return nil, err
	}

	description, descriptionExists := doc.Find(`meta[property="og:description"]`).Attr("content")
	title, titleExists := doc.Find(`meta[property="og:title"]`).Attr("content")
	image, _ := doc.Find(`meta[property="og:image"]`).Attr("content")
	url, _ := doc.Find(`meta[property="og:url"]`).Attr("content")

	if !descriptionExists || !titleExists {
		redirectUrl, redirectUrlExists := doc.Find(`meta[property="og:url"]`).Attr("content")
		if !redirectUrlExists {
			return nil, nil
		}
		return scraper.FetchData(redirectUrl)
	}
	return &FetchedData{
		Description: description,
		Title:       title,
		URL:         url,
		Image:       image,
	}, nil
}
