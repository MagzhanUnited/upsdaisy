package main

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	// Replace "YOUR_BOT_API_TOKEN" with the token you received from the BotFather.
	bot, err := tgbotapi.NewBotAPI("6517147428:AAG_OkAcpjQOolOjzXko3rjvE4lcwtFFWhU")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	// Create an update configuration with an offset.
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60
	// Continuously fetch updates using long polling.
	updates, err := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		fmt.Println("update.Message.Text:", update.Message.Text)
		// Reply to the user's message.
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "SUIIIII")
		_, err := bot.Send(msg)
		if err != nil {
			log.Panic(err)
		}
	}
}
