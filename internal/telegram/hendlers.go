package telegram

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/giicoo/telegramStickBot/internal/images"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

func (b *APIBot) hendlerCommads(message *tgbotapi.Message) error {
	switch message.Text {
	case "/start":
		{
			return b.hendlerStart(message)
		}
	default:
		{
			return b.hendlerDefault(message)
		}
	}
}
func (b *APIBot) hendlerDefault(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Я не понимаю ваш язык, только байты файлов :)")
	if _, err := b.bot.Send(msg); err != nil {
		return err
	}
	return nil
}
func (b *APIBot) hendlerStart(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Отправьте фото как файл, чтобы я преобразовал его размер для стикеров :)")
	_, err := b.bot.Send(msg)
	if err != nil {
		return err
	}
	return nil
}

func (b *APIBot) hendlerMessages(message *tgbotapi.Message) error {
	chat_id, msg_id := message.Chat.ID, int64(message.MessageID)
	if message.Document != nil {
		id := message.Document.FileID
		url, err := b.bot.GetFileDirectURL(id)
		if err != nil {
			return err
		}
		b.getToSendImage(url, chat_id, msg_id)
		b.deleteSendedFile(chat_id, msg_id)
	} else if message.Text != "" {
		b.hendlerDefault(message)
	} else {
		msg := tgbotapi.NewMessage(chat_id, "Пожалуйста отправьте файлом (без сжатия)")
		if _, err := b.bot.Send(msg); err != nil {
			return err
		}
	}

	return nil
}

func (b *APIBot) deleteSendedFile(chat_id int64, msg_id int64) error {
	path_in := fmt.Sprintf("internal/images/files/in/%v_%v.png", chat_id, msg_id)
	path_out := fmt.Sprintf("internal/images/files/out/%v_%v.png", chat_id, msg_id)

	if err := os.Remove(path_in); err != nil {
		return err
	}

	if err := os.Remove(path_out); err != nil {
		return err
	}
	return nil
}
func (b *APIBot) getToSendImage(url string, chat_id int64, msg_id int64) error {
	logrus.Info(chat_id, " ...Downloading image...")
	err := b.downloadImage(url, chat_id, msg_id)
	if err != nil {
		return err
	}

	logrus.Info(chat_id, " ...Image resize.")
	err = images.ResizeImage(chat_id, msg_id)
	if err != nil {
		return err
	}

	logrus.Info(chat_id, " ...Send file...")
	err = b.sendFile(chat_id, msg_id)
	if err != nil {
		return err
	}

	return nil
}
func (b *APIBot) downloadImage(url string, chat_id int64, msg_id int64) error {
	path := fmt.Sprintf("internal/images/files/in/%v_%v.png", chat_id, msg_id)

	out, err := os.Create(path)

	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func (b *APIBot) sendFile(chat_id int64, msg_id int64) error {
	path := fmt.Sprintf("internal/images/files/out/%v_%v.png", chat_id, msg_id)
	photo := tgbotapi.NewDocument(chat_id, tgbotapi.FilePath(path))
	if _, err := b.bot.Send(photo); err != nil {
		return err
	}
	return nil
}
