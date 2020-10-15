package bing

import (
	"github.com/PuerkitoBio/goquery"
	"io"
)

func ParseAutocomplete(r io.Reader) (res []string, err error) {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, err
	}

	doc.Find("span.sa_tm_text").Each(func(i int, s *goquery.Selection) {
		res = append(res, s.Text())
	})

	return
}
