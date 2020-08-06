package google

import (
	"github.com/PuerkitoBio/goquery"
	"io"
	"regexp"
	"strconv"
	"strings"
)

const (
	organicClasses = "div.rc"
	organicTitleClasses = "h3.LC20lb"
	organicURLClasses = "div.rc > div > a"
	organicDisplayURLClasses = "div.TbwUpd > cite"
	organicDescriptionClasses = "span.st"
)

func ParseGoogleResult(r io.Reader) (res *Serp, err error) {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, err
	}

	res = &Serp{}

	resultStats := doc.Find("div#result-stats")
	matches := regexp.MustCompile(`\d+(,\d+)*`).FindAllString(resultStats.Text(), -1)
	if len(matches) > 0 {
		res.TotalResultCount, err = strconv.ParseInt(strings.ReplaceAll(matches[0], ",", ""), 0, 64)
		if err != nil {
			return nil, err
		}
	}

	searchNode := doc.Find("div > div#search")
	searchNode.Find(organicClasses).Each(func(pos int, el *goquery.Selection) {
		res.OrganicItems = append(res.OrganicItems, Item{
			Position:        pos,
			PositionOverall: pos,
			Title:           el.Find(organicTitleClasses).Text(),
			DisplayURL:      el.Find(organicDisplayURLClasses).Text(),
			Description:     el.Find(organicDescriptionClasses).Text(),
			URL: el.Find(organicURLClasses).AttrOr("href", ""),
		})
	})

	searchNode.Find("div.cbphWd").Each(func(i int, el *goquery.Selection) {
		res.RelatedQuestions = append(res.RelatedQuestions, el.Text())
	})

	doc.Find("div.top-pla-group-inner").
		Find("div.pla-unit-container").
		Each(func(i int, el *goquery.Selection) {
			res.TopPLAItems = append(res.TopPLAItems, PLAItem{
				URL:    el.Find("div.pla-unit-title > a:last-child").AttrOr("href", ""),
				Title:  el.Find("a.pla-unit-title-link").Text(),
				Source: el.Find("div.LbUacb > span.VZqTOd").Text(),
				Price:  el.Find("div.e10twf").Text(),
			})
		})

	doc.Find("div.commercial-unit-desktop-rhs").
		Find("div.pla-unit-container").
		Each(func(i int, el *goquery.Selection) {
			res.CommercialUnitPlA = append(res.CommercialUnitPlA, PLAItem{
				URL:    el.Find("div.pla-unit-title > a:last-child").AttrOr("href", ""),
				Title:  el.Find("a.pla-unit-title-link").Text(),
				Source: el.Find("div.LbUacb > span.VZqTOd").Text(),
				Price:  el.Find("div.e10twf").Text(),
			})
		})

	doc.Find("#tads ol > li").Each(func(i int, el *goquery.Selection) {
		// selectors depend on where tads was taken from, this one is a top under PLA items
		res.PaidItems = append(res.PaidItems, PaidItem{
			Position:    i,
			Title:       el.Find("div[role='heading']").Text(),
			URL:         el.Find("a[data-ved]").AttrOr("href", ""),
			Description: el.Find("div.lyLwlc > span").Text(),
		})
	})

	doc.Find("tr[valign='top'] > td").Each(func(i int, el *goquery.Selection) {
		if el.HasClass("YyVfkd") {
			res.Pagination.Current, err = strconv.ParseInt(el.Text(), 0, 64)
		} else if el.HasClass("d6cvqb") {
			res.Pagination.Next = el.Find("a").AttrOr("href", "")
		} else {
			href := el.Find("a").AttrOr("href", "")
			res.Pagination.OtherPages = append(res.Pagination.OtherPages, href)
		}
	})

	return res, nil
}



