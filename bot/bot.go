package bot

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Init(token string) *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic("Ошибка инициализации бота.", err)
	}

	return bot
}

func GetUpdates(bot *tgbotapi.BotAPI) tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	return bot.GetUpdatesChan(u)
}
