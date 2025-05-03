package main

import (
	"vps_admin_bot/config"
	"vps_admin_bot/telbot"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(200, "Welcome to HTTPS!")
	})

	router.POST("/webhook", telbot.WebhookHandler)

	// 启动HTTPS服务
	router.RunTLS(config.ServerCFG.Server.Addr, config.ServerCFG.Server.CertFile, config.ServerCFG.Server.KeyFile)
}
