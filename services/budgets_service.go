package services

import (
	"time"

	"github.com/expense-tracker-api/models"
	"github.com/expense-tracker-api/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateNewBudget(req models.CreateBudgetRequestModel, userId string) error {
	objID, _ := primitive.ObjectIDFromHex(userId)
	re := models.BudgetModel{
		UserID:    objID,
		Category:  req.Category,
		Amount:    req.Amount,
		Period:    req.Period,
		StartDate: time.Now(),
		EndDate:   endTimeForPeriod(time.Now(), req.Period),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return repository.CreateNewBudget(re)
}

func endTimeForPeriod(startDate time.Time, period string) time.Time {
	startDate = startDate.UTC()

	switch period {
	case "daily":
		return startDate.AddDate(0, 0, 1).Add(-time.Nanosecond)

	case "weekly":
		return startDate.AddDate(0, 0, 7).Add(-time.Nanosecond)

	case "monthly":
		return startDate.AddDate(0, 1, 0).Add(-time.Nanosecond)

	case "yearly":
		return startDate.AddDate(1, 0, 0).Add(-time.Nanosecond)

	default:
		return startDate
	}
}
