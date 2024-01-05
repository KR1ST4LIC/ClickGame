package manager

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

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
				fmt.Println(err)
			}
			b = k
		}
		if b {
			m.ParsingUserMessage(ctx, update, bot)
			return
		} else {
			err := m.storage.InsertUser(ctx, update.Message.From.ID, update.Message.From.FirstName)
			//инсертнуть в ван сек
			if err != nil {
				fmt.Println(err)
			}
			err = m.storage.InsertMinUser(ctx, update.Message.From.ID)
			if err != nil {
				fmt.Println(err)
			}
			m.ParsingUserMessage(ctx, update, bot)
			return
		}
	}

	if update.CallbackQuery != nil {
		m.ParsingUserCbd(ctx, update, bot)
		return
	}
}

func (m *Manager) GetUserStatus(ctx context.Context, userID int64) string {
	status, err := m.storage.GetStatus(ctx, userID)
	if err != nil {
		fmt.Println(err)
	}
	return status
}

func (m *Manager) ParsingUserMessage(ctx context.Context, update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	text := update.Message.Text
	userID := update.Message.From.ID

	switch text {
	case "/start":
		msg := user.StartMsg(userID)
		_, err := bot.Send(msg)
		if err != nil {
			fmt.Println(err)
		}
	case "🖱 кликать 🖱":
		msg := user.ClickIKB(userID)
		_, err := bot.Send(msg)
		if err != nil {
			fmt.Println(err)
		}
	case "💶бустер автокликов💶":
		status := m.GetUserStatus(ctx, userID)
		if status == "boost" {
			msg := tgbotapi.NewMessage(userID, "у вас уже включен бустер")
			_, err := bot.Send(msg)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			m.storage.UpdateStatus(ctx, userID, "boost")
			msg := tgbotapi.NewMessage(userID, "✅бустер афк кликов включен!✅")
			_, err := bot.Send(msg)
			if err != nil {
				fmt.Println(err)
			}
		}
		/*else {
			msg := tgbotapi.NewMessage(userID, "Бустер включен")
			// в редиску пихать то, что в след раз он получит x20 монет
		}*/
	}
}

func (m *Manager) ParsingUserCbd(ctx context.Context, update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	cbd := strings.Split(update.CallbackData(), "?")

	switch cbd[0] {
	case "getmoney":
		multi, _ := strconv.Atoi(cbd[1])
		click, err := m.storage.GetMulti(ctx, update.CallbackQuery.From.ID)
		if err != nil {
			fmt.Println(err)
		}
		balance := float32(multi) * click.Multiplier * float32(click.Click)
		err = m.storage.InsertMin(ctx, update.CallbackQuery.From.ID, int64(balance))
		if err != nil {
			fmt.Println(err)
		}
		msg := tgbotapi.NewMessage(update.CallbackQuery.From.ID, "Монеты были начислены")
		_, err = bot.Send(msg)
		if err != nil {
			fmt.Println(err)
		}

	}
}
