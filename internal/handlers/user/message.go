package user

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func StartMsg(userID int64) tgbotapi.MessageConfig {
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

func ClickIKB(userID int64) tgbotapi.MessageConfig {
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
