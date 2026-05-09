package utils

import (
	"bytes"
	"encoding/csv"
	"fmt"

	"github.com/expense-tracker-api/models"
	"github.com/xuri/excelize/v2"
)

func GenerateCSV(expenses []models.ExpenseModel) ([]byte, error) {

	var buffer bytes.Buffer
	writer := csv.NewWriter(&buffer)

	writer.Write([]string{
		"title",
		"amount",
		"category",
		"expense_date",
		"type",
		"payment_method",
	})

	for _, e := range expenses {
		writer.Write([]string{
			e.Title,
			fmt.Sprintf("%.2f", e.Amount),
			e.Category,
			e.ExpenseDate.Format("2006-01-02"),
			e.Type,
		})
	}

	writer.Flush()

	return buffer.Bytes(), writer.Error()
}

func GenerateXLSX(expenses []models.ExpenseModel) ([]byte, error) {

	f := excelize.NewFile()

	sheet := "Expenses"
	f.SetSheetName("Sheet1", sheet)

	// header
	headers := []string{
		"Title",
		"Amount",
		"Category",
		"Expense Date",
		"Type",
		"Payment Method",
	}

	for i, h := range headers {
		cell := fmt.Sprintf("%c1", 'A'+i)
		f.SetCellValue(sheet, cell, h)
	}

	// rows
	for i, e := range expenses {

		row := i + 2

		f.SetCellValue(sheet, fmt.Sprintf("A%d", row), e.Title)
		f.SetCellValue(sheet, fmt.Sprintf("B%d", row), e.Amount)
		f.SetCellValue(sheet, fmt.Sprintf("C%d", row), e.Category)
		f.SetCellValue(
			sheet,
			fmt.Sprintf("D%d", row),
			e.ExpenseDate.Format("2006-01-02"),
		)
		f.SetCellValue(sheet, fmt.Sprintf("E%d", row), e.Type)
		f.SetCellValue(sheet, fmt.Sprintf("F%d", row), e.PaymentMethod)
	}

	buffer, err := f.WriteToBuffer()
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
