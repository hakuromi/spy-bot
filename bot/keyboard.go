package bot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func MainKeyboard() tgbotapi.ReplyKeyboardMarkup {
	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("[dota] создать игру"),
			tgbotapi.NewKeyboardButton("[clash royale] создать игру"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("присоединиться к игре"),
			tgbotapi.NewKeyboardButton("начать игру"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("завершить игру"),
		),
	)

	keyboard.ResizeKeyboard = true
	return keyboard
}
