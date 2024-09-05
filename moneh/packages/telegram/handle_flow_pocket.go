package telegram

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"moneh/modules/pockets/models"
	"net/http"
	"os"
	"strconv"
	"time"

	tele_bot "github.com/go-telegram-bot-api/telegram-bot-api"
)

type APIResponsePocket struct {
	Status  int                      `json:"status"`
	Message string                   `json:"message"`
	Data    []models.GetPocketExport `json:"data"`
}

func HandlePocketExport(callback *tele_bot.CallbackQuery, bot *tele_bot.BotAPI) {
	resp, err := http.Get("http://127.0.0.1:1323/api/v2/pockets")
	if err != nil {
		bot.Send(tele_bot.NewMessage(callback.Message.Chat.ID, "Something is wrong : failed to fetch data "+err.Error()))
		return
	}
	defer resp.Body.Close()

	var res APIResponsePocket
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		bot.Send(tele_bot.NewMessage(callback.Message.Chat.ID, "Something is wrong : failed to parse API response "+err.Error()))
		return
	}

	var csvBuffer bytes.Buffer
	writer := csv.NewWriter(&csvBuffer)
	writer.Write([]string{"Pocket Name", "Description", "Type", "Limit", "Created At"})

	for _, data := range res.Data {
		writer.Write([]string{
			data.PocketsName,
			data.PocketsDesc,
			data.PocketsType,
			strconv.Itoa(data.PocketsLimit),
			data.CreatedAt,
		})
	}
	writer.Flush()

	if err := writer.Error(); err != nil {
		bot.Send(tele_bot.NewMessage(callback.Message.Chat.ID, "Something is wrong : failed to generate CSV "+err.Error()))
		return
	}

	dt := time.Now().Format("2006-01-02 15:04:05")
	csvFile, err := os.CreateTemp("", "pockets-export-"+dt+"-*.csv")
	if err != nil {
		bot.Send(tele_bot.NewMessage(callback.Message.Chat.ID, "Something is wrong : failed to create temp file "+err.Error()))
		return
	}
	defer os.Remove(csvFile.Name())

	if _, err := csvFile.Write(csvBuffer.Bytes()); err != nil {
		bot.Send(tele_bot.NewMessage(callback.Message.Chat.ID, "Something is wrong : failed to write CSV to file "+err.Error()))
		return
	}
	csvFile.Close()

	msg := tele_bot.NewDocumentUpload(callback.Message.Chat.ID, csvFile.Name())
	msg.Caption = "Here is your exported pockets in CSV"
	if _, err := bot.Send(msg); err != nil {
		bot.Send(tele_bot.NewMessage(callback.Message.Chat.ID, "Failed to send CSV: "+err.Error()))
	}
}
