package repository

import (
	"context"
	"time"

	"github.com/expense-tracker-api/config"
	"github.com/expense-tracker-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func GetBudget(req models.GetBudgetsRequestModel) ([]models.BudgetModel, error) {
	col := config.GetCollection("budgets")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	userObjID, err := primitive.ObjectIDFromHex(req.UserId)
	var filter interface{}
	if err == nil {
		filter = bson.M{"user_id": userObjID}
	} else {
		filter = bson.M{"user_id": req.UserId}
	}

	if req.Category != "" {
		filter.(bson.M)["category"] = req.Category
	}

	if req.Page > 0 && req.Limit > 0 {
		skip := (req.Page - 1) * req.Limit
		opts := options.Find().SetSkip(int64(skip)).SetLimit(int64(req.Limit))
		cursor, err := col.Find(ctx, filter, opts)
		if err != nil {
			return nil, err
		}
		defer cursor.Close(ctx)

		var budgets []models.BudgetModel
		if err = cursor.All(ctx, &budgets); err != nil {
			return nil, err
		}
		return budgets, nil
	}

	cursor, err := col.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var budgets []models.BudgetModel
	if err = cursor.All(ctx, &budgets); err != nil {
		return nil, err
	}
	return budgets, nil
}
