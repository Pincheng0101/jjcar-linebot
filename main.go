package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pincheng0101/go-linebot-server-template/config"
	"github.com/pincheng0101/go-linebot-server-template/controller"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	linebot_controller, err := controller.NewLineBotController(
		cfg.ChannelSecret,
		cfg.ChannelToken,
	)
	if err != nil {
		panic(err)
	}

	router := gin.Default()
	router.StaticFile("/qrcode", "./qrcode.html")
	router.POST("/callback", linebot_controller.Callback)
	router.POST("/points/reward", linebot_controller.RewardPoints)
	router.GET("/quota", linebot_controller.Quota)

	if err := router.Run(":8080"); err != nil {
		fmt.Println(err)
	}
}
