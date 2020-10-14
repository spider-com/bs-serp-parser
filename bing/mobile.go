package bing

import (
	"github.com/PuerkitoBio/goquery"
	"io"
	"strings"
)

func parseMobile(r io.Reader) (*serp, error) {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, err
	}

	res := &serp{}
	doc.Find("ol#b_results li.b_algo").Each(func(i int, el *goquery.Selection) {
		for _, p := range specialPages {
			if strings.Contains(el.Find("h2 > a").AttrOr("href", ""), p) {
				return
			}
		}

		res.Items = append(res.Items, item{
			Position:        i,
			PositionOverall: res.CountItems(),
			Title:           el.Find("div.b_algoheader > a > h2").Text(),
			URL:             el.Find(".b_algoheader > a").AttrOr("href", ""),
			DisplayURL:      el.Find(".b_attribution > cite").Text(),
		})
	})

	return res, nil
}