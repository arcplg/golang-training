package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type User struct {
	ID        bson.ObjectID `json:"_id" bson:"_id"`
	Avatar    string        `json:"avatar"`
	Email     string        `json:"email"`
	Name      string        `json:"name"`
	Phone     string        `json:"phone"`
	CreatedAt time.Time     `json:"createdAt"`
	UpdatedAt time.Time     `json:"updatedAt"`
	DeletedAt *time.Time    `json:"deletedAt,omitempty"`
}
