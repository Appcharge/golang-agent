package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"golang-agent/pkg/signature"
	"io"
	"net/http"
	"strings"
)

func AuthenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.RequestURI == "/mocker/health" || strings.HasPrefix(c.Request.RequestURI, "/mocker/swagger") {
			c.Next()
		} else {
			signatureService := signature.NewSignatureHashingService()

			body, err := io.ReadAll(c.Request.Body)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request body"})
			}
			c.Request.Body = io.NopCloser(bytes.NewBuffer(body))

			signatureFromHeader := c.GetHeader("signature")

			_, err = signatureService.SignPayload(signatureFromHeader, string(body))
			if err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			} else {
				c.Next()
			}
		}
	}
}
