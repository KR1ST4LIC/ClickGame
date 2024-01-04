package admin

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func AdminCbd(bot *tgbotapi.BotAPI) {

}

func ParsingAdminMessage(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	if update.Message.Text == "/admin" {
		LoginAdmin(bot, update)
		return
	}
}
