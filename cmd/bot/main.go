package main

import (
	"os"

	"github.com/giicoo/telegramStickBot/internal/config"
	"github.com/giicoo/telegramStickBot/internal/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {

	// get env variables
	if err := godotenv.Load(); err != nil {
		logrus.Fatal("No .env file found")
	}

	// get configs
	cfg, err := config.InitConfig()
	if err != nil {
		logrus.Fatal("No config file found")
	}

	// create Bot
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TG_TOKEN"))
	if err != nil {
		logrus.Fatal("Error with Token ", err)
	}

	// init Bot
	APIBot := telegram.NewBot(bot, cfg)
	if err := APIBot.Start(); err != nil {
		logrus.Fatal("Start don't work ", err)
	}
}
