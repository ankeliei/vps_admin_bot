package main

import (
	"log"
	"vps_admin_bot/config"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	serverCFG, err := config.LoadJSONConfig("/config/config.json")
	if err != nil {
		panic(err)
	}

	router.GET("/", func(c *gin.Context) {
		c.String(200, "Welcome to HTTPS!")
	})
	log.Print(serverCFG)
	// 启动HTTPS服务
	router.RunTLS(serverCFG.Server.Addr, serverCFG.Server.CertFile, serverCFG.Server.KeyFile)
}
