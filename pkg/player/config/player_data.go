package config

type PlayerDataFile struct {
	Player PlayerData
}
type PlayerData struct {
	Status             string `json:"status"`
	PlayerProfileImage string `json:"playerProfileImage"`
	PublisherPlayerID  string `json:"publisherPlayerId"`
	PlayerName         string `json:"playerName"`
	Offers             []struct {
		PublisherOfferID string `json:"publisherOfferId"`
		OfferPlayerAvail int    `json:"offerPlayerAvailability"`
		ProductsSequence []struct {
			Index    int `json:"index"`
			Products []struct {
				PublisherProductID string `json:"publisherProductId"`
				Quantity           int    `json:"quantity"`
				BadgeText          string `json:"badgeText,omitempty"`
				RarityProductInfo  struct {
					Stars       int    `json:"stars"`
					TooltipText string `json:"tooltipText"`
					RarityText  string `json:"rarityText"`
				} `json:"rarityProductInfo,omitempty"`
			} `json:"products"`
		} `json:"productsSequence"`
	} `json:"offers"`
	Balances []struct {
		PublisherProductID string `json:"publisherProductId"`
		Balance            int    `json:"balance"`
	} `json:"balances"`
	Segments []interface{} `json:"segments"`
	Focus    struct {
		PublisherBundleID *string `json:"publisherBundleId"`
	} `json:"focus"`
}
