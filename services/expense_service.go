package services

import (
	"encoding/csv"
	"errors"
	"io"
	"time"

	"github.com/expense-tracker-api/models"
	"github.com/expense-tracker-api/repository"
	"github.com/expense-tracker-api/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateNewExpense(req models.CreateExpenseRequestModel, userId string) error {
	objID, _ := primitive.ObjectIDFromHex(userId)
	re := models.ExpenseModel{
		UserID:        objID,
		Title:         req.Title,
		Description:   req.Description,
		Amount:        req.Amount,
		ExpenseDate:   req.ExpenseDate.Time(),
		Category:      req.Category,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		Type:          req.Type,
		PaymentMethod: req.PaymentMethod,
	}
	return repository.CreateNewExpense(re)
}

func GetExpense(req models.ExpenseGetQueryModel) ([]models.ExpenseModel, error) {
	return repository.GetExpense(req, false)
}

func DeleteExpense(id, userId string) error {
	return repository.DeleteExpense(id, userId)
}

func ImportExpenses(userId string, file io.Reader) (*models.CSVImportResponse, error) {

	reader := csv.NewReader(file)
	reader.TrimLeadingSpace = true

	records, err := reader.ReadAll()
	if err != nil {
		return nil, errors.New("invalid CSV format")
	}

	if len(records) < 2 {
		return nil, errors.New("CSV must contain header and data")
	}

	headerMap := utils.BuildHeaderMap(records[0])

	var successCount, failedCount int
	var failedRows []models.CSVRowError
	var validExpenses []models.ExpenseModel

	userObjID, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, errors.New("invalid user ID")
	}

	now := time.Now()

	for i, row := range records[1:] {
		rowNumber := i + 2

		expense, err := utils.ParseExpenseRow(row, headerMap)
		if err != nil {
			failedCount++
			failedRows = append(failedRows, models.CSVRowError{
				RowNumber: rowNumber,
				Error:     err.Error(),
			})
			continue
		}

		expense.UserID = userObjID
		expense.CreatedAt = now
		expense.UpdatedAt = now

		validExpenses = append(validExpenses, *expense)
		successCount++
	}
	if len(validExpenses) > 0 {
		err = repository.ImportExpenses(validExpenses)
		if err != nil {
			return nil, err
		}
	}

	return &models.CSVImportResponse{
		TotalRows:    len(records) - 1,
		SuccessCount: successCount,
		FailedCount:  failedCount,
		FailedRows:   failedRows,
	}, nil
}

func ExportExpenses(
	format string,
	req models.ExpenseGetQueryModel,
) ([]byte, string, string, error) {

	expenses, err := repository.GetExpense(req, true)
	if err != nil {
		return nil, "", "", err
	}

	switch format {

	case "csv":
		data, err := utils.GenerateCSV(expenses)
		return data,
			"expenses.csv",
			"text/csv",
			err

	case "xlsx":
		data, err := utils.GenerateXLSX(expenses)
		return data,
			"expenses.xlsx",
			"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
			err

	default:
		return nil, "", "", errors.New("unsupported export format")
	}
}
