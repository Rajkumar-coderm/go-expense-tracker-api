package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserSignupReq struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name" json:"name" binding:"required"`
	Email     string             `bson:"email" json:"email" binding:"required,email"`
	Password  string             `bson:"password" json:"password" binding:"required"`
	CreatedAt int64              `bson:"createdAt" json:"created_at"`
	UpdatedAt int64              `bson:"updatedAt" json:"updated_at"`
}
