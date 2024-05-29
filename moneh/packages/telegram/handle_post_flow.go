package telegram

import (
	"moneh/modules/bots/dct"

	tele_bot "github.com/go-telegram-bot-api/telegram-bot-api"
)

func HandleAddFlow(callback *tele_bot.CallbackQuery, bot *tele_bot.BotAPI) {
	userId := callback.Message.Chat.ID
	bot.Send(tele_bot.NewMessage(userId, "Preparing the field..."))

	inputFlowType := tele_bot.NewMessage(userId, "Select Flow Type : ")
	inputFlowType.ReplyMarkup = tele_bot.NewInlineKeyboardMarkup(
		tele_bot.NewInlineKeyboardRow(
			tele_bot.NewInlineKeyboardButtonData("Income", "flows_category_income"),
		),
		tele_bot.NewInlineKeyboardRow(
			tele_bot.NewInlineKeyboardButtonData("Spending", "flows_category_spending"),
		),
	)
	bot.Send(inputFlowType)

	UserStates[userId] = "waiting_for_flow_type"
	UserInputs[userId] = make(map[string]string)
}

func HandleFlowTypeInput(update tele_bot.Update, bot *tele_bot.BotAPI) {
	userId := update.Message.Chat.ID
	flowType := update.Message.Text

	if flowType != "flows_category_income" && flowType != "flows_category_spending" {
		bot.Send(tele_bot.NewMessage(userId, "Please select a valid flow type: Income or Spending."))
		return
	}

	UserInputs[userId]["flow_type"] = flowType
	UserStates[userId] = "waiting_for_flow_category"

	dct, err := dct.GetDctByType("flows_category")
	if err != nil {
		bot.Send(tele_bot.NewMessage(userId, err.Error()))
		return
	}

	var categoryButtons [][]tele_bot.InlineKeyboardButton
	for _, entry := range dct {
		btn := tele_bot.NewInlineKeyboardButtonData(entry.DctName, "flows_category_"+entry.DctName)
		categoryButtons = append(categoryButtons, tele_bot.NewInlineKeyboardRow(btn))
	}

	inputFlowCat := tele_bot.NewMessage(userId, "Select Category : ")
	inputFlowCat.ReplyMarkup = tele_bot.NewInlineKeyboardMarkup(categoryButtons...)

	bot.Send(inputFlowCat)
}

func HandleFlowCategoryInput(update tele_bot.Update, bot *tele_bot.BotAPI) {
	userId := update.Message.Chat.ID
	flowCategory := update.Message.Text

	UserInputs[userId]["flow_category"] = flowCategory
}
