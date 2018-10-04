package main

import (
	"log"
	"strings"

	"gopkg.in/telegram-bot-api.v4"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(BotToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {

		if update.Message == nil {
			continue
		}

		if err != nil {
			log.Println(err.Error())
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		text := FormatMessage(strings.ToLower(update.Message.Text))
		incomingMessage := strings.Split(text, " ")

		message := getSendMessage(incomingMessage)

		if !empty(message) {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
		}
	}
}
