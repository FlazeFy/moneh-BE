package telegram

import (
	"moneh/modules/bots/flow"

	tele_bot "github.com/go-telegram-bot-api/telegram-bot-api"
)

func HandleGetAllFlow(callback *tele_bot.CallbackQuery, bot *tele_bot.BotAPI) {
	userId := callback.Message.Chat.ID
	bot.Send(tele_bot.NewMessage(userId, "Displaying all flows..."))

	res, err := flow.GetAllFlowBot()
	if err != nil {
		bot.Send(tele_bot.NewMessage(userId, err.Error()))
		return
	}
	bot.Send(tele_bot.NewMessage(userId, string(res)))
}
