package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Allowed categories:
// Groceries
// Leisure
// Electronics
// Utilities
// Clothing
// Health
// Others

type ExpenseModel struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID        primitive.ObjectID `bson:"user_id" json:"user_id"`
	Title         string             `bson:"title" json:"title" binding:"required"`
	Amount        float64            `bson:"amount" json:"amount" binding:"required"`
	Category      string             `bson:"category" json:"category" binding:"required,oneof=Groceries Leisure Electronics Utilities Clothing Health Others"`
	Description   string             `bson:"description,omitempty" json:"description,omitempty"`
	ExpenseDate   time.Time          `bson:"expense_date" json:"expense_date" binding:"required"`
	CreatedAt     time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt     time.Time          `bson:"updated_at" json:"updated_at"`
	Type          string             `bson:"type" json:"type" binding:"required,oneof=Expense Income"`
	PaymentMethod string             `json:"payment_method" binding:"required"`
}
