package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Room struct {
	Id        uuid.UUID  `json:"id" gorm:"column:id;primaryKey"`
	Name      string     `json:"name" gorm:"column:name"`
	CreateBy  uuid.UUID  `json:"create_by" gorm:"column:created_by"`
	CreatedAt time.Time  `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"column:deleted_at"`
}

func (r *Room) BeforeCreate(tx *gorm.DB) (err error) {
	r.Id = uuid.New()
	return
}
