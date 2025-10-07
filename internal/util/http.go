package util

import (
	"io"
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

func GetHtml(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	err = resp.Body.Close()
	if err != nil {
		return "", err
	}

	return string(body), nil
}
