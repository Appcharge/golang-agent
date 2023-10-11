package config

type OfferSnapshotProduct struct {
	Name   string `json:"name"`
	SKU    string `json:"sku"`
	Amount int    `json:"amount"`
}

type OfferSnapshot struct {
	PublisherID  string                 `json:"publisherId"`
	OfferID      string                 `json:"offerId"`
	Name         string                 `json:"name"`
	SKU          string                 `json:"sku"`
	Products     []OfferSnapshotProduct `json:"products"`
	Price        float32                `json:"price"`
	CurrencyCode string                 `json:"currencyCode"`
	Description  string                 `json:"description"`
}

type GetOrdersResponse struct {
	TotalCount int                 `json:"totalCount"`
	Results    []OrderItemResponse `json:"results"`
}

type OrderItemResponse struct {
	ID                    string         `json:"_id"`
	PublisherID           string         `json:"publisherId"`
	PaymentID             string         `json:"paymentId"`
	PaymentURL            string         `json:"paymentUrl"`
	PublisherPurchaseID   string         `json:"publisherPurchaseId"`
	Currency              string         `json:"currency"`
	AmountTotal           float32        `json:"amountTotal"`
	CurrencySymbol        string         `json:"currencySymbol"`
	OfferSetID            string         `json:"offersetId"`
	BundleName            string         `json:"bundleName"`
	Provider              string         `json:"provider"`
	UTMSource             string         `json:"utmSource"`
	UTMMedium             string         `json:"utmMedium"`
	UTMCampaign           string         `json:"utmCampaign"`
	ModifiedAt            string         `json:"modifiedAt"`
	CreatedAt             string         `json:"createdAt"`
	PlayerID              string         `json:"playerId"`
	ClientGAID            string         `json:"clientGaId"`
	State                 string         `json:"state"`
	Reason                string         `json:"reason"`
	PublisherErrorMessage string         `json:"publisherErrorMessage"`
	Retry                 int            `json:"retry"`
	OfferSnapshot         *OfferSnapshot `json:"offerSnapshot,omitempty"`
}
