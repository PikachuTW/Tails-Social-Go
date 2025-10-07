package util

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func GetDoc(url string) (*goquery.Document, error) {
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}
	return doc, nil
}
