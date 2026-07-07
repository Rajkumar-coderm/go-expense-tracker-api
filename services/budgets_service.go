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
		return time.Date(
			startDate.Year(),
			startDate.Month(),
			startDate.Day(),
			23, 59, 59, int(time.Second-time.Nanosecond),
			time.UTC,
		)

	case "weekly":
		daysUntilSunday := (7 - int(startDate.Weekday())) % 7
		end := startDate.AddDate(0, 0, daysUntilSunday)

		return time.Date(
			end.Year(),
			end.Month(),
			end.Day(),
			23, 59, 59, int(time.Second-time.Nanosecond),
			time.UTC,
		)

	case "monthly":
		firstDayNextMonth := time.Date(
			startDate.Year(),
			startDate.Month()+1,
			1,
			0, 0, 0, 0,
			time.UTC,
		)
		return firstDayNextMonth.Add(-time.Nanosecond)

	case "yearly":
		return time.Date(
			startDate.Year(),
			time.December,
			31,
			23, 59, 59, int(time.Second-time.Nanosecond),
			time.UTC,
		)

	default:
		return startDate
	}
}

func GetBudget(req models.GetBudgetsRequestModel) ([]models.BudgetModel, error) {
	return repository.GetBudget(req)
}
