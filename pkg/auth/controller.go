package auth

import (
	"github.com/gin-gonic/gin"
	"golang-agent/pkg/auth/config"
	"golang-agent/pkg/auth/service"
	"golang-agent/pkg/auth/service/auth_method"
	"net/http"
)

// @Summary Authenticate user.
// @Description Authenticates user using Facebook, Google, AppleId, Token or Password.
// @Param request body config.AuthenticationRequest true "Authentication Request"
// @Success 200 {string} string "OK"
// @Failure 400 {object} string "Bad Request"
// @Router /auth [post]
func Auth(c *gin.Context) {
	var input config.AuthenticationRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
	}

	authService := &service.AuthenticationServiceImpl{
		FacebookAuth: &auth_method.FacebookAuthInit{},
		GoogleAuth:   &auth_method.GoogleAuthInit{},
		AppleAuth:    &auth_method.AppleAuthInit{},
	}

	authResponse, err := authService.AuthenticatePlayer(&input)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, authResponse)
}
