package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/hakuromi/spy-bot/bot"
	"github.com/hakuromi/spy-bot/game"
)

func main() {
	var manager game.Manager
	manager.NewGame()

	botAPI := bot.Init("8578620255:AAHIQsDmMdA33Kpryao6AOC3U8m7OCQKYiQ")
	updates := bot.GetUpdates(botAPI)

	commands := []tgbotapi.BotCommand{
		{Command: "start", Description: "Запуск бота"},
		{Command: "newgame", Description: "Создать новую игру"},
		{Command: "join", Description: "Присоединиться к игре"},
		{Command: "begin", Description: "Начать игру"},
	}

	config := tgbotapi.NewSetMyCommands(commands...)
	_, err := botAPI.Request(config)
	if err != nil {
		log.Println("Не удалось зарегистрировать команды:", err)
	}

	log.Println("Бот запущен.")

	for update := range updates {
		if update.Message != nil {
			switch update.Message.Text {
			case "/start":
				bot.HandleStart(botAPI, update.Message.Chat.ID)
			case "/newgame":
				bot.HandleNewGame(botAPI, update.Message.Chat.ID, &manager)
			case "/join":
				bot.HandleJoin(botAPI, update.Message, &manager)
			case "/begin":
				bot.HandleBegin(botAPI, update.Message.Chat.ID, &manager)
			}

		}
	}
}
