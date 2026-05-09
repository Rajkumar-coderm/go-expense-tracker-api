package models

import (
	"strings"
	"time"
)

// Add Expense Request
type CreateExpenseRequestModel struct {
	Title         string   `json:"title" binding:"required"`
	Amount        float64  `json:"amount" binding:"required,gt=0"`
	Category      string   `json:"category" binding:"required,oneof=Groceries Leisure Electronics Utilities Clothing Health Others"`
	Description   string   `json:"description,omitempty"`
	ExpenseDate   JSONTime `json:"expense_date" binding:"required"`
	Type          string   `json:"type" binding:"required,oneof=Expense Income"`
	PaymentMethod string   `json:"payment_method" binding:"required"`
}

var ExpensePaymentMethods = map[string]bool{
	"cash":          true,
	"upi":           true,
	"credit_card":   true,
	"debit_card":    true,
	"wallet":        true,
	"bank_transfer": true,
	"cheque":        true,
	"other":         true,
}

var IncomePaymentMethods = map[string]bool{
	"bank_transfer": true,
	"upi":           true,
	"cash":          true,
	"cheque":        true,
	"paypal":        true,
	"stripe":        true,
	"razorpay":      true,
	"paytm":         true,
	"cupons":        true,
	"other":         true,
}

type JSONTime time.Time

// UnmarshalJSON handles the incoming user format
func (jt *JSONTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02 15:04:05", s)
	if err != nil {
		return err
	}
	*jt = JSONTime(t)
	return nil
}

// Cast back to time.Time easily
func (jt JSONTime) Time() time.Time {
	return time.Time(jt)
}

// Update Expense Request
type UpdateExpenseRequestModel struct {
	Title       string   `json:"title,omitempty"`
	Amount      float64  `json:"amount,omitempty"`
	Category    string   `json:"category,omitempty" binding:"omitempty,oneof=Groceries Leisure Electronics Utilities Clothing Health Others"`
	Description string   `json:"description,omitempty"`
	ExpenseDate JSONTime `json:"expense_date"`
}

type ExpenseGetQueryModel struct {
	DateFilter ExpenseFilterQueryModel `json:"date"`
	Page       int                     `json:"page,omitempty"`
	Limit      int                     `json:"limit,omitempty"`
	UserId     string                  `json:"id,omitempty"`
	Type       string                  `json:"type,omitempty" binding:"omitempty,oneof=Expense Income"`
}

// Query Params Example:
// /expenses?filter=week
// /expenses?filter=month
// /expenses?filter=3months
// /expenses?filter=custom&start_date=2026-01-01&end_date=2026-02-01

type ExpenseFilterQueryModel struct {
	Filter    string `form:"filter" binding:"omitempty,oneof=week month 3months custom"`
	StartDate string `form:"start_date,omitempty"`
	EndDate   string `form:"end_date,omitempty"`
}
