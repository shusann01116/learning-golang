package main

import (
	"log/slog"

	"github.com/xuri/excelize/v2"
)

func main() {
	out := excelize.NewFile()
	out.SetCellValue("Sheet1", "A1", "Hello Excel!")
	if err := out.SaveAs("HelloExcel.xlsx"); err != nil {
		slog.Error("failed to save file", slog.Any("error", err))
		return
	}

	in, err := excelize.OpenFile("HelloExcel.xlsx")
	if err != nil {
		slog.Error("failed to open file", slog.Any("error", err))
		return
	}
	defer func() {
		err := in.Close()
		if err != nil {
			slog.Error("failed to close file", slog.Any("error", err))
		}
	}()

	value, err := in.GetCellValue("Sheet1", "A1")
	if err != nil {
		slog.Error("failed to read value", slog.Any("error", err))
		return
	}
	slog.Info("got value", slog.String("value", value))
}
