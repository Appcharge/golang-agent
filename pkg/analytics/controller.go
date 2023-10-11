package analytics

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"golang-agent/pkg/analytics/config"
	"golang-agent/pkg/analytics/service"
	"net/http"
)

// @Summary Get Analytics Data.
// @Description Retrieve analytics data from Appcharge API's.
// @Param request body config.GetAnalyticsRequest true "Player Info Sync Request"
// @Success 200 {object} config.GetAnalyticsResponse
// @Failure 400 {string} string "Bad Request"
// @Router /analytics [post]
func GetAnalytics(c *gin.Context) {
	var input config.GetAnalyticsRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
	}

	jsonData, err := json.Marshal(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal JSON"})
	}

	analyticsService := service.NewAnalyticsService()

	analyticsData, err := analyticsService.GetAnalytics(string(jsonData))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, analyticsData)
}
