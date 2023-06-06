package main

import (
	"os"

	"github.com/giicoo/telegramStickBot/internal/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := godotenv.Load(); err != nil {
		logrus.Fatal("No .env file found")
	}
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TG_TOKEN"))
	if err != nil {
		logrus.Fatal("Error with Token ", err)
	}

	APIBot := telegram.NewBot(bot)
	if err := APIBot.Start(); err != nil {
		logrus.Fatal("Start don't work ", err)
	}
}
