package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Message struct {
	Id      uuid.UUID `json:"id" gorm:"column:id;primaryKey"`
	RoomId  uuid.UUID `json:"room_id" gorm:"column:room_id;not null"`
	UserId  uuid.UUID `json:"user_id" gorm:"column:user_id;not null"`
	Content string    `json:"content" gorm:"column:content"`
	SentAt  time.Time `json:"sent_at" gorm:"column:sent_at"`
}

func (m *Message) BeforeCreate(tx *gorm.DB) (err error) {
	m.Id = uuid.New()
	return
}
