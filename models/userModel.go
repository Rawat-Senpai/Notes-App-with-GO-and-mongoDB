package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `bson:"_id"`
	Name          *string            `json:"name" validate:"required,min=2,max=20"`
	Passowrd      *string            `json:"password" validate:"email,required"`
	Email         *string            `json:"phone" validate:"required"`
	Token         *string            `json:"token"`
	Refresh_token *string            `json:"refreshToken"`
	User_id       string             `json:"user_id"`
	Created_at    time.Time          `json:"created_at"`
	Updated_at    time.Time          `json:"updated_at"`
	Phone         *string            `json:"phone" validate:"required"`
}