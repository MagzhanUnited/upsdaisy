package main

import (
	"fmt"
	"log"
	handler "mytelegrambot/Handler"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	// Replace "YOUR_BOT_API_TOKEN" with the token you received from the BotFather.
	bot, err := tgbotapi.NewBotAPI("6614980196:AAGjdnaxpNgybhFJQA_blvPL3OQjHEMuYd8")
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

		handleCommands(bot, update.Message)

		fmt.Println("update.Message.Text:", update.Message.Text)
		if isYouTubeURL(update.Message.Text) {
			audioBytes, audioextraction := handler.VideoData{Videourl: update.Message.Text}.ServeAudio()
			audioFile := tgbotapi.FileBytes{Name: audioextraction.AudioName, Bytes: audioBytes}
			audio := tgbotapi.NewAudioUpload(update.Message.Chat.ID, audioFile)
			_, err := bot.Send(audio)
			if err != nil {
				log.Panic(err)
			}
		}

	}
}

func handleCommands(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) {
	command := msg.Command()

	switch command {
	case "start":
		msg := tgbotapi.NewMessage(msg.Chat.ID, "Ютуб видеоның ссылкасын жібер")
		bot.Send(msg)
	default:
		if !isYouTubeURL(msg.Text) {
			msg := tgbotapi.NewMessage(msg.Chat.ID, "Мен сізді түсінбедім. бастау үшін /start деп жазыңыз.")
			bot.Send(msg)
		}
	}
}

func isYouTubeURL(text string) bool {
	return strings.HasPrefix(text, "https")
}
