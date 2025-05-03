package telbot

import (
	"log"

	"vps_admin_bot/config"

	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	Bot *tgbotapi.BotAPI
)

func init() {
	var err error
	Bot, err = tgbotapi.NewBotAPI(config.ServerCFG.Telebot.BotApi)
	if err != nil {
		log.Panic(err)
	}

	_, err = tgbotapi.NewWebhook(config.ServerCFG.Server.Addr + "/webhook")
	if err != nil {
		log.Fatal(err)
	}
}
func handleMessage(msg *tgbotapi.Message) {
	chatID := msg.Chat.ID

	// 处理文本消息
	if msg.Text != "" {
		reply := tgbotapi.NewMessage(chatID, "收到文本消息: "+msg.Text)
		Bot.Send(reply)
		return
	}
}

func WebhookHandler(c *gin.Context) {
	var update tgbotapi.Update
	if err := c.BindJSON(&update); err != nil {
		log.Printf("解析更新失败: %v", err)
		return
	}

	if update.Message != nil {
		handleMessage(update.Message)
	}
}
