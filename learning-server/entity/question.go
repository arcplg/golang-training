package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

/** Entity */
type Media struct {
	ID   bson.ObjectID `bson:"_id"`
	Type *string       `bson:"type,omitempty"`
	Url  *string       `bson:"url,omitempty"`
}

type Block struct {
	ID     bson.ObjectID `bson:"_id"`
	Label  string        `bson:"label,omitempty"`
	Text   string        `bson:"text,omitempty"`
	Media  Media         `bson:"media,omitempty"`
	Blocks []Block       `bson:"block"`
}

type Exam struct {
	ID           bson.ObjectID `bson:"_id"`
	Title        *string       `bson:"title"`
	Description  *string       `bson:"description,omitempty"`
	ThumbnailUrl *string       `bson:"thumbnailUrl,omitempty"`
	Questions    []Question    `bson:"questions"`
	Answers      []Answer      `bson:"answers"`
	AnyTime      bool          `bson:"anyTime"`
	StartAt      *time.Time    `bson:"startAt"`
	EndAt        *time.Time    `bson:"endAt"`
	PublishedAt  *time.Time    `bson:"publishedAt,omitempty"`
	CreatedAt    time.Time     `bson:"createdAt"`
	CreatedBy    bson.ObjectID `bson:"createdBy,omitempty"`
	UpdatedAt    time.Time     `bson:"updatedAt"`
	UpdatedBy    bson.ObjectID `bson:"UpdatedBy,omitempty"`
	DeletedAt    *time.Time    `bson:"deletedAt,omitempty"`
	DeletedBy    bson.ObjectID `bson:"deletedBy,omitempty"`
}

type ExamInput struct {
	Title        *string    `json:"title,omitempty"`
	Description  *string    `json:"description,omitempty"`
	ThumbnailUrl *string    `json:"thumbnailUrl,omitempty"`
	AnyTime      *bool      `json:"anyTime"`
	StartAt      *time.Time `json:"startAt,omitempty"`
	EndAt        *time.Time `json:"endAt,omitempty"`
}

type Answer struct {
	ID        bson.ObjectID `bson:"_id"`
	UserId    bson.ObjectID `bson:"userId"`
	Lock      bool          `bson:"lock,omitempty"`
	Questions []Question    `bson:"questions"`
	StartAt   time.Time     `bson:"startAt"`
	EndAt     time.Time     `bson:"endAt"`
}

type Question struct {
	ID            bson.ObjectID `bson:"_id,omitempty"`
	Name          *string       `bson:"name,omitempty"`
	Note          *string       `bson:"note,omitempty"`
	Text          *string       `bson:"text,omitempty"`
	Media         *Media        `bson:"media,omitempty"`
	Options       []Block       `bson:"options"`
	CorrectOption []Block       `bson:"correctOption"`
	PublishedAt   *time.Time    `bson:"publishedAt,omitempty"`
	CreatedAt     time.Time     `bson:"createdAt"`
	CreatedBy     bson.ObjectID `bson:"createdBy,omitempty"`
	UpdatedAt     time.Time     `bson:"updatedAt"`
	UpdatedBy     bson.ObjectID `bson:"UpdatedBy,omitempty"`
	DeletedAt     *time.Time    `bson:"deletedAt,omitempty"`
	DeletedBy     bson.ObjectID `bson:"DeletedBy,omitempty"`
}
type QuestionInput struct {
	Name          *string `bson:"name,omitempty"`
	Note          *string `bson:"note,omitempty"`
	Text          *string `bson:"text,omitempty"`
	Media         *Media  `bson:"media,omitempty"`
	Options       []Block `bson:"options"`
	CorrectOption []Block `bson:"correctOption"`
}

type QuestionTemplate struct {
	ID     bson.ObjectID `bson:"_id"`
	Name   *string       `bson:"name,omitempty"`
	Note   *string       `bson:"note,omitempty"`
	Text   *string       `bson:"text,omitempty"`
	Media  *Media        `bson:"media,omitempty"`
	Blocks []Block       `bson:"blocks"`
}
