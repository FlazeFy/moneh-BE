package telegram

import (
	"os"

	"github.com/tuotoo/qrcode"
)

func DecodeQRCode(url string) (string, error) {
	file, err := os.Open(url)
	if err != nil {
		return "", err
	}
	defer file.Close()

	qr, err := qrcode.Decode(file)
	if err != nil {
		return "", err
	}

	return qr.Content, nil
}
