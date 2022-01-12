package bot

import (
	"fmt"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	info "github.com/nzim-dev/waha-finder-bot"
	scrapper "github.com/nzim-dev/waha-finder-bot/pkg/scrapper"
	"github.com/sirupsen/logrus"
)

func InitBot() error {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
	if err != nil {
		return err
	}

	bot.Debug = true

	logrus.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			logrus.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			//set search credentials
			credentials, err := info.NewRequest(update.Message.Text)
			if err != nil {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "wrong input, try again")
				msg.ReplyToMessageID = update.Message.MessageID
				bot.Send(msg)
				continue
			}

			res := fmt.Sprintf("%v", scrapper.ScrapData(credentials))
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, res)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}

	return nil
}
