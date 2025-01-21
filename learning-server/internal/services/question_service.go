package services

import (
	"context"
	"fmt"
	"learning-server/entity"
	"learning-server/internal/db"
	"learning-server/internal/validation"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func ListExam(ctx context.Context) ([]*entity.Exam, error) {
	collection := db.GetCollection("questions")
	filter := bson.D{}
	opts := options.Find().SetSort(bson.D{{Key: "createdAt", Value: 1}})

	cursor, err := collection.Find(ctx, filter, opts)

	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var exams []*entity.Exam
	for cursor.Next(ctx) {
		var question entity.Exam
		err := cursor.Decode(&question)
		if err != nil {
			return nil, err
		}
		exams = append(exams, &question)
	}

	return exams, nil
}

func FindExamById(ctx context.Context, id string) (*entity.Exam, error) {
	collection := db.GetCollection("questions")
	var groupQuestion entity.Exam
	_id, _ := bson.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}
	err := collection.FindOne(ctx, filter).Decode(&groupQuestion)
	if err != nil {
		return nil, err
	}

	return &groupQuestion, nil
}

func CreateExam(ctx context.Context, input entity.ExamInput) (*entity.Exam, error) {
	if err := validation.ValidateStruct(input); err != nil {
		return nil, fmt.Errorf("validation failed: %v", err)
	}

	examInput := &entity.Exam{
		ID:           bson.NewObjectID(),
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

	res, err := collection.InsertOne(ctx, examInput)
	if err != nil {
		return nil, err
	}

	var exam entity.Exam
	filter := bson.M{"_id": res.InsertedID}
	err = collection.FindOne(ctx, filter).Decode(&exam)
	if err != nil {
		return nil, err
	}

	return &exam, nil
}

func ListQuestions(ctx context.Context) ([]*entity.Question, error) {
	collection := db.GetCollection("questions")

	pipeline := mongo.Pipeline{
		{{"$unwind", bson.D{{"path", "$questions"}}}},
		{{"$replaceRoot", bson.D{{"newRoot", "$questions"}}}},
	}

	cursor, err := collection.Aggregate(ctx, pipeline)

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

func ListQuestionTemplates(ctx context.Context) ([]*entity.QuestionTemplate, error) {
	collection := db.GetCollection("questionTemplates")
	filter := bson.D{}
	opts := options.Find()

	cursor, err := collection.Find(ctx, filter, opts)

	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var questionTemplate []*entity.QuestionTemplate
	for cursor.Next(ctx) {
		var question entity.QuestionTemplate
		err := cursor.Decode(&question)
		if err != nil {
			return nil, err
		}
		questionTemplate = append(questionTemplate, &question)
	}

	return questionTemplate, nil
}

func AddQuestionIntoExam(ctx context.Context, examId string, input entity.QuestionInput) (*entity.Exam, error) {
	if err := validation.ValidateStruct(input); err != nil {
		return nil, fmt.Errorf("validation failed: %v", err)
	}

	collection := db.GetCollection("questions")
	var exam entity.Exam
	_id, e := bson.ObjectIDFromHex(examId)
	if e != nil {
		return nil, e
	}

	filter := bson.M{"_id": _id}

	question := &entity.Question{
		ID:     bson.NewObjectID(),
		Name:   input.Name,
		Note:   input.Note,
		Text:   input.Text,
		Media:  input.Media,
		Blocks: input.Blocks,
	}

	update := bson.M{
		"$push": bson.M{
			"questions": question,
		},
	}
	collection.FindOneAndUpdate(ctx, filter, update)

	err := collection.FindOne(ctx, filter).Decode(&exam)
	if err != nil {
		return nil, err
	}

	return &exam, nil
}
