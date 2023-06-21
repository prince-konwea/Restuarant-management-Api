package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Table struct {
	ID               primitive.ObjectID `json:"_id"`
	Number_of_guests int64              `json:"number_of_guests" validate:"required"`
	Table_number     int64              `json:"table_number"    validate:"required"`
	Created_at       time.Time          `json:"created_at"`
	Updated_at       time.Time          `json:"updated_at"`
	Table_id         string             `json:"table_id"`
}
