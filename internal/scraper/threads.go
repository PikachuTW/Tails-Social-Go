package scraper

import (
	"regexp"
	"tails-social-go/internal/util"
)

type ThreadsScraper struct {
	BaseScraper
}

func NewThreadsScraper() Scraper {
	return &ThreadsScraper{
		BaseScraper{
			sourceName: "Threads",
			regex:      regexp.MustCompile("https://www\\.threads\\.com/\\S+"),
			embedColor: 0,
		},
	}
}

func (scraper *ThreadsScraper) FetchData(link string) (*FetchedData, error) {
	doc, err := util.GetDoc(link)
	if err != nil {
		return nil, err
	}

	description, descriptionExists := doc.Find(`meta[property="og:description"]`).Attr("content")
	title, titleExists := doc.Find(`meta[property="og:title"]`).Attr("content")
	image, _ := doc.Find(`meta[property="og:image"]`).Attr("content")
	url, urlExists := doc.Find(`meta[property="og:url"]`).Attr("content")

	if !descriptionExists || !titleExists {
		if !urlExists || url == link {
			return nil, nil
		}
		return scraper.FetchData(url)
	}
	return &FetchedData{
		Description: description,
		Title:       title,
		URL:         url,
		Image:       image,
	}, nil
}
