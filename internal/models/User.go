package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id        uuid.UUID  `json:"id" gorm:"column:id;primaryKey"`
	Username  string     `json:"username" gorm:"column:username;not null"`
	Password  string     `json:"password" gorm:"column:password;not null"`
	CreatedAt time.Time  `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"column:deleted_at"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Id = uuid.New()
	return
}
