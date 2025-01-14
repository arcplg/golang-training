package services

import (
	"context"
	"fmt"
	"learning-server/entity"
	"learning-server/internal/db"
	"learning-server/internal/validation"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func GetGroupQuestions(ctx context.Context) ([]*entity.GroupQuestion, error) {
	collection := db.GetCollection("questions")
	filter := bson.D{}
	opts := options.Find().SetSort(bson.D{{Key: "createdAt", Value: 1}})

	cursor, err := collection.Find(ctx, filter, opts)

	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var groupQuestions []*entity.GroupQuestion
	for cursor.Next(ctx) {
		var question entity.GroupQuestion
		err := cursor.Decode(&question)
		if err != nil {
			return nil, err
		}
		groupQuestions = append(groupQuestions, &question)
	}

	return groupQuestions, nil
}

func FindGroupQuestion(ctx context.Context, id string) (*entity.GroupQuestion, error) {
	collection := db.GetCollection("questions")
	var groupQuestion entity.GroupQuestion
	_id, _ := bson.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}
	err := collection.FindOne(ctx, filter).Decode(&groupQuestion)
	if err != nil {
		return nil, err
	}

	return &groupQuestion, nil
}

func CreateGroupQuestion(ctx context.Context, input entity.GroupQuestionInput) (*entity.GroupQuestion, error) {
	if err := validation.ValidateStruct(input); err != nil {
		return nil, fmt.Errorf("validation failed: %v", err)
	}
	groupQuestionInput := &entity.GroupQuestionInput{
		Title:        input.Title,
		Description:  input.Description,
		ThumbnailUrl: input.ThumbnailUrl,
		AnyTime:      input.AnyTime,
		StartAt:      input.StartAt,
		EndAt:        input.EndAt,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	collection := db.GetCollection("questions")

	res, err := collection.InsertOne(ctx, groupQuestionInput)
	if err != nil {
		return nil, err
	}

	var groupQuestion entity.GroupQuestion
	filter := bson.M{"_id": res.InsertedID}
	err = collection.FindOne(ctx, filter).Decode(&groupQuestion)
	if err != nil {
		return nil, err
	}

	return &groupQuestion, nil
}

func GetQuestions(ctx context.Context) ([]*entity.Question, error) {
	collection := db.GetCollection("questions")
	cursor, err := collection.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var questions []*entity.Question
	for cursor.Next(ctx) {
		var question entity.Question
		err := cursor.Decode(&question)
		if err != nil {
			return nil, err
		}
		questions = append(questions, &question)
	}

	return questions, nil
}

func CreateQuestion(ctx context.Context, input entity.QuestionInput) (*entity.Question, error) {
	if err := validation.ValidateStruct(input); err != nil {
		return nil, fmt.Errorf("validation failed: %v", err)
	}

	questionItemsList := make([]entity.QuestionItemInput, len(input.QuestionItemInput))
	for i, input := range input.QuestionItemInput {
		questionItemsList[i] = entity.QuestionItemInput{
			ID:           bson.NewObjectID(),
			Key:          input.Key,
			OriginNumber: input.OriginNumber,
			Text:         input.Text,
			ImageUrl:     input.ImageUrl,
			VideoUrl:     input.VideoUrl,
			YoutubeUrl:   input.YoutubeUrl,
		}
	}

	questionInput := &entity.QuestionInput{
		ID:                bson.NewObjectID(),
		OriginNumber:      input.OriginNumber,
		Text:              input.Text,
		ImageUrl:          input.ImageUrl,
		VideoUrl:          input.VideoUrl,
		YoutubeUrl:        input.YoutubeUrl,
		QuestionItemInput: questionItemsList,
	}

	collection := db.GetCollection("questions")
	res, err := collection.InsertOne(ctx, questionInput)
	if err != nil {
		return nil, err
	}

	var question entity.Question
	filter := bson.M{"_id": res.InsertedID}
	err = collection.FindOne(ctx, filter).Decode(&question)
	if err != nil {
		return nil, err
	}

	return &question, nil
}
