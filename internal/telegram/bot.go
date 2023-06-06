package telegram

import (
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

type APIBot struct {
	bot *tgbotapi.BotAPI
}

func NewBot(bot *tgbotapi.BotAPI) *APIBot {
	return &APIBot{
		bot: bot,
	}
}

func (b *APIBot) Start() error {
	logrus.Info("Authorized on account ", b.bot.Self.UserName)

	os.MkdirAll("internal/images/files/in", os.ModePerm)
	os.MkdirAll("internal/images/files/out", os.ModePerm)

	updates := b.initUpdateChannel()
	b.checkingUpdates(updates)

	return nil
}

func (b *APIBot) initUpdateChannel() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return b.bot.GetUpdatesChan(u)
}

func (b *APIBot) checkingUpdates(updates tgbotapi.UpdatesChannel) error {
	for update := range updates {
		if update.Message != nil { // If we got a message

			if update.Message.IsCommand() { // check commands
				b.hendlerCommads(update.Message)
				continue
			}

			b.hendlerMessages(update.Message) //check message

		}
	}
	return nil
}
