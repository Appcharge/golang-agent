package config

import (
	"time"
)

type PlayerUpdateBalanceRequest struct {
	AppChargePaymentID     string    `json:"appChargePaymentId"`
	PurchaseDateAndTimeUtc time.Time `json:"purchaseDateAndTimeUtc"`
	GameID                 string    `json:"gameId"`
	PlayerID               string    `json:"playerId"`
	BundleName             string    `json:"bundleName"`
	BundleID               string    `json:"bundleId"`
	SKU                    string    `json:"sku"`
	PriceInCents           int       `json:"priceInCents"`
	Currency               string    `json:"currency"`
	PriceInDollar          float64   `json:"priceInDollar"`
	SubTotal               float64   `json:"subTotal"`
	Tax                    float64   `json:"tax"`
	OriginalPriceInDollar  int       `json:"originalPriceInDollar"`
	Action                 string    `json:"action"`
	ActionStatus           string    `json:"actionStatus"`
	Products               []Product `json:"products"`
}

type Product struct {
	Amount int    `json:"amount"`
	Sku    string `json:"sku"`
	Name   string `json:"name"`
}

type PlayerUpdateBalanceResponse struct {
	PublisherPurchaseId string `json:"publisherPurchaseId"`
	PurchaseTime        string `json:"purchaseTime"`
}
