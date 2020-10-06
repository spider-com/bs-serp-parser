package google

type item struct {
	Position        int    `json:"position"`
	PositionOverall int    `json:"positionOverall"`
	Description     string `json:"description"`
	Title           string `json:"title"`
	URL             string `json:"url"`
	DisplayURL      string `json:"displayURL"`
}

type paidItem struct {
	Position    int    `json:"position"`
	Description string `json:"description"`
	Title       string `json:"title"`
	URL         string `json:"url"`
}

type plaItem struct {
	URL    string `json:"url"`
	Title  string `json:"title"`
	Source string `json:"source"`
	Price  string `json:"price"`
}

type pagination struct {
	Current    int64    `json:"current"`
	Next       string   `json:"next"`
	OtherPages []string `json:"otherPages"`
}

type serp struct {
	TotalResultCount  int64      `json:"totalResultCount"`
	OrganicItems      []item     `json:"organicItems"`
	PaidItems         []paidItem `json:"paidItems"`
	TopPLAItems       []plaItem  `json:"topPLAItems"`
	CommercialUnitPLA []plaItem  `json:"commercialUnitPLA"`
	RelatedQuestions  []string   `json:"relatedQuestions"`
	Pagination        pagination `json:"pagination"`
}
