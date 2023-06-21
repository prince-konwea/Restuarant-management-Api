package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `bson:"_id"`
	First_name    *string            `json:"first_name" validate:"required min=2,max=50"`
	Last_name     *string            `json:"last_name"  validate:"required min=2,max=50"`
	Password      *string            `json:"password" validate:"required min=6"`
	Email         *string            `json:"email" validate:"required"`
	Avatar        *string            `json:"avatar""`
	Phone         *string            `json:"phone" validate:"required"`
	Token         *string            `json:"token" validate:"required"`
	Refresh_token *string            `json:"refresh_token" validate:"required"`
	Created_at    *time.Time         `json:"created_at" validate:"required"`
	Updated_at    *time.Time         `json:"updated_at" validate:"required"`
	User_id       *string            `json:"user_id" validate:"required"`
}
