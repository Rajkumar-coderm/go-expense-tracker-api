package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type LoginSignupRes struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name" json:"name" binding:"required"`
	Email     string             `bson:"email" json:"email" binding:"required,email"`
	CreatedAt int64              `bson:"createdAt" json:"created_at"`
	Token     string             `bson:"token" json:"token"`
}
