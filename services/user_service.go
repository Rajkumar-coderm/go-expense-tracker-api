package services

import (
	"time"

	"github.com/expense-tracker-api/models"
	"github.com/expense-tracker-api/repository"
	"github.com/expense-tracker-api/utils"
)

func UserSignup(rq models.UserSignupReq) (models.LoginSignupRes, error) {
	password, pasErr := utils.HashPassword(rq.Password)
	if pasErr != nil {
		return models.LoginSignupRes{}, pasErr
	}

	payload := models.UserSignupReq{
		Email:     rq.Email,
		Name:      rq.Name,
		Password:  password,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	id, err := repository.UserSignup(payload)
	if err != nil {
		return models.LoginSignupRes{}, err
	}

	payload.ID = id

	token, tokenErr := utils.GenerateAccessToken(payload.ID.Hex(), payload.Email, "use")
	if tokenErr != nil {
		return models.LoginSignupRes{}, tokenErr
	}

	finalRes := models.LoginSignupRes{
		ID:        payload.ID,
		Name:      payload.Name,
		Email:     payload.Email,
		CreatedAt: payload.CreatedAt,
		Token:     token,
	}
	return finalRes, nil
}

func UserLogin(req models.LoginReq) (error, models.LoginSignupRes) {
	err, user := repository.UserLogin(req)
	if err != nil {
		return err, models.LoginSignupRes{}
	}

	token, tokenErr := utils.GenerateAccessToken(user.ID.Hex(), user.Email, "use")
	if tokenErr != nil {
		return tokenErr, models.LoginSignupRes{}
	}

	finalRes := models.LoginSignupRes{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		Token:     token,
	}
	return nil, finalRes
}

func GetUserById(id string) (models.UserProfile, error) {
	return repository.GetUserById(id)
}
