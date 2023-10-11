package player

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"golang-agent/pkg/player/config"
	"golang-agent/pkg/player/service"
	"net/http"
)

// @Summary Sync Player Info.
// @Description Sync's all info about player from the DB.
// @Param request body config.PlayerInfoSyncRequest true "Player Info Sync Request"
// @Success 200 {object} config.PlayerData
// @Failure 400 {string} string "Bad Request"
// @Router /playerInfoSync [post]
func SyncInfo(c *gin.Context) {
	var input config.PlayerInfoSyncRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
	}

	playerService := service.NewPlayerService()

	playerData, err := playerService.SyncInfo(input.PlayerId)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.JSON(http.StatusOK, playerData)
}

// @Summary Update Player Balance.
// @Description Update balance for given player.
// @Param request body config.PlayerUpdateBalanceRequest true "Player Update Balance"
// @Success 200 {object} config.PlayerUpdateBalanceResponse
// @Failure 400 {object} string "Bad Request"
// @Router /playerUpdateBalance [post]
func UpdateBalance(c *gin.Context) {
	var input config.PlayerUpdateBalanceRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
	}
	playerService := service.NewPlayerService()

	jsonData, err := json.Marshal(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal JSON"})
	}

	playerUpdateBalanceResponse, err := playerService.UpdateBalance(string(jsonData))

	c.JSON(http.StatusOK, playerUpdateBalanceResponse)
}
