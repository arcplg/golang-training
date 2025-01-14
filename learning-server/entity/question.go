package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type GroupQuestion struct {
	ID           bson.ObjectID `bson:"_id"`
	Title        *string       `bson:"title"`
	Description  *string       `bson:"description,omitempty"`
	ThumbnailUrl *string       `bson:"thumbnailUrl,omitempty"`
	Questions    []Question    `bson:"questions"`
	Answers      []Answer      `bson:"answers"`
	CreatedBy    bson.ObjectID `bson:"createdBy"`
	AnyTime      bool          `bson:"anyTime"`
	StartAt      *time.Time    `bson:"startAt"`
	EndAt        *time.Time    `bson:"endAt"`
	PublishedAt  *time.Time    `bson:"publishedAt,omitempty"`
	CreatedAt    time.Time     `bson:"createdAt"`
	UpdatedAt    time.Time     `bson:"updatedAt"`
	DeletedAt    *time.Time    `bson:"deletedAt,omitempty"`
}

type Answer struct {
	ID        bson.ObjectID `bson:"_id"`
	UserId    bson.ObjectID `bson:"userId"`
	Lock      bool          `bson:"lock,omitempty"`
	Questions []Question    `bson:"questions"`
	StartAt   time.Time     `bson:"startAt"`
	EndAt     time.Time     `bson:"endAt"`
	CreatedAt time.Time     `bson:"createdAt"`
	UpdatedAt time.Time     `bson:"updatedAt"`
	DeletedAt *time.Time    `bson:"deletedAt,omitempty"`
}

type Question struct {
	ID            bson.ObjectID  `bson:"_id"`
	OriginNumber  int            `bson:"originNumber"`
	Note          *string        `bson:"note,omitempty"`
	Text          *string        `bson:"text,omitempty"`
	ImageUrl      *string        `bson:"imageUrl,omitempty"`
	VideoUrl      *string        `bson:"videoUrl,omitempty"`
	YoutubeUrl    *string        `bson:"youtubeUrl,omitempty"`
	QuestionItems []QuestionItem `bson:"questionItems"`
	PublishedAt   *time.Time     `bson:"publishedAt,omitempty"`
	CreatedAt     time.Time      `bson:"createdAt"`
	CreatedBy     bson.ObjectID  `bson:"createdBy"`
	UpdatedAt     time.Time      `bson:"updatedAt"`
	DeletedAt     *time.Time     `bson:"deletedAt,omitempty"`
}

type QuestionItem struct {
	ID           bson.ObjectID `bson:"_id"`
	Key          *string       `bson:"key,omitempty"`
	OriginNumber int           `bson:"originNumber"`
	Text         *string       `bson:"text,omitempty"`
	ImageUrl     *string       `bson:"imageUrl,omitempty"`
	VideoUrl     *string       `bson:"videoUrl,omitempty"`
	YoutubeUrl   *string       `bson:"youtubeUrl,omitempty"`
}

type QuestionTemplate struct {
	ID                    bson.ObjectID          `bson:"_id"`
	OriginNumber          int                    `bson:"originNumber"`
	Note                  *string                `bson:"note,omitempty"`
	Text                  *string                `bson:"text,omitempty"`
	ImageUrl              *string                `bson:"imageUrl,omitempty"`
	VideoUrl              *string                `bson:"videoUrl,omitempty"`
	YoutubeUrl            *string                `bson:"youtubeUrl,omitempty"`
	QuestionTemplateItems []QuestionTemplateItem `bson:"questionTemplateItems"`
	PublishedAt           *time.Time             `bson:"publishedAt,omitempty"`
	CreatedAt             time.Time              `bson:"createdAt"`
	UpdatedAt             time.Time              `bson:"updatedAt"`
	DeletedAt             *time.Time             `bson:"deletedAt,omitempty"`
}
type QuestionTemplateItem struct {
	ID           bson.ObjectID `bson:"_id"`
	Key          *string       `bson:"key,omitempty"`
	OriginNumber int           `bson:"originNumber"`
	Text         *string       `bson:"text,omitempty"`
	ImageUrl     *string       `bson:"imageUrl,omitempty"`
	VideoUrl     *string       `bson:"videoUrl,omitempty"`
	YoutubeUrl   *string       `bson:"youtubeUrl,omitempty"`
}

type QuestionItemInput struct {
	ID           bson.ObjectID `bson:"_id"`
	Key          *string       `bson:"key,omitempty"`
	OriginNumber int           `bson:"originNumber"`
	Text         *string       `bson:"text,omitempty"`
	ImageUrl     *string       `bson:"imageUrl,omitempty"`
	VideoUrl     *string       `bson:"videoUrl,omitempty"`
	YoutubeUrl   *string       `bson:"youtubeUrl,omitempty"`
}

type QuestionInput struct {
	ID                bson.ObjectID       `bson:"_id"`
	OriginNumber      int                 `bson:"originNumber"`
	Note              *string             `bson:"note,omitempty"`
	Text              *string             `bson:"text,omitempty"`
	ImageUrl          *string             `bson:"imageUrl,omitempty"`
	VideoUrl          *string             `bson:"videoUrl,omitempty"`
	YoutubeUrl        *string             `bson:"youtubeUrl,omitempty"`
	QuestionItemInput []QuestionItemInput `bson:"questionItems"`
}

type GroupQuestionInput struct {
	Title        *string    `bson:"title"`
	Description  *string    `bson:"description,omitempty"`
	ThumbnailUrl *string    `bson:"thumbnailUrl,omitempty"`
	AnyTime      bool       `bson:"anyTime"`
	StartAt      *time.Time `bson:"startAt,omitempty"`
	EndAt        *time.Time `bson:"endAt,omitempty"`
	CreatedAt    time.Time  `bson:"createdAt,omitempty"`
	UpdatedAt    time.Time  `bson:"updatedAt,omitempty"`
}
