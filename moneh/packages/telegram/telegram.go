package telegram

import (
	"log"
	"moneh/configs"
	"strings"

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

	// Handle post flow
	case "waiting_for_flow_type":
		flowType := update.Message.Text
		HandleFlowTypeInput(&update, bot, flowType)
	case "waiting_for_flow_category":
		HandleFlowCategoryInput(update, bot)
	case "waiting_for_flow_name":
		HandleFlowNameInput(update, bot)
	case "waiting_for_flow_desc":
		HandleFlowDescInput(update, bot)
	case "waiting_for_flow_ammount":
		SubmitFlow(update, bot)

	// Handle post pocket
	case "waiting_for_pocket_type":
		pocketType := update.Message.Text
		HandlePocketTypeInput(&update, bot, pocketType)
	case "waiting_for_pocket_name":
		HandlePocketNameInput(update, bot)
	case "waiting_for_pocket_desc":
		HandlePocketDescInput(update, bot)
	case "waiting_for_pocket_limit":
		SubmitPocket(update, bot)

	// Others
	default:
		if update.Message.Chat.Type == "group" {
			if strings.HasPrefix(update.Message.Text, "/start/moneh") {
				HandleStartCommand(update, bot)
			} else if update.Message.Text == "/stop" {
				// HandleStopCommand(update, bot)
			}
		} else {
			// Private chat handling
			if update.Message.Text == "/start" {
				HandleStartCommand(update, bot)
			} else if update.Message.Text == "/stop" {
				// HandleStopCommand(update, bot)
			}
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
	case "menu_dashboard":
		HandleDashboardMenu(callback, bot)
	case "flow_get_list_flow":
		HandleGetAllFlow(callback, bot)
	case "flow_add":
		HandleAddFlow(callback, bot)
	case "flows_category_income", "flows_category_spending":
		flowType := callback.Data
		HandleFlowTypeInput(&update, bot, flowType)
	case "pockets_type_Bank":
		pocketType := callback.Data
		HandlePocketTypeInput(&update, bot, pocketType)
	case "pocket_get_list_pocket":
		HandleGetAllPocket(callback, bot)
	case "pocket_add":
		HandleAddPocket(callback, bot)
	default:
		bot.Send(tele_bot.NewMessage(userId, "Unknown option selected."))
	}

	// Acknowledge the callback query
	callbackResponse := tele_bot.NewCallback(callback.ID, "")
	bot.AnswerCallbackQuery(callbackResponse)
}
