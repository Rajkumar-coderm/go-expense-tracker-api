package utils

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/expense-tracker-api/models"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateShortCode(length int) string {
	code := make([]byte, length)
	for i := range code {
		code[i] = charset[rand.Intn(len(charset))]
	}
	return string(code)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	return string(bytes), err
}

func CheckPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword),
		[]byte(password),
	)
}

type ApiError struct {
	Field string `json:"field"`
	Msg   string `json:"message"`
}

func FormatValidationError(err error) []ApiError {
	var ve validator.ValidationErrors
	var out []ApiError

	if errors.As(err, &ve) {
		for _, fe := range ve {
			msg := getCustomMsg(fe)
			out = append(out, ApiError{
				Field: fe.Field(),
				Msg:   fmt.Sprintf("The field '%s' %s", fe.Field(), msg),
			})
		}
		return out
	}

	return []ApiError{{Field: "request", Msg: "Invalid JSON format"}}
}

func FormatDBError(err error) (int, map[string]string) {
	if mongo.IsDuplicateKeyError(err) {
		return 409, map[string]string{
			"msg": "This resource (or short code) already exists. Please try again.",
		}
	}

	msg := err.Error()

	if strings.Contains(msg, "NotFound") || errors.Is(err, mongo.ErrNoDocuments) {
		return 404, map[string]string{
			"msg": "The requested resource was not found.",
		}
	}

	if strings.Contains(msg, "NotAccessToResouce") {
		return 403, map[string]string{
			"msg": "You don't have permission to perform this action."}
	}

	return 500, map[string]string{
		"msg": err.Error(),
	}
}

func getCustomMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "is mandatory and cannot be empty"
	case "url":
		return "must be a valid URL"
	case "oneof":
		return "must be one of the following: " + fe.Param()
	default:
		return "is invalid"
	}
}

func ParseExpenseRow(row []string, headerMap map[string]int) (*models.ExpenseModel, error) {

	get := func(field string) string {
		if idx, ok := headerMap[field]; ok && idx < len(row) {
			return strings.TrimSpace(row[idx])
		}
		return ""
	}

	title := get("title")
	amountStr := get("amount")
	category := get("category")
	dateStr := get("expense_date")
	typeStr := get("type")
	paymentMode := get("payment_method")

	if paymentMode == "" {
		return nil, errors.New("payment_method is required")
	}

	if !ValidatePaymentMethod(strings.ToLower(typeStr), strings.ToLower(paymentMode)) {
		return nil, errors.New("invalid payment method for the given transaction type")
	}

	if title == "" {
		return nil, errors.New("title is required")
	}

	if category == "" {
		return nil, errors.New("category is required")
	}

	if typeStr == "" {
		return nil, errors.New("type is required")
	}

	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil || amount <= 0 {
		return nil, errors.New("invalid amount")
	}

	expenseDate, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return nil, errors.New("invalid expense_date format, use YYYY-MM-DD")
	}

	return &models.ExpenseModel{
		Title:         title,
		Amount:        amount,
		Category:      category,
		ExpenseDate:   expenseDate,
		Type:          typeStr,
		PaymentMethod: paymentMode,
	}, nil
}

func BuildHeaderMap(header []string) map[string]int {
	m := make(map[string]int)
	for i, col := range header {
		m[strings.ToLower(strings.TrimSpace(col))] = i
	}
	return m
}

func ValidatePaymentMethod(
	transactionType string,
	paymentMethod string,
) bool {

	switch transactionType {

	case "expense":
		return models.ExpensePaymentMethods[paymentMethod]

	case "income":
		return models.IncomePaymentMethods[paymentMethod]

	default:
		return false
	}
}
