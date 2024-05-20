package tests

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

type Record struct {
	Context   string
	Title     string
	Request   string
	Response  string
	CreatedAt time.Time
}

func WriteAudit(record Record) error {
	filename := "../../docs/test_audit.csv"
	fmt.Println(record)

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return fmt.Errorf("failed to open file for appending: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write record
	row := []string{
		record.Context,
		record.Title,
		record.Request,
		record.Response,
		record.CreatedAt.Format(time.RFC3339),
	}
	if err := writer.Write(row); err != nil {
		return fmt.Errorf("failed to write record: %w", err)
	}

	return nil
}
