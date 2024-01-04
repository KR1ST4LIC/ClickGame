package user

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func ParsingUserMessage(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	text := update.Message.Text
	userID := update.Message.From.ID
	//msgID := update.Message.MessageID
	//chatID := update.Message.Chat.ID

	switch text {
	case "/start":
		msg := startMsg(userID)
		bot.Send(msg)

	case "🖱 кликать 🖱":
		msg := clickIKB(userID)
		_, err := bot.Send(msg)
		if err != nil {
			fmt.Println(err)
		}
	case "💶бустер автокликов💶":
		/*if статус бустера == активен{
			msg := tgbotapi.NewMessage(userID, "У вас уже включен бустер")
		} else {
			msg := tgbotapi.NewMessage(userID, "Бустер включен")
			// в редиску пихать то, что в след раз он получит x20 монет
		}*/
	}
}

func startMsg(userID int64) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(userID, "Хорошей игры!")
	msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("🖱 кликать 🖱"),
			tgbotapi.NewKeyboardButton("💶бустер автокликов💶"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("🏪 магазин 🏪"),
			tgbotapi.NewKeyboardButton("помощь"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("💰 реклама 💰"),
			tgbotapi.NewKeyboardButton("🏠 кланы 🏠"),
		),
	)
	return msg
}

func clickIKB(userID int64) tgbotapi.MessageConfig {
	markup := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("x1", "getmoney?1"),
			tgbotapi.NewInlineKeyboardButtonData("x1", "getmoney?1"),
			tgbotapi.NewInlineKeyboardButtonData("x1", "getmoney?1"),
			tgbotapi.NewInlineKeyboardButtonData("x1", "getmoney?1"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("x1", "getmoney?1"),
			tgbotapi.NewInlineKeyboardButtonData("x1", "getmoney?1"),
			tgbotapi.NewInlineKeyboardButtonData("x1", "getmoney?1"),
			tgbotapi.NewInlineKeyboardButtonData("x1", "getmoney?1"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("x1", "getmoney?1"),
			tgbotapi.NewInlineKeyboardButtonData("x1", "getmoney?1"),
			tgbotapi.NewInlineKeyboardButtonData("x2", "getmoney?2"),
			tgbotapi.NewInlineKeyboardButtonData("x1", "getmoney?1"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("x1", "getmoney?1"),
			tgbotapi.NewInlineKeyboardButtonData("x1", "getmoney?1"),
			tgbotapi.NewInlineKeyboardButtonData("x1", "getmoney?1"),
			tgbotapi.NewInlineKeyboardButtonData("x1", "getmoney?1"),
		),
	)
	msg := tgbotapi.NewMessage(userID, "кликай ниже")
	msg.ReplyMarkup = markup
	return msg
}
