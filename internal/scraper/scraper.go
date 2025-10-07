package scraper

import (
	"regexp"
)

type FetchedData struct {
	Title       string
	Description string
	URL         string
	Image       string
}

type Scraper interface {
	FetchData(link string) (*FetchedData, error)
	Match(content string) string
	EmbedColor() int
	SourceName() string
}

type BaseScraper struct {
	sourceName string
	regex      *regexp.Regexp
	embedColor int
}

func (scraper *BaseScraper) Match(content string) string {
	match := scraper.regex.FindString(content)
	return match
}

func (scraper *BaseScraper) EmbedColor() int {
	return scraper.embedColor
}

func (scraper *BaseScraper) SourceName() string {
	return scraper.sourceName
}
