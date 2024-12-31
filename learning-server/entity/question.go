package entity

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Question struct {
	ID          bson.ObjectID `json:"_id" bson:"_id"`
	Title       string        `json:"title"`
	Description *string       `json:"description,omitempty"`
}

type NewQuestion struct {
	Title       string  `json:"title" validate:"required,min=5,max=1000"`
	Description *string `json:"description,omitempty"`
}
