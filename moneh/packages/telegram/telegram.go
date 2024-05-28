package telegram

import (
	"log"
	"moneh/configs"
	"moneh/modules/bots/flow"

	tele_bot "github.com/go-telegram-bot-api/telegram-bot-api"
)

func InitTeleBot() {
	bot, err := tele_bot.NewBotAPI(configs.GetConfigTele().TELE_TOKEN)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tele_bot.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			if update.Message.Text == "/start" {
				firstName := update.Message.From.FirstName
				msg := tele_bot.NewMessage(update.Message.Chat.ID, "Hello "+firstName+"! Welcome to Moneh Bot")

				msg.ReplyMarkup = tele_bot.NewInlineKeyboardMarkup(
					tele_bot.NewInlineKeyboardRow(
						tele_bot.NewInlineKeyboardButtonData("Flow", "menu_flow"),
						tele_bot.NewInlineKeyboardButtonData("Pocket", "menu_pocket"),
					),
				)
				bot.Send(msg)
				bot.Send(tele_bot.NewMessage(update.Message.Chat.ID, "Select what menu do you want to do"))
			}
		} else if update.CallbackQuery != nil {
			callback := update.CallbackQuery

			switch callback.Data {

			case "menu_flow":
				msg := tele_bot.NewMessage(callback.Message.Chat.ID, "Selected Menu : Flow")
				msg.ReplyMarkup = tele_bot.NewInlineKeyboardMarkup(
					tele_bot.NewInlineKeyboardRow(
						tele_bot.NewInlineKeyboardButtonData("See All Flow", "flow_get_list_flow"),
						tele_bot.NewInlineKeyboardButtonData("See Stats", "flow_stats"),
					),
				)
				bot.Send(msg)

			case "menu_pocket":
				msg := tele_bot.NewMessage(callback.Message.Chat.ID, "Selected Menu : Pocket")
				msg.ReplyMarkup = tele_bot.NewInlineKeyboardMarkup(
					tele_bot.NewInlineKeyboardRow(
						tele_bot.NewInlineKeyboardButtonData("See All Pocket", "pocket_get_list_pocket"),
						tele_bot.NewInlineKeyboardButtonData("See Stats", "pocket_stats"),
					),
				)
				bot.Send(msg)

			case "flow_get_list_flow":
				bot.Send(tele_bot.NewMessage(callback.Message.Chat.ID, "Displaying all flows..."))
				res, _ := flow.GetAllFlowBot()
				if err != nil {
					bot.Send(tele_bot.NewMessage(callback.Message.Chat.ID, err.Error()))
				}
				bot.Send(tele_bot.NewMessage(callback.Message.Chat.ID, string(res)))

			case "flow_stats":
				responseText := "Displaying flow stats..."
				msg := tele_bot.NewMessage(callback.Message.Chat.ID, responseText)
				bot.Send(msg)
			case "pocket_get_list_pocket":
				responseText := "Displaying all pockets..."
				msg := tele_bot.NewMessage(callback.Message.Chat.ID, responseText)
				bot.Send(msg)
			case "pocket_stats":
				responseText := "Displaying pocket stats..."
				msg := tele_bot.NewMessage(callback.Message.Chat.ID, responseText)
				bot.Send(msg)
			default:
				responseText := "Unknown option selected."
				msg := tele_bot.NewMessage(callback.Message.Chat.ID, responseText)
				bot.Send(msg)
			}

			// Acknowledge the callback query
			callbackResponse := tele_bot.NewCallback(callback.ID, "")
			bot.AnswerCallbackQuery(callbackResponse)
		}
	}

}
