package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (b *APIBot) handlerCommands(message *tgbotapi.Message) error {
	switch message.Text {
	case "/start":
		{
			return b.handlerStart(message)
		}
	default:
		{
			return b.handlerDefaultCommand(message)
		}
	}
}

func (b *APIBot) handlerStart(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, b.cfg.Messages.Start)
	_, err := b.bot.Send(msg)
	if err != nil {
		return err
	}
	return nil
}

func (b *APIBot) handlerDefaultCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, b.cfg.Messages.DefaultCommand)
	if _, err := b.bot.Send(msg); err != nil {
		return err
	}
	return nil
}
