package google

type Item struct {
	Position int
	PositionOverall int
	Description string
	Title string
	URL string
	DisplayURL string
}

type PaidItem struct {
	Position int
	Description string
	Title string
	URL string
}

type PLAItem struct {
	URL string
	Title string
	Source string
	Price string
}

type Pagination struct {
	Current int64
	Next string
	OtherPages []string
}

type Serp struct {
	TotalResultCount int64
	OrganicItems []Item
	PaidItems []PaidItem
	TopPLAItems []PLAItem
	CommercialUnitPlA []PLAItem
	RelatedQuestions []string
	Pagination Pagination
}