package repository

import (
	"context"
	"errors"
	"time"

	"github.com/expense-tracker-api/config"
	"github.com/expense-tracker-api/models"
	"github.com/expense-tracker-api/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UserLogin(req models.LoginReq) (error, models.UserSignupReq) {
	col := config.GetCollection("user")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	var user models.UserSignupReq

	result := col.FindOne(ctx, bson.M{"email": req.Email}).Decode(&user)

	if result != nil {
		return errors.New("User No Exists with this email please try again"), models.UserSignupReq{}
	}

	if utils.CheckPassword(user.Password, req.Password) != nil {
		return errors.New("Invalid password"), models.UserSignupReq{}
	}

	return nil, user

}

func UserSignup(req models.UserSignupReq) (primitive.ObjectID, error) {
	col := config.GetCollection("user")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result := col.FindOne(ctx, bson.M{"email": req.Email})

	if result.Err() == nil {
		return primitive.NilObjectID, errors.New("User Already Exists with same email please try with different email")
	}

	res, err := col.InsertOne(ctx, req)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return res.InsertedID.(primitive.ObjectID), err
}

func GetUserById(id string) (models.UserProfile, error) {
	col := config.GetCollection("user")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var user models.UserProfile
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.UserProfile{}, err
	}
	result := col.FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
	return user, result
}
