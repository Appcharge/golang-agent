package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"golang-agent/common/middleware"
	_ "golang-agent/common/swagger"
	"golang-agent/pkg/analytics"
	"golang-agent/pkg/auth"
	"golang-agent/pkg/health"
	"golang-agent/pkg/offer"
	"golang-agent/pkg/order"
	"golang-agent/pkg/player"
	"os"
)

// @title 	Golang Template API
// @description A Mocker service API in Go using Gin framework.
// @host 	0.0.0.0:8000
// @BasePath /mocker
func main() {
	router := gin.Default()
	router.Use(middleware.AuthenticationMiddleware())
	mockerGroup := router.Group("/mocker")
	// GET controllers
	mockerGroup.GET("/health", health.CheckHealth)
	mockerGroup.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// POST controllers
	mockerGroup.POST("/auth", auth.Auth)
	mockerGroup.POST("/playerInfoSync", player.SyncInfo)
	mockerGroup.POST("/playerUpdateBalance", player.UpdateBalance)
	mockerGroup.POST("/orders", order.GetOrders)
	mockerGroup.POST("/analytics", analytics.GetAnalytics)
	mockerGroup.POST("/offer", offer.CreateOffer)

	//get APP_PORT from the env
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	addr := fmt.Sprintf("%s:%s", "0.0.0.0", os.Getenv("APP_PORT"))
	router.Run(addr)
}
