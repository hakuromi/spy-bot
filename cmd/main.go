package main

import (
	"log"

	"github.com/hakuromi/spy-bot/bot"
	"github.com/hakuromi/spy-bot/game"
)

func main() {
	var manager game.Manager
	manager.NewGame()

	botAPI := bot.Init("8578620255:AAHIQsDmMdA33Kpryao6AOC3U8m7OCQKYiQ")
	updates := bot.GetUpdates(botAPI)

	log.Println("Бот запущен.")

	for update := range updates {
		if update.Message != nil {
			switch update.Message.Text {
			case "/start":
				bot.HandleStart(botAPI, update.Message.Chat.ID)
			case "[dota] создать игру":
				bot.HandleNewGame(botAPI, update.Message.Chat.ID, update.Message.From.ID, update.Message.From.UserName, &manager, "dota")
			case "[clash royale] создать игру":
				bot.HandleNewGame(botAPI, update.Message.Chat.ID, update.Message.From.ID, update.Message.From.UserName, &manager, "clash")
			case "присоединиться к игре":
				bot.HandleJoin(botAPI, update.Message, &manager)
			case "начать игру":
				bot.HandleBegin(botAPI, update.Message.Chat.ID, &manager)
			case "завершить игру":
				bot.HandleEnd(botAPI, update.Message.Chat.ID, &manager)
			}

		}
	}
}
