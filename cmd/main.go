package main

import (
	"github.com/joho/godotenv"
	"github.com/nzim-dev/waha-finder-bot/pkg/bot"
	scrapper "github.com/nzim-dev/waha-finder-bot/pkg/scrapper"
	"github.com/sirupsen/logrus"
)

func init() {
	if err := scrapper.InitConfig(); err != nil {
		logrus.Fatalf("err init config: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("err init godotenv: %s", err.Error())
	}
}

func main() {
	if err := bot.InitBot(); err != nil {
		logrus.Fatalf("err init bot: %s", err.Error())
	}
}
