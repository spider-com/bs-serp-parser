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

type serpDesktop struct {
	RelatedQuestions  []string   `json:"relatedQuestions"`
	OrganicItems      []item     `json:"organicItems"`
	TotalResultCount  int64      `json:"totalResultCount"`
	PaidItems         []paidItem `json:"paidItems"`
	TopPLAItems       []plaItem  `json:"topPLAItems"`
	CommercialUnitPLA []plaItem  `json:"commercialUnitPLA"`
	Pagination        pagination `json:"pagination"`
}

type mobileItem struct {
	item
	IsAMP bool
}

type serpMobile struct {
	RelatedQuestions []string     `json:"relatedQuestions"`
	OrganicItems     []mobileItem `json:"organicItems"`
}
