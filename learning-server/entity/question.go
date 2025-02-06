package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

/** Entity */
type Media struct {
	ID   bson.ObjectID `bson:"_id"`
	Type string        `bson:"type"`
	Url  string        `bson:"url"`
}

type Block struct {
	ID     bson.ObjectID `bson:"_id"`
	Label  *string       `bson:"label"`
	Text   *string       `bson:"text"`
	Media  *Media        `bson:"media"`
	Blocks []Block       `bson:"blocks"`
}

type Exam struct {
	ID           bson.ObjectID  `bson:"_id"`
	Title        string         `bson:"title"`
	Description  *string        `bson:"description"`
	ThumbnailUrl *string        `bson:"thumbnailUrl"`
	Questions    []Question     `bson:"questions"`
	Answers      []Answer       `bson:"answers"`
	AnyTime      bool           `bson:"anyTime"`
	StartAt      *time.Time     `bson:"startAt"`
	EndAt        *time.Time     `bson:"endAt"`
	PublishedAt  *time.Time     `bson:"publishedAt"`
	CreatedAt    time.Time      `bson:"createdAt"`
	CreatedBy    *bson.ObjectID `bson:"createdBy"`
	UpdatedAt    time.Time      `bson:"updatedAt"`
	UpdatedBy    *bson.ObjectID `bson:"UpdatedBy"`
	DeletedAt    *time.Time     `bson:"deletedAt"`
	DeletedBy    *bson.ObjectID `bson:"deletedBy"`
}

type ExamInput struct {
	Title        string     `json:"title" validate:"required,min=5,max=100"`
	Description  *string    `json:"description"`
	ThumbnailUrl *string    `json:"thumbnailUrl"`
	AnyTime      bool       `json:"anyTime"`
	StartAt      *time.Time `json:"startAt"`
	EndAt        *time.Time `json:"endAt"`
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
	ID            bson.ObjectID  `bson:"_id"`
	Code          string         `bson:"code"`
	Blocks        []Block        `bson:"blocks"`
	Options       []Block        `bson:"options"`
	CorrectOption []Block        `bson:"correctOption"`
	CreatedAt     time.Time      `bson:"createdAt"`
	CreatedBy     *bson.ObjectID `bson:"createdBy"`
	UpdatedAt     time.Time      `bson:"updatedAt"`
	UpdatedBy     *bson.ObjectID `bson:"UpdatedBy"`
	DeletedAt     *time.Time     `bson:"deletedAt"`
	DeletedBy     *bson.ObjectID `bson:"DeletedBy"`
}
type QuestionInput struct {
	ID            string       `bson:"_id"`
	Code          string       `bson:"code"`
	Blocks        []BlockInput `bson:"blocks"`
	Options       []BlockInput `bson:"options"`
	CorrectOption []BlockInput `bson:"correctOption"`
}

type BlockInput struct {
	ID     *string      `bson:"_id,omitempty"`
	Label  *string      `bson:"label,omitempty"`
	Text   *string      `bson:"text,omitempty"`
	Media  *MediaInput  `bson:"media,omitempty"`
	Blocks []BlockInput `bson:"blocks"`
}

type MediaInput struct {
	ID   *string `bson:"_id,omitempty"`
	Type *string `bson:"type,omitempty"`
	URL  *string `bson:"url,omitempty"`
}

type QuestionTemplate struct {
	ID      bson.ObjectID `bson:"_id"`
	Code    string        `bson:"code"`
	Blocks  []Block       `bson:"blocks"`
	Options []Block       `bson:"options"`
}
