package telegram

import (
	"moneh/modules/bots/pocket"

	tele_bot "github.com/go-telegram-bot-api/telegram-bot-api"
)

func HandleGetAllPocket(callback *tele_bot.CallbackQuery, bot *tele_bot.BotAPI) {
	userId := callback.Message.Chat.ID
	bot.Send(tele_bot.NewMessage(userId, "Displaying all pockets..."))

	res, err := pocket.GetAllPocketBot()
	if err != nil {
		bot.Send(tele_bot.NewMessage(userId, err.Error()))
		return
	}
	bot.Send(tele_bot.NewMessage(userId, string(res)))
}
