package bing

import (
	"github.com/PuerkitoBio/goquery"
	"io"
	"regexp"
	"strconv"
	"strings"
)

var (
	specialPages = []string{"wikipedia"}
)

func Parse(r io.Reader) (*Serp, error) {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, err
	}

	res := &Serp{}
	resultStats := doc.Find("span.sb_count")
	matches := regexp.MustCompile(`\d+(,\d+)*`).FindAllString(resultStats.Text(), -1)
	if len(matches) > 0 {
		res.TotalResultCount, err = strconv.ParseInt(strings.ReplaceAll(matches[0], ",", ""), 0, 64)
		if err != nil {
			return nil, err
		}
	}

	doc.Find("#b_results > li.b_ad:first-child div.sb_add").Each(func(i int, el *goquery.Selection) {
		res.AdItems = append(res.AdItems, Item{
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

		res.Items = append(res.Items, Item{
			Position:        i,
			PositionOverall: res.CountItems(),
			Description:     el.Find(".b_caption > p").Text(),
			Title:           el.Find("h2 > a").Text(),
			URL:             el.Find(".b_caption > div cite").Text(),
			DisplayURL:      el.Find(".b_caption > div cite").Text(),
		})
	})

	doc.Find("#b_results > li.b_adBottom div.sb_add").Each(func(i int, el *goquery.Selection) {
		res.AdBottomItems = append(res.AdBottomItems, Item{
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

	pagination := doc.Find("li.b_pag")
	res.Pagination.Current, err = strconv.ParseInt(pagination.Find("a.sb_pagS_bp").Text(), 0, 64)
	if err != nil {
		return nil, err
	}

	res.Pagination.Next = pagination.Find("a.sb_pagN").AttrOr("href", "")
	pagination.Find("a.sb_bp").Each(func(i int, el *goquery.Selection) {
		href := el.AttrOr("href", "")
		if href == "" || href == res.Pagination.Next {
			return
		}

		res.Pagination.OtherPages = append(res.Pagination.OtherPages, href)
	})

	return res, nil
}
