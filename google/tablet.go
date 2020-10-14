package google

import (
	"github.com/PuerkitoBio/goquery"
	"io"
)

func parseTablet(r io.Reader) (*serpDesktop, error) {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, err
	}

	res := &serpDesktop{}
	searchNode := doc.Find("div#rso")
	searchNode.Find("div.mnr-c.xpd").Each(func(pos int, el *goquery.Selection) {
		res.OrganicItems = append(res.OrganicItems, item{
			Position:        pos,
			PositionOverall: pos,
			Title:           el.Find(`div[role="heading"] > div`).Text(),
			DisplayURL:      el.Find("div > span > span:first-child").Text(),
			Description:     el.Find("div.BmP5tf > div > span:last-child").Text(),
			URL:             el.Find("div.mnr-c > div > div > div:first-child > a").AttrOr("href", ""),
		})
	})

	searchNode.Find("div.related-question-pair div.hide-focus-ring:last-child").Each(func(i int, el *goquery.Selection) {
		res.RelatedQuestions = append(res.RelatedQuestions, el.Text())
	})

	return res, nil
}
