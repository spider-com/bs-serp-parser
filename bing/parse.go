package bing

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	ut "github.com/spider-com/bs-serp-parser"
	"io"
	"regexp"
	"strconv"
	"strings"
)

const (
	domain = "https://www.bing.com"
)

var (
	specialPages = []string{"wikipedia"}
)

func ParseJSON(r io.Reader) (res []byte, err error) {
	v, err := parse(r)
	if err != nil {
		return
	}

	return json.Marshal(v)
}

func parse(r io.Reader) (*serp, error) {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, err
	}

	res := &serp{}
	resultStats := doc.Find("span.sb_count")
	matches := regexp.MustCompile(`\d+(,\d+)*`).FindAllString(resultStats.Text(), -1)
	if len(matches) > 0 {
		res.TotalResultCount, err = strconv.ParseInt(strings.ReplaceAll(matches[0], ",", ""), 0, 64)
		if err != nil {
			return nil, err
		}
	}

	doc.Find("#b_results > li.b_ad:first-child div.sb_add").Each(func(i int, el *goquery.Selection) {
		res.AdItems = append(res.AdItems, item{
			Position:        i,
			PositionOverall: i,
			Description:     el.Find(".b_caption > p").Text(),
			Title:           el.Find("h2 > a").Text(),
			URL:             el.Find(".b_caption > div cite").Text(),
			DisplayURL:      el.Find(".b_caption > div cite").Text(),
		})
	})

	doc.Find("main > ol#b_results li.b_algo").Each(func(i int, el *goquery.Selection) {
		for _, p := range specialPages {
			if strings.Contains(el.Find("h2 > a").AttrOr("href", ""), p) {
				return
			}
		}

		res.Items = append(res.Items, item{
			Position:        i,
			PositionOverall: res.CountItems(),
			Description:     el.Find(".b_caption > p").Text(),
			Title:           el.Find("h2 > a").Text(),
			URL:             el.Find(".b_caption > div cite").Text(),
			DisplayURL:      el.Find(".b_caption > div cite").Text(),
		})
	})

	doc.Find("#b_results > li.b_adBottom div.sb_add").Each(func(i int, el *goquery.Selection) {
		res.AdBottomItems = append(res.AdBottomItems, item{
			Position:        i,
			PositionOverall: res.CountItems(),
			Description:     el.Find(".b_caption > p").Text(),
			Title:           el.Find("h2 > a").Text(),
			URL:             el.Find(".b_caption > div cite").Text(),
			DisplayURL:      el.Find(".b_caption > div cite").Text(),
		})
	})

	doc.Find("div.b_rich .b_vlist2col > ul > li").Each(func(i int, el *goquery.Selection) {
		res.RelatedQuestions = append(res.RelatedQuestions, el.Text())
	})

	pag := doc.Find("li.b_pag")
	res.Pagination.Current, err = strconv.ParseInt(pag.Find("a.sb_pagS_bp").Text(), 0, 64)
	if err != nil {
		return nil, err
	}

	pag.Find("a.sb_bp").Each(func(i int, el *goquery.Selection) {
		href := el.AttrOr("href", "")
		if href == "" {
			return
		}

		href = ut.PrependDomainToHRef(domain, href)
		if el.HasClass("sb_pagN") {
			res.Pagination.Next = href
		} else {
			res.Pagination.OtherPages = append(res.Pagination.OtherPages, href)
		}
	})

	return res, nil
}
