package telegram

import (
	"fmt"
	"log"
	"moneh/modules/bots/dct"
	"moneh/modules/flows/models"
	"moneh/modules/flows/repositories"
	"moneh/packages/helpers/converter"
	"strconv"
	"strings"

	tele_bot "github.com/go-telegram-bot-api/telegram-bot-api"
)

func HandleAddFlow(callback *tele_bot.CallbackQuery, bot *tele_bot.BotAPI) {
	userId := callback.Message.Chat.ID
	bot.Send(tele_bot.NewMessage(userId, "Preparing the field..."))

	inputFlowType := tele_bot.NewMessage(userId, "Select Flow Type : ")
	inputFlowType.ReplyMarkup = tele_bot.NewInlineKeyboardMarkup(
		tele_bot.NewInlineKeyboardRow(
			tele_bot.NewInlineKeyboardButtonData("Income", "flows_type_income"),
		),
		tele_bot.NewInlineKeyboardRow(
			tele_bot.NewInlineKeyboardButtonData("Spending", "flows_type_spending"),
		),
	)
	bot.Send(inputFlowType)

	UserStates[userId] = "waiting_for_flow_type"
	UserInputs[userId] = make(map[string]string)
}

func HandleFlowTypeInput(update *tele_bot.Update, bot *tele_bot.BotAPI, flowType string) {
	if update.Message == nil {
		log.Println("Message in update is nil")
		return
	}

	userId := update.Message.Chat.ID

	if flowType != "flows_type_income" && flowType != "flows_type_spending" {
		bot.Send(tele_bot.NewMessage(userId, "Please select a valid flow type: Income or Spending."))
		return
	}

	UserInputs[userId]["flows_type"] = flowType
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

	UserInputs[userId]["flows_category"] = flowCategory

	bot.Send(tele_bot.NewMessage(userId, "Type your flow name :"))
	UserStates[userId] = "waiting_for_flow_name"
}

func HandleFlowNameInput(update tele_bot.Update, bot *tele_bot.BotAPI) {
	userId := update.Message.Chat.ID
	flowName := update.Message.Text

	UserInputs[userId]["flows_name"] = flowName

	bot.Send(tele_bot.NewMessage(userId, "Type your flow desc :"))
	UserStates[userId] = "waiting_for_flow_desc"
}

func HandleFlowDescInput(update tele_bot.Update, bot *tele_bot.BotAPI) {
	userId := update.Message.Chat.ID
	flowDesc := update.Message.Text

	UserInputs[userId]["flows_desc"] = flowDesc

	bot.Send(tele_bot.NewMessage(userId, "Type your flow ammount (Rp.) :"))
	UserStates[userId] = "waiting_for_flow_ammount"
}

func SubmitFlow(update tele_bot.Update, bot *tele_bot.BotAPI) {
	userId := update.Message.Chat.ID
	flowAmmount := update.Message.Text

	UserInputs[userId]["flows_ammount"] = flowAmmount

	// Clean
	UserInputs[userId]["flows_category"] = strings.Replace(UserInputs[userId]["flows_category"], "flows_category_", "", 1)
	UserInputs[userId]["flows_type"] = strings.Replace(UserInputs[userId]["flows_type"], "flows_type_", "", 1)

	var res strings.Builder

	intAmount, err := strconv.Atoi(UserInputs[userId]["flows_ammount"])
	if err != nil {
		bot.Send(tele_bot.NewMessage(userId, "Something error with flow ammount : "+err.Error()))
		return
	}

	amount := converter.ConvertPriceNumber(intAmount)
	res.WriteString(fmt.Sprintf(`
			Detail of submited flow

			Type : %s
			Category : %s
			Name : %s
			Ammount : Rp. %s,00

			Notes : %s
		`,
		UserInputs[userId]["flows_type"],
		UserInputs[userId]["flows_category"],
		UserInputs[userId]["flows_name"],
		amount,
		UserInputs[userId]["flows_desc"],
	))

	bot.Send(tele_bot.NewMessage(userId, res.String()))
	bot.Send(tele_bot.NewMessage(userId, "Sending flow..."))

	var obj models.GetFlow

	obj.FlowsType = UserInputs[userId]["flows_type"]
	obj.FlowsCategory = UserInputs[userId]["flows_category"]
	obj.FlowsName = UserInputs[userId]["flows_name"]
	obj.FlowsDesc = UserInputs[userId]["flows_desc"]
	obj.FlowsAmmount = intAmount
	obj.FlowsTag = "null"
	obj.IsShared = 0

	result, err := repositories.PostFlow(obj)
	if err != nil {
		bot.Send(tele_bot.NewMessage(userId, err.Error()))
		return
	}

	bot.Send(tele_bot.NewMessage(userId, result.Message))
}
