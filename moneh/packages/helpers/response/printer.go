package response

import (
	"fmt"
	"os"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func ResponsePrinter(types, ctx, body string) {
	msgFailed := "Failed to print"
	dt := time.Now().Format("2006-01-02 15:04:05")
	purpose := "log_" + ctx + "_" + dt + ".txt"

	if types == "txt" {
		d1 := []byte(body)
		err := os.WriteFile("factories/logs/"+purpose, d1, 0644)
		check(err)
	} else {
		fmt.Println("\n" + msgFailed + " : File type is not valid")
	}
}
