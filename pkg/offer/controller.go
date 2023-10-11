package offer

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"golang-agent/pkg/offer/config"
	"golang-agent/pkg/offer/service"
	"net/http"
)

// @Summary Create offer.
// @Description Create's offer in Appcharge systems.
// @Param request body config.CreateOfferRequest true "Player Info Sync Request"
// @Success 200 {object} config.OfferResponse
// @Failure 400 {string} string "Bad Request"
// @Router /offer [post]
func CreateOffer(c *gin.Context) {
	var input config.CreateOfferRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
	}

	jsonData, err := json.Marshal(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal JSON"})
	}

	offerService := service.NewOfferService()

	createOfferData, err := offerService.CreateOffer(string(jsonData))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
	}

	c.JSON(http.StatusOK, createOfferData)
}

// @Summary Update Player Balance.
// @Description Update balance for given player.
// @Param request body config.PlayerUpdateBalanceRequest true "Player Update Balance"
// @Success 200 {object} config.PublisherUpdateBalanceResponse
// @Failure 400 {object} string "Bad Request"
// @Router /offer [put]
//func UpdateOffer(c *gin.Context) {
//	var input config.PlayerUpdateBalanceRequest
//
//	if err := c.ShouldBindJSON(&input); err != nil {
//		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
//	}
//	playerService := service.NewPlayerService()
//
//	jsonData, err := json.Marshal(input)
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal JSON"})
//	}
//
//	playerUpdateBalanceResponse, err := playerService.UpdateBalance(string(jsonData))
//
//	c.JSON(http.StatusOK, playerUpdateBalanceResponse)
//}
