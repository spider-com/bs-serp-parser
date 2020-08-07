package bing

type Pagination struct {
	Current int64
	Next string
	OtherPages []string
}

type Item struct {
	Position int
	PositionOverall int
	Description string
	Title string
	URL string
	DisplayURL string
}

type Serp struct {
	TotalResultCount int64
	Items []Item
	AdItems []Item
	AdBottomItems []Item
	RelatedQuestions []string
	Pagination Pagination
}

func (br Serp) CountItems() int {
	return len(br.Items) + len(br.AdItems) + len(br.Items)
}
