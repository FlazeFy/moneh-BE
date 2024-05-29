package telegram

import (
	"log"
	"moneh/modules/bots/flow"

	tele_bot "github.com/go-telegram-bot-api/telegram-bot-api"
)

func HandleGetDailyFlow(callback *tele_bot.CallbackQuery, bot *tele_bot.BotAPI) {
	userId := callback.Message.Chat.ID

	bot.Send(tele_bot.NewMessage(userId, "How many days from now do you want to track?"))
	UserStates[userId] = "waiting_for_days_track"
}

func FetchDailyFLow(update *tele_bot.Update, bot *tele_bot.BotAPI, days string) {
	if update.Message == nil {
		log.Println("Message in update is nil")
		return
	}

	userId := update.Message.Chat.ID
	bot.Send(tele_bot.NewMessage(userId, "Displaying all flows..."))

	res, err := flow.GetAllFlowDaily(days)
	if err != nil {
		bot.Send(tele_bot.NewMessage(userId, err.Error()))
		return
	}
	bot.Send(tele_bot.NewMessage(userId, string(res)))
}
