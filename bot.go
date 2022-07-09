package main

import (
	"log"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	cmd "cmd"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("5433221809:AAFoubVvGt5d5y9yFRCcBLWSrIs5KvOCwrY")
	if err != nil {
		log.Panic(err)
	}


	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		if !update.Message.IsCommand() { // ignore any non-command Messages
			continue
		}

		// Create a new MessageConfig. We don't have text yet,
		// so we leave it empty.
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		// Extract the command from the Message.
		cmd.Cmd()		

		if _, err := bot.Send(msg); err != nil {			log.Panic(err)
		}
	}
}
