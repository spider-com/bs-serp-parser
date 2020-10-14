package google

import (
	"github.com/PuerkitoBio/goquery"
	ut "github.com/spider-com/bs-serp-parser"
	"io"
	"regexp"
	"strconv"
	"strings"
)

const (
	domain = "https://www.google.com"
)

func parseDesktop(r io.Reader) (*serpDesktop, error) {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, err
	}

	res := &serpDesktop{}
	resultStats := doc.Find("div#result-stats")
	matches := regexp.MustCompile(`\d+(,\d+)*`).FindAllString(resultStats.Text(), -1)
	if len(matches) > 0 {
		res.TotalResultCount, err = strconv.ParseInt(strings.ReplaceAll(matches[0], ",", ""), 0, 64)
		if err != nil {
			return nil, err
		}
	}

	searchNode := doc.Find("div > div#search")
	searchNode.Find("div.rc").Each(func(pos int, el *goquery.Selection) {
		res.OrganicItems = append(res.OrganicItems, item{
				Position:        pos,
				PositionOverall: pos,
				Title:           el.Find("div.rc > div > * > h3 > span").Text(),
				DisplayURL:      el.Find("div.rc cite").Text(),
				Description:     el.Find("div.rc > div span > span").Text(),
				URL: el.Find("div.rc > div > a").AttrOr("href", ""),
		})
	})
	searchNode.Find("div.related-question-pair div.hide-focus-ring:last-child").Each(func(i int, el *goquery.Selection) {
		res.RelatedQuestions = append(res.RelatedQuestions, el.Text())
	})

	doc.Find("div.top-pla-group-inner").
		Find("div.pla-unit-container").
		Each(func(i int, el *goquery.Selection) {
			res.TopPLAItems = append(res.TopPLAItems, plaItem{
				URL:    el.Find("div.pla-unit-title > a:last-child").AttrOr("href", ""),
				Title:  el.Find("a.pla-unit-title-link").Text(),
				Source: el.Find("div.LbUacb > span.VZqTOd").Text(),
				Price:  el.Find("div.e10twf").Text(),
			})
		})

	doc.Find("div.commercial-unit-desktop-rhs").
		Find("div.pla-unit-container").
		Each(func(i int, el *goquery.Selection) {
			res.CommercialUnitPLA = append(res.CommercialUnitPLA, plaItem{
				URL:    el.Find("div.pla-unit-title > a:last-child").AttrOr("href", ""),
				Title:  el.Find("a.pla-unit-title-link").Text(),
				Source: el.Find("div.LbUacb > span.VZqTOd").Text(),
				Price:  el.Find("div.e10twf").Text(),
			})
		})

	doc.Find("#tads ol > li").Each(func(i int, el *goquery.Selection) {
		// selectors depend on where tads was taken from, this one is a top under PLA items
		res.PaidItems = append(res.PaidItems, paidItem{
			Position:    i,
			Title:       el.Find("div[role='heading']").Text(),
			URL:         el.Find("a[data-ved]").AttrOr("href", ""),
			Description: el.Find("div.lyLwlc > span").Text(),
		})
	})

	doc.Find("tr[valign='top'] > td").Each(func(i int, el *goquery.Selection) {
		if el.Text() == "" {
			return
		}

		if el.Is("td[role='heading']") {
			res.Pagination.Next = ut.PrependDomainToHRef(domain, el.Find("a").AttrOr("href", ""))
		} else if el.AttrOr("class", "") != "" {
			res.Pagination.Current, err = strconv.ParseInt(el.Text(), 0, 64)
		} else {
			href := el.Find("a").AttrOr("href", "")
			res.Pagination.OtherPages = append(res.Pagination.OtherPages, ut.PrependDomainToHRef(domain, href))
		}
	})

	return res, nil
}



