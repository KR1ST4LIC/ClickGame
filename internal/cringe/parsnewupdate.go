package cringe

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"click-game/internal/cringe/handlers/user"
)

func ParsingNewUpdateMsg(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	user.ParsingUserMessage(update, bot)
}

//
//func ParsingNewUpdateCbd(bot *tgbotapi.BotAPI, update tgbotapi.Update, database *sql.DB) {
//	cbd := strings.Split(update.CallbackData(), "/")
//	if cbd[0] == "admin" {
//		admin.AdminCbd(bot)
//		return
//	}
//	if cbd[0] == "user" {
//		user.UserCbd(bot)
//		return
//	}
//}
