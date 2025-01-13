package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type GroupQuestion struct {
	ID           bson.ObjectID `json:"_id" bson:"_id"`
	Title        *string       `json:"title"`
	Description  *string       `json:"description,omitempty"`
	ThumbnailUrl *string       `json:"thumbnailUrl,omitempty"`
	Questions    []Question    `json:"questions"`
	Answers      []Answer      `json:"answers"`
	CreatedBy    bson.ObjectID `json:"createdBy" bson:"createdBy"`
	AnyTime      bool          `json:"anyTime"`
	StartAt      *time.Time    `json:"startAt"`
	EndAt        *time.Time    `json:"endAt"`
	PublishedAt  *time.Time    `json:"publishedAt,omitempty"`
	CreatedAt    time.Time     `json:"createdAt"`
	UpdatedAt    time.Time     `json:"updatedAt"`
	DeletedAt    *time.Time    `json:"deletedAt,omitempty"`
}

type Answer struct {
	ID        bson.ObjectID `json:"_id" bson:"_id"`
	UserId    bson.ObjectID `json:"userId" bson:"_id"`
	Lock      bool          `json:"lock,omitempty"`
	Questions []Question    `json:"questions"`
	StartAt   time.Time     `json:"startAt"`
	EndAt     time.Time     `json:"endAt"`
	CreatedAt time.Time     `json:"createdAt"`
	UpdatedAt time.Time     `json:"updatedAt"`
	DeletedAt *time.Time    `json:"deletedAt,omitempty"`
}

type Question struct {
	ID            bson.ObjectID  `json:"_id" bson:"_id"`
	OriginNumber  int            `json:"originNumber"`
	Note          *string        `json:"note,omitempty"`
	Text          *string        `json:"text,omitempty"`
	ImageUrl      *string        `json:"imageUrl,omitempty"`
	VideoUrl      *string        `json:"videoUrl,omitempty"`
	YoutubeUrl    *string        `json:"youtubeUrl,omitempty"`
	QuestionItems []QuestionItem `json:"questionItems"`
	PublishedAt   *time.Time     `json:"publishedAt,omitempty"`
	CreatedAt     time.Time      `json:"createdAt"`
	CreatedBy     bson.ObjectID  `json:"createdBy" bson:"CreatedBy"`
	UpdatedAt     time.Time      `json:"updatedAt"`
	DeletedAt     *time.Time     `json:"deletedAt,omitempty"`
}

type QuestionItem struct {
	ID           bson.ObjectID `json:"_id" bson:"_id"`
	Key          *string       `json:"key,omitempty"`
	OriginNumber int           `json:"originNumber"`
	Text         *string       `json:"text,omitempty"`
	ImageUrl     *string       `json:"imageUrl,omitempty"`
	VideoUrl     *string       `json:"videoUrl,omitempty"`
	YoutubeUrl   *string       `json:"youtubeUrl,omitempty"`
}

type QuestionTemplate struct {
	ID                    bson.ObjectID          `json:"_id" bson:"_id"`
	OriginNumber          int                    `json:"originNumber"`
	Note                  *string                `json:"note,omitempty"`
	Text                  *string                `json:"text,omitempty"`
	ImageUrl              *string                `json:"imageUrl,omitempty"`
	VideoUrl              *string                `json:"videoUrl,omitempty"`
	YoutubeUrl            *string                `json:"youtubeUrl,omitempty"`
	QuestionTemplateItems []QuestionTemplateItem `json:"questionTemplateItems"`
	PublishedAt           *time.Time             `json:"publishedAt,omitempty"`
	CreatedAt             time.Time              `json:"createdAt"`
	UpdatedAt             time.Time              `json:"updatedAt"`
	DeletedAt             *time.Time             `json:"deletedAt,omitempty"`
}
type QuestionTemplateItem struct {
	ID           bson.ObjectID `json:"_id" bson:"_id"`
	Key          *string       `json:"key,omitempty"`
	OriginNumber int           `json:"originNumber"`
	Text         *string       `json:"text,omitempty"`
	ImageUrl     *string       `json:"imageUrl,omitempty"`
	VideoUrl     *string       `json:"videoUrl,omitempty"`
	YoutubeUrl   *string       `json:"youtubeUrl,omitempty"`
}

type QuestionItemInput struct {
	ID           bson.ObjectID `json:"_id" bson:"_id"`
	Key          *string       `json:"key,omitempty"`
	OriginNumber int           `json:"originNumber"`
	Text         *string       `json:"text,omitempty"`
	ImageUrl     *string       `json:"imageUrl,omitempty"`
	VideoUrl     *string       `json:"videoUrl,omitempty"`
	YoutubeUrl   *string       `json:"youtubeUrl,omitempty"`
}

type QuestionInput struct {
	ID                bson.ObjectID       `json:"_id" bson:"_id"`
	OriginNumber      int                 `json:"originNumber"`
	Note              *string             `json:"note,omitempty"`
	Text              *string             `json:"text,omitempty"`
	ImageUrl          *string             `json:"imageUrl,omitempty"`
	VideoUrl          *string             `json:"videoUrl,omitempty"`
	YoutubeUrl        *string             `json:"youtubeUrl,omitempty"`
	QuestionItemInput []QuestionItemInput `json:"questionItems"`
}

type GroupQuestionInput struct {
	Title        *string    `json:"title"`
	Description  *string    `json:"description,omitempty"`
	ThumbnailUrl *string    `json:"thumbnailUrl,omitempty"`
	AnyTime      bool       `json:"anyTime"`
	StartAt      *time.Time `json:"startAt,omitempty"`
	EndAt        *time.Time `json:"endAt,omitempty"`
	CreatedAt    time.Time  `json:"createdAt,omitempty"`
	UpdatedAt    time.Time  `json:"updatedAt,omitempty"`
}
