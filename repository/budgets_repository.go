package repository

import (
	"context"
	"time"

	"github.com/expense-tracker-api/config"
	"github.com/expense-tracker-api/models"
)

func CreateNewBudget(req models.BudgetModel) error {

	col := config.GetCollection("budgets")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := col.InsertOne(ctx, req)

	if err != nil {
		return err
	}
	return nil
}

func GetBudget(userId string) ([]models.BudgetModel, error) {
	col := config.GetCollection("budgets")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	results, err := col.Find(ctx, map[string]interface{}{"user_id": userId})
	if err != nil {
		return nil, err
	}
	var budgets []models.BudgetModel
	if err = results.All(ctx, &budgets); err != nil {
		return nil, err
	}
	return budgets, nil
}
