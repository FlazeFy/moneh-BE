package telegram

import (
	"log"
	"moneh/configs"

	tele_bot "github.com/go-telegram-bot-api/telegram-bot-api"
)

var UserStates = make(map[int64]string)
var UserInputs = make(map[int64]map[string]string)

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
	if err != nil {
		log.Fatal(err)
	}

	for update := range updates {
		if update.Message != nil {
			handleMessage(update, bot)
		} else if update.CallbackQuery != nil {
			handleCallbackQuery(update, bot)
		}
	}
}

func handleMessage(update tele_bot.Update, bot *tele_bot.BotAPI) {
	userId := update.Message.Chat.ID

	switch UserStates[userId] {
	case "waiting_for_flow_type":
		HandleFlowTypeInput(update, bot)
	case "waiting_for_flow_category":
		HandleFlowCategoryInput(update, bot)
	default:
		if update.Message.Text == "/start" {
			HandleStartCommand(update, bot)
		}
	}
}

func handleCallbackQuery(update tele_bot.Update, bot *tele_bot.BotAPI) {
	callback := update.CallbackQuery
	userId := callback.Message.Chat.ID

	switch callback.Data {
	case "menu_flow":
		HandleFlowMenu(callback, bot)
	case "menu_pocket":
		HandlePocketMenu(callback, bot)
	case "flow_get_list_flow":
		HandleGetAllFlow(callback, bot)
	case "flow_add":
		HandleAddFlow(callback, bot)
	case "flow_stats":
		handleFlowStats(callback, bot)
	case "pocket_get_list_pocket":
		HandleGetAllPocket(callback, bot)
	case "pocket_add":
		handleAddPocket(callback, bot)
	case "pocket_stats":
		handlePocketStats(callback, bot)
	default:
		bot.Send(tele_bot.NewMessage(userId, "Unknown option selected."))
	}

	// Acknowledge the callback query
	callbackResponse := tele_bot.NewCallback(callback.ID, "")
	bot.AnswerCallbackQuery(callbackResponse)
}

func handleFlowStats(callback *tele_bot.CallbackQuery, bot *tele_bot.BotAPI) {
	userId := callback.Message.Chat.ID
	responseText := "Displaying flow stats..."
	msg := tele_bot.NewMessage(userId, responseText)
	bot.Send(msg)
}

func handleAddPocket(callback *tele_bot.CallbackQuery, bot *tele_bot.BotAPI) {
	// ...
}

func handlePocketStats(callback *tele_bot.CallbackQuery, bot *tele_bot.BotAPI) {
	userId := callback.Message.Chat.ID
	responseText := "Displaying pocket stats..."
	msg := tele_bot.NewMessage(userId, responseText)
	bot.Send(msg)
}
