package entity

import (
	"github.com/99designs/gqlgen/graphql"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type File struct {
	ID   bson.ObjectID `json:"_id" bson:"_id"`
	Name string        `json:"name"`
	Type string        `json:"content"`
	Path string        `json:"path"`
	Url  string        `json:"url"`
}

type UploadFile struct {
	File graphql.Upload `json:"file"`
}
