package health

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary Get API health status.
// @Description Shows status of an API
// @Success 200 {string} string  "ok"
// @Router /health [get]
func CheckHealth(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "OK")
}
