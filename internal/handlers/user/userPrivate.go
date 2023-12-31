package user

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jackc/pgx/v4/pgxpool"
)

func UserCbd(bot *tgbotapi.BotAPI) {

}

func ParsingUserMessage(update tgbotapi.Update, bot *tgbotapi.BotAPI, database *pgxpool.Pool) {
	//text := update.Message.Text
	//userID := update.Message.From.ID
	//name := update.Message.From.FirstName
	//chatID := update.Message.Chat.ID

	/*var balance int
	_ = db.QueryRow("SELECT balance FROM users WHERE tg_id = ($1);",
		userID,
	).Scan(&balance)*/
	//switch text {
	//case "/start":
	//	k := internal.CheckUser(database, userID)
	//	if !k {
	//		err := internal.InsertUser(database, userID, name)
	//		if err != nil {
	//			fmt.Println(err)
	//		}
	//	}
	//
	//}
}
