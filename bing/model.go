package bing

type pagination struct {
	Current    int64    `json:"current"`
	Next       string   `json:"next"`
	OtherPages []string `json:"otherPages"`
}

type item struct {
	Position        int    `json:"position"`
	PositionOverall int    `json:"positionOverall"`
	Description     string `json:"description"`
	Title           string `json:"title"`
	URL             string `json:"url"`
	DisplayURL      string `json:"displayURL"`
}

type serp struct {
	TotalResultCount int64      `json:"totalResultCount"`
	Items            []item     `json:"item"`
	AdItems          []item     `json:"adItems"`
	AdBottomItems    []item     `json:"adBottomItems"`
	RelatedQuestions []string   `json:"relatedQuestions"`
	Pagination       pagination `json:"pagination"`
}

func (br serp) CountItems() int {
	return len(br.Items) + len(br.AdItems) + len(br.Items)
}
