package telegram

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/giicoo/telegramStickBot/internal/imageWork"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

func (b *APIBot) handlerDefault(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, b.cfg.Messages.Default)
	if _, err := b.bot.Send(msg); err != nil {
		return err
	}
	return nil
}

func (b *APIBot) handlerMessages(message *tgbotapi.Message) error {

	chat_id, msg_id := message.Chat.ID, int64(message.MessageID)

	// if send document
	if message.Document != nil {

		id := message.Document.FileID

		url, err := b.bot.GetFileDirectURL(id)
		if err != nil {
			return err
		}

		if err := b.getToSendImage(url, chat_id, msg_id); err != nil {
			return err
		}

	} else {

		msg := tgbotapi.NewMessage(chat_id, b.cfg.Messages.SendNotFile)

		if _, err := b.bot.Send(msg); err != nil {
			return err
		}
	}

	return nil
}

func (b *APIBot) getToSendImage(url string, chat_id int64, msg_id int64) error {
	ext := strings.Split(url, ".")[len(strings.Split(url, "."))-1]
	path_in := fmt.Sprintf("%v/%v_%v.%v", b.cfg.PathIn, chat_id, msg_id, ext)
	// path_in_jpeg := fmt.Sprintf("%v/%v_%v.jpeg", b.cfg.PathIn, chat_id, msg_id)
	path_out := fmt.Sprintf("%v/%v_%v.png", b.cfg.PathOut, chat_id, msg_id)

	logrus.Info(chat_id, " ...Downloading image...")
	err := b.downloadImage(url, path_in)
	if err != nil {
		return err
	}

	logrus.Info(chat_id, " ...Image resize.")
	err = imageWork.ResizeImage(path_in, path_out)
	if err != nil {
		return err
	}

	logrus.Info(chat_id, " ...Send file...")
	err = b.sendFile(chat_id, path_out)
	if err != nil {
		return err
	}

	logrus.Info(chat_id, " ...Delete file...")
	if err := b.deleteSendedFile(path_in, path_out); err != nil {
		return err
	}

	return nil
}
func (b *APIBot) downloadImage(url, path_in string) error {
	in, err := os.Create(path_in)
	if err != nil {
		return err
	}
	defer in.Close()

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

	_, err = io.Copy(in, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func (b *APIBot) sendFile(chat_id int64, path_out string) error {
	photo := tgbotapi.NewDocument(chat_id, tgbotapi.FilePath(path_out))
	if _, err := b.bot.Send(photo); err != nil {
		return err
	}
	return nil
}

func (b *APIBot) deleteSendedFile(path_in, path_out string) error {

	if err := os.Remove(path_in); err != nil {
		return err
	}

	if err := os.Remove(path_out); err != nil {
		return err
	}
	return nil
}
