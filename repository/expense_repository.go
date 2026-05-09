package repository

import (
	"context"
	"errors"
	"time"

	"github.com/expense-tracker-api/config"
	"github.com/expense-tracker-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateNewExpense(req models.ExpenseModel) error {
	col := config.GetCollection("expenses")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := col.InsertOne(ctx, req)

	if err != nil {
		return err
	}
	return nil
}

func GetExpense(req models.ExpenseGetQueryModel, isGetAll bool) ([]models.ExpenseModel, error) {
	col := config.GetCollection("expenses")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_userId, _ := primitive.ObjectIDFromHex(req.UserId)

	filter := bson.M{"user_id": _userId}

	now := time.Now()
	var startDate time.Time

	if req.Type != "" {
		filter["type"] = req.Type
	}

	switch req.DateFilter.Filter {
	case "week":
		startDate = now.AddDate(0, 0, -7)
		filter["expense_date"] = bson.M{"$gte": startDate}
	case "month":
		startDate = now.AddDate(0, -1, 0)
		filter["expense_date"] = bson.M{"$gte": startDate}
	case "3months":
		startDate = now.AddDate(0, -3, 0)
		filter["expense_date"] = bson.M{"$gte": startDate}
	case "custom":
		start, _ := time.Parse("2006-01-02", req.DateFilter.StartDate)
		end, _ := time.Parse("2006-01-02", req.DateFilter.EndDate)
		filter["expense_date"] = bson.M{
			"$gte": start,
			"$lte": end,
		}
	}

	findOptions := options.Find()
	if isGetAll {
		findOptions.SetSort(bson.D{{Key: "expense_date", Value: -1}}) // Newest first
	} else {
		findOptions.SetLimit(int64(req.Limit))
		findOptions.SetSkip(int64((req.Page - 1) * req.Limit))
		findOptions.SetSort(bson.D{{Key: "expense_date", Value: -1}}) // Newest first
	}

	cursor, err := col.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var expenses []models.ExpenseModel
	if err = cursor.All(ctx, &expenses); err != nil {
		return nil, err
	}

	return expenses, nil
}

func DeleteExpense(id, userId string) error {
	col := config.GetCollection("expenses")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	userIdObj, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return err
	}

	idObj, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": idObj}
	result := col.FindOne(ctx, filter)
	if err := result.Err(); err != nil {
		return err
	}

	var expense bson.M
	if err := result.Decode(&expense); err != nil {
		return err
	}

	storedUserID, ok := expense["user_id"].(primitive.ObjectID)
	if !ok || storedUserID != userIdObj {
		return errors.New("NotAccessToResouce")
	}

	_, err = col.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}

func ImportExpenses(expenses []models.ExpenseModel) error {
	col := config.GetCollection("expenses")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var docs []interface{}
	for _, e := range expenses {
		docs = append(docs, e)
	}
	_, err := col.InsertMany(ctx, docs)
	return err
}
