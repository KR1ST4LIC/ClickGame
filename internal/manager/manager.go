package manager

import (
	"context"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"

	"click-game/internal/handlers/user"
	"click-game/internal/storage"
)

type Manager struct {
	storage storage.Storage
}

func NewManager(s storage.Storage) *Manager {
	return &Manager{
		storage: s,
	}
}

func (m *Manager) Run(ctx context.Context, token string) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		localUpdate := update
		m.handleBotUpdates(ctx, localUpdate, bot)
	}
}

//for {
//m.storage.GetUser()
//}

func (m *Manager) handleBotUpdates(ctx context.Context, update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	if update.Message != nil {
		var b bool
		b = true
		if update.Message.Text == "/start" {
			k, err := m.storage.ExistUser(ctx, update.Message.From.ID)
			if err != nil {
				fmt.Println(errors.Wrap(err, "failed exist user"))
			}
			b = k
		}
		if b {
			user.ParsingUserMessage(update, bot)
			return
		} else {
			err := m.storage.InsertUser(ctx, update.Message.From.ID, update.Message.From.FirstName)
			if err != nil {
				fmt.Println(err)
			}
			user.ParsingUserMessage(update, bot)
			return
		}
	}

	if update.CallbackQuery != nil {
		//user.ParsingNewUpdateCbd(bot, update)
		return
	}
}
