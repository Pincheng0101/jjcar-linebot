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
	router.POST("/callback", linebot_controller.Callback)
	router.GET("/quota", linebot_controller.Quota)
	router.StaticFile("/qrcode", "./qrcode.html")

	if err := router.Run(":8080"); err != nil {
		fmt.Println(err)
	}
}
