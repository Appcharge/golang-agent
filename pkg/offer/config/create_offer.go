package config

type CreateOfferRequest struct {
	PublisherOfferID string             `json:"publisherOfferId"`
	Name             string             `json:"name"`
	Type             string             `json:"type"`
	Intervals        []Interval         `json:"intervals"`
	OfferUIID        string             `json:"offerUiId"`
	Active           bool               `json:"active"`
	Segments         []string           `json:"segments"`
	ProductsSequence []ProductsSequence `json:"productsSequence"`
	CreatedBy        string             `json:"createdBy"`
}
type ProductsSequence struct {
	Index              int       `json:"index"`
	Products           []Product `json:"products"`
	PriceInUsdCents    int       `json:"priceInUsdCents"`
	PlayerAvailability int       `json:"playerAvailability"`
}

type Product struct {
	PublisherProductId string `json:"publisherProductId"`
	Quantity           int    `json:"quantity"`
}

type Interval struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}

type OfferResponse struct {
	Result  string `json:"result"`
	Status  string `json:"status"`
	Message string `json:"message"`
}
