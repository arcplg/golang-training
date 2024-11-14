package models

import (
	"github.com/google/uuid"
	"time"
)

type RoomMember struct {
	UserId   uuid.UUID `json:"user_id" gorm:"column:user_id;not null"`
	RoomId   uuid.UUID `json:"room_id" gorm:"column:room_id;not null"`
	JoinedAt time.Time `json:"joined_at" gorm:"column:joined_at;not null"`
}
