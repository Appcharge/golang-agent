package order

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"golang-agent/pkg/order/config"
	"golang-agent/pkg/order/service"
	"net/http"
)

// @Summary Get orders.
// @Description Get's orders from Appcharge internal API.
// @Param request body config.GetOrdersRequest true "Orders Request"
// @Success 200 {object} config.GetOrdersResponse
// @Failure 400 {object} string "Bad Request"
// @Router /getOrders [post]
func GetOrders(c *gin.Context) {
	var input config.GetOrdersRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orderService := service.NewOrderService()

	jsonData, err := json.Marshal(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal JSON"})
	}

	getOrdersResponse, err := orderService.GetOrders(string(jsonData))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get Orders"})
	}

	c.JSON(http.StatusOK, getOrdersResponse)
}
