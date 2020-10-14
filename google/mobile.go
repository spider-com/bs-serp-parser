package google

import (
	"github.com/PuerkitoBio/goquery"
	"io"
)

func parseMobile(r io.Reader) (*serpMobile, error) {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, err
	}

	res := &serpMobile{}
	searchNode := doc.Find("div#rso")
	searchNode.Find("div.U3THc").Each(func(pos int, el *goquery.Selection) {
		ampSpan := el.Find(`div.mnr-c > div > div > div:first-child > a > div > span:last-child`)
		res.OrganicItems = append(res.OrganicItems, mobileItem{
			item: item{
				Position:        pos,
				PositionOverall: pos,
				Title:           el.Find(`div[role="heading"] > div`).Text(),
				DisplayURL:      el.Find("div > span > span:first-child").Text(),
				Description:     el.Find("div.BmP5tf > div > span:last-child").Text(),
				URL: el.Find("div.mnr-c > div > div > div:first-child > a").AttrOr("href", ""),
			},
			IsAMP: ampSpan.AttrOr("aria-label", "") == "AMP logo",
		})
	})

	searchNode.Find("div.related-question-pair > div:first-child").Each(func(i int, el *goquery.Selection) {
		if attr, ok := el.Attr("data-q"); ok {
			res.RelatedQuestions = append(res.RelatedQuestions, attr)
		}
	})

	return res, nil
}
