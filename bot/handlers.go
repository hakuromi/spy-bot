package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/hakuromi/spy-bot/game"
	"github.com/hakuromi/spy-bot/models"
)

func HandleStart(bot *tgbotapi.BotAPI, chatID int64) {
	message := "Здарова! Это игра в шпиона. Чтобы присоединиться к игре, используй /join"
	msg := tgbotapi.NewMessage(chatID, message)
	bot.Send(msg)
}

func HandleNewGame(bot *tgbotapi.BotAPI, chatID int64, manager *game.Manager) {
	manager.NewGame()
	message := "Новая игра создана! Игроки могут присоединяться с помощью /join!"
	msg := tgbotapi.NewMessage(chatID, message)
	bot.Send(msg)
}

func HandleJoin(bot *tgbotapi.BotAPI, message *tgbotapi.Message, manager *game.Manager) {
	player := models.Player{
		ID:   message.From.ID,
		Name: message.From.UserName,
	}

	err := manager.AddPlayer(player)
	var reply string
	if err != nil {
		reply = "Игрок не присоединен."
	} else {
		reply = "Игрок" + player.Name + "присоединился!"
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, reply)
	bot.Send(msg)
}

func HandleBegin(bot *tgbotapi.BotAPI, chatID int64, manager *game.Manager) {
	if !manager.CanStart() {
		message := "Нельзя начать игру. Нужно минимум 3 игрока."
		for _, p := range manager.Game.Players {
			msg := tgbotapi.NewMessage(p.ID, message)
			bot.Send(msg)
		}
		return
	}
	manager.Start()

	msg := tgbotapi.NewMessage(chatID, "Игра началась!")
	bot.Send(msg)
	roles := manager.GetRoles()
	for userID, role := range roles {
		msg := tgbotapi.NewMessage(userID, role)
		bot.Send(msg)
	}
}
