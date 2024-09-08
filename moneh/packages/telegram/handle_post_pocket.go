package telegram

import (
	"fmt"
	"log"
	"moneh/modules/bots/dct"
	"moneh/modules/pockets/models"
	"moneh/modules/pockets/repositories"
	"moneh/packages/helpers/converter"
	"strconv"
	"strings"

	tele_bot "github.com/go-telegram-bot-api/telegram-bot-api"
)

func HandleAddPocket(callback *tele_bot.CallbackQuery, bot *tele_bot.BotAPI) {
	userId := callback.Message.Chat.ID
	bot.Send(tele_bot.NewMessage(userId, "Preparing the field..."))

	dct, err := dct.GetDctByType("pockets_type")
	if err != nil {
		bot.Send(tele_bot.NewMessage(userId, err.Error()))
		return
	}

	var typeButtons [][]tele_bot.InlineKeyboardButton
	for _, entry := range dct {
		btn := tele_bot.NewInlineKeyboardButtonData(entry.DctName, "pockets_type_"+entry.DctName)
		typeButtons = append(typeButtons, tele_bot.NewInlineKeyboardRow(btn))
	}

	inputPocketType := tele_bot.NewMessage(userId, "Select Type : ")
	inputPocketType.ReplyMarkup = tele_bot.NewInlineKeyboardMarkup(typeButtons...)

	bot.Send(inputPocketType)

	UserStates[userId] = "waiting_for_pocket_type"
	UserInputs[userId] = make(map[string]string)
}

func HandlePocketTypeInput(update *tele_bot.Update, bot *tele_bot.BotAPI, pocketType string) {
	if update.Message == nil {
		log.Println("Message in update is nil")
		return
	}

	userId := update.Message.Chat.ID

	UserInputs[userId]["pockets_type"] = pocketType

	bot.Send(tele_bot.NewMessage(userId, "Type your pocket name :"))
	UserStates[userId] = "waiting_for_pocket_name"
}

func HandlePocketNameInput(update tele_bot.Update, bot *tele_bot.BotAPI) {
	userId := update.Message.Chat.ID
	pocketName := update.Message.Text

	UserInputs[userId]["pockets_name"] = pocketName

	bot.Send(tele_bot.NewMessage(userId, "Type your pocket desc :"))
	UserStates[userId] = "waiting_for_pocket_desc"
}

func HandlePocketDescInput(update tele_bot.Update, bot *tele_bot.BotAPI) {
	userId := update.Message.Chat.ID
	pocketDesc := update.Message.Text

	UserInputs[userId]["pockets_desc"] = pocketDesc

	bot.Send(tele_bot.NewMessage(userId, "Type your pocket limit (Rp.) :"))
	UserStates[userId] = "waiting_for_pocket_limit"
}

func SubmitPocket(update tele_bot.Update, bot *tele_bot.BotAPI) {
	userId := update.Message.Chat.ID
	pocketLimit := update.Message.Text

	UserInputs[userId]["pockets_limit"] = pocketLimit

	// Clean
	UserInputs[userId]["pockets_type"] = strings.Replace(UserInputs[userId]["pockets_type"], "pockets_type_", "", 1)

	var res strings.Builder

	intLimit, err := strconv.Atoi(UserInputs[userId]["pockets_limit"])
	if err != nil {
		bot.Send(tele_bot.NewMessage(userId, "Something error with pockets limit : "+err.Error()))
		return
	}

	limit := converter.ConvertPriceNumber(intLimit)
	res.WriteString(fmt.Sprintf(`
			Detail of submited pocket

			Type : %s
			Name : %s
			Limit : Rp. %s,00

			Notes : %s
		`,
		UserInputs[userId]["pockets_type"],
		UserInputs[userId]["pockets_name"],
		limit,
		UserInputs[userId]["pockets_desc"],
	))

	bot.Send(tele_bot.NewMessage(userId, res.String()))
	bot.Send(tele_bot.NewMessage(userId, "Sending flow..."))

	var obj models.GetPocketHeaders

	obj.PocketsType = UserInputs[userId]["pockets_type"]
	obj.PocketsName = UserInputs[userId]["pockets_name"]
	obj.PocketsDesc = UserInputs[userId]["pockets_desc"]
	obj.PocketsLimit = intLimit

	result, err := repositories.PostPocket(obj, fmt.Sprintf("%d", userId))
	if err != nil {
		bot.Send(tele_bot.NewMessage(userId, err.Error()))
		return
	}

	bot.Send(tele_bot.NewMessage(userId, result.Message))
}
