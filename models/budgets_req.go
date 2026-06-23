package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BudgetModel struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID    primitive.ObjectID `bson:"user_id" json:"user_id"`
	Category  string             `bson:"category" json:"category"`
	Amount    float64            `bson:"amount" json:"amount"`
	Period    string             `bson:"period" json:"period"`
	StartDate time.Time          `bson:"start_date" json:"start_date"`
	EndDate   time.Time          `bson:"end_date" json:"end_date"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}

type CreateBudgetRequestModel struct {
	Category string  `json:"category" binding:"required,oneof=Groceries Leisure Electronics Utilities Clothing Health Others"`
	Amount   float64 `json:"amount" binding:"required"`
	Period   string  `json:"period" binding:"required,oneof=daily weekly monthly yearly"`
}

// {
//   "category": "Groceries",
//   "amount": 5000,
//   "period": "monthly"
// }
