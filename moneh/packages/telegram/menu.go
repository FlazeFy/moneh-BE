package telegram

import (
	tele_bot "github.com/go-telegram-bot-api/telegram-bot-api"
)

func HandleStartCommand(update tele_bot.Update, bot *tele_bot.BotAPI) {
	userId := update.Message.Chat.ID
	firstName := update.Message.From.FirstName

	msg := tele_bot.NewMessage(userId, "Hello "+firstName+"! Welcome to Moneh Bot")
	msg.ReplyMarkup = tele_bot.NewInlineKeyboardMarkup(
		tele_bot.NewInlineKeyboardRow(
			tele_bot.NewInlineKeyboardButtonData("Flow", "menu_flow"),
		),
		tele_bot.NewInlineKeyboardRow(
			tele_bot.NewInlineKeyboardButtonData("Pocket", "menu_pocket"),
		),
	)
	bot.Send(msg)
	bot.Send(tele_bot.NewMessage(userId, "Select what menu do you want to do"))
}

func HandleFlowMenu(callback *tele_bot.CallbackQuery, bot *tele_bot.BotAPI) {
	userId := callback.Message.Chat.ID
	msg := tele_bot.NewMessage(userId, "Selected Menu : Flow")
	msg.ReplyMarkup = tele_bot.NewInlineKeyboardMarkup(
		tele_bot.NewInlineKeyboardRow(
			tele_bot.NewInlineKeyboardButtonData("See All Flow", "flow_get_list_flow"),
		),
		tele_bot.NewInlineKeyboardRow(
			tele_bot.NewInlineKeyboardButtonData("Add Flow", "flow_add"),
		),
		tele_bot.NewInlineKeyboardRow(
			tele_bot.NewInlineKeyboardButtonData("See Stats", "flow_stats"),
		),
	)
	bot.Send(msg)
}

func HandlePocketMenu(callback *tele_bot.CallbackQuery, bot *tele_bot.BotAPI) {
	userId := callback.Message.Chat.ID
	msg := tele_bot.NewMessage(userId, "Selected Menu : Pocket")
	msg.ReplyMarkup = tele_bot.NewInlineKeyboardMarkup(
		tele_bot.NewInlineKeyboardRow(
			tele_bot.NewInlineKeyboardButtonData("See All Pocket", "pocket_get_list_pocket"),
		),
		tele_bot.NewInlineKeyboardRow(
			tele_bot.NewInlineKeyboardButtonData("Add Pocket", "pocket_add"),
		),
		tele_bot.NewInlineKeyboardRow(
			tele_bot.NewInlineKeyboardButtonData("See Stats", "pocket_stats"),
		),
	)
	bot.Send(msg)
}
