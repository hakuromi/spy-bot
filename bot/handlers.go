package bot

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/hakuromi/spy-bot/game"
	"github.com/hakuromi/spy-bot/models"
)

func HandleStart(bot *tgbotapi.BotAPI, chatID int64) {
	message := "здарова! это игра в шпиона. присоединяйся к игре или создай свою!"
	msg := tgbotapi.NewMessage(chatID, message)

	msg.ReplyMarkup = MainKeyboard()
	bot.Send(msg)
}

func HandleNewGame(bot *tgbotapi.BotAPI, chatID int64, userID int64, username string, manager *game.Manager) {
	if len(manager.Game.Players) > 0 {
		msg := tgbotapi.NewMessage(chatID, "игра уже существует!")
		bot.Send(msg)
		return
	}
	manager.NewGame()

	player := models.Player{
		ID:   userID,
		Name: username,
	}

	manager.AddPlayer(player)
	message := "новая игра создана! \nадмин @" + username + ". зови друзей!"
	msg := tgbotapi.NewMessage(chatID, message)
	bot.Send(msg)
}

func HandleJoin(bot *tgbotapi.BotAPI, message *tgbotapi.Message, manager *game.Manager) {
	if len(manager.Game.Players) == 0 {
		msg := tgbotapi.NewMessage(message.Chat.ID, "игра еще не создана.\nсначала создайте игру.")
		bot.Send(msg)
		return
	}

	if manager.Game.Active {
		msg := tgbotapi.NewMessage(message.Chat.ID, "игра уже идет. присоединиться нельзя.")
		bot.Send(msg)
		return
	}

	player := models.Player{
		ID:   message.From.ID,
		Name: message.From.UserName,
	}

	err := manager.AddPlayer(player)
	var reply string

	if err != nil {
		if err.Error() == "игрок уже в игре." {
			reply = "вы уже присоединились к игре."
		} else {
			reply = "игрок не присоединен." + err.Error()
		}
	} else {
		reply = fmt.Sprintf("игрок @%s присоединился! \nвсего игроков: %d.", player.Name, len(manager.Game.Players))
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, reply)
	bot.Send(msg)
}

func HandleBegin(bot *tgbotapi.BotAPI, chatID int64, manager *game.Manager) {
	if !manager.CanStart() {
		message := "нельзя начать игру.\nнужно минимум 3 игрока.\n" + game.PlayerList(manager.Game.Players)
		msg := tgbotapi.NewMessage(chatID, message)
		bot.Send(msg)
		return
	}
	manager.Start()

	msg := tgbotapi.NewMessage(chatID, "стартуем!\n"+game.PlayerList(manager.Game.Players))
	bot.Send(msg)
	roles := manager.GetRoles()
	for userID, role := range roles {
		msg := tgbotapi.NewMessage(userID, role)
		bot.Send(msg)
	}
}

func HandleEnd(bot *tgbotapi.BotAPI, chatID int64, manager *game.Manager) {
	if len(manager.Game.Players) == 0 {
		msg := tgbotapi.NewMessage(chatID, "игра ещё не создана.")
		bot.Send(msg)
		return
	}

	if !manager.Game.Active {
		msg := tgbotapi.NewMessage(chatID, "игра уже завершена.")
		bot.Send(msg)
		return
	}

	manager.End()

	msg := tgbotapi.NewMessage(chatID, "игра завершена. можно начинать новую партию или позвать кого-нибудь еще.")
	bot.Send(msg)
}
