package admin

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func LoginAdmin(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.From.ID, "Ты успешно вошел в админ-панель")
	_, err := bot.Send(msg)
	if err != nil {
		log.Fatal(err)
	}
}
