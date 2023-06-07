package telegram

import (
	"os"

	"github.com/giicoo/telegramStickBot/internal/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

type APIBot struct {
	bot *tgbotapi.BotAPI
	cfg *config.Config
}

func NewBot(bot *tgbotapi.BotAPI, cfg *config.Config) *APIBot {
	return &APIBot{
		bot: bot,
		cfg: cfg,
	}
}

func (b *APIBot) Start() error {
	logrus.Info("Authorized on account ", b.bot.Self.UserName)

	err := os.MkdirAll(b.cfg.PathIn, os.ModePerm)
	if err != nil {
		return err
	}
	err = os.MkdirAll(b.cfg.PathOut, os.ModePerm)
	if err != nil {
		return err
	}

	updates := b.initUpdateChannel()
	if err = b.checkingUpdates(updates); err != nil {
		return err
	}

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

			// check commands
			if update.Message.IsCommand() {
				if err := b.handlerCommands(update.Message); err != nil {
					return err
				}
				continue
			}

			//check message
			if update.Message.Text != "" {
				if err := b.handlerDefault(update.Message); err != nil {
					return err
				}
				continue
			}

			if err := b.handlerMessages(update.Message); err != nil {
				return nil
			}

		}
	}
	return nil
}
