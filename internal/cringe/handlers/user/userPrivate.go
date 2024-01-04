package user

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func UserCbd(bot *tgbotapi.BotAPI) {

}

func ParsingUserMessage(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	text := update.Message.Text
	userID := update.Message.From.ID
	//name := update.Message.From.FirstName
	//chatID := update.Message.Chat.ID

	//var balance int
	//_ = db.QueryRow("SELECT balance FROM users WHERE tg_id = ($1);",
	//	userID,
	//).Scan(&balance)

	switch text {
	case "/start":
		//k := internal.CheckUser(database, userID)
		//if !k {
		//	err := internal.InsertUser(database, userID, name)
		//	if err != nil {
		//		fmt.Println(err)
		//	}
		//} else if k {
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
		bot.Send(msg)

		//}
	case "🖱 кликать 🖱":
		msg := tgbotapi.NewMessage(userID, "Кликай ниже")
		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("x1", "user/getmoney?1"),
				tgbotapi.NewInlineKeyboardButtonData("x1", "user/getmoney?1"),
				tgbotapi.NewInlineKeyboardButtonData("x1", "user/getmoney?1"),
				tgbotapi.NewInlineKeyboardButtonData("x1", "user/getmoney?1"),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("x1", "user/getmoney?1"),
				tgbotapi.NewInlineKeyboardButtonData("x1", "user/getmoney?1"),
				tgbotapi.NewInlineKeyboardButtonData("x1", "user/getmoney?1"),
				tgbotapi.NewInlineKeyboardButtonData("x1", "user/getmoney?1"),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("x1", "user/getmoney?1"),
				tgbotapi.NewInlineKeyboardButtonData("x1", "user/getmoney?1"),
				tgbotapi.NewInlineKeyboardButtonData("x2", "user/getmoney?2"),
				tgbotapi.NewInlineKeyboardButtonData("x1", "user/getmoney?1"),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("x1", "user/getmoney?1"),
				tgbotapi.NewInlineKeyboardButtonData("x1", "user/getmoney?1"),
				tgbotapi.NewInlineKeyboardButtonData("x1", "user/getmoney?1"),
				tgbotapi.NewInlineKeyboardButtonData("x1", "user/getmoney?1"),
			),
		)
		//msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
		//	tgbotapi.NewKeyboardButtonRow(
		//		tgbotapi.NewKeyboardButton("Назад в главное меню"),
		//	),
		//)
		bot.Send(msg)
	case "💶бустер автокликов💶":
		/*if статус бустера == активен{
			msg := tgbotapi.NewMessage(userID, "У вас уже включен бустер")
		} else {
			msg := tgbotapi.NewMessage(userID, "Бустер включен")
			// в редиску пихать то, что в след раз он получит x20 монет
		}*/

	}
}
