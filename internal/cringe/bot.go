package cringe

import (
	"context"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Start(ctx context.Context, bot *tgbotapi.BotAPI) {
	database, err := UploadDataBase(ctx)
	if err != nil {
		log.Fatal(err)
	}
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	fmt.Println(database.Ping(ctx))

	for update := range updates {
		localUpdate := update

		go func(update tgbotapi.Update) {
			if update.Message != nil {

				ParsingNewUpdateMsg(update, bot)
				return
			}

			if update.CallbackQuery != nil {
				//ParsingNewUpdateCbd(bot, update, database)
				return
			}
		}(localUpdate)
	}
}
