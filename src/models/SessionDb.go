package models

import (
	"cfa-tools-api/src/e"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type SessionDb struct {
	Collection *mongo.Collection
	SessionModel SessionModel
}

func NewSessionDb(collection *mongo.Collection, sessionModel SessionModel) SessionDb {
	return SessionDb{
		Collection: collection,
		SessionModel: sessionModel,
	}
}

func (v SessionDb) Insert() (*SessionResult, *e.HttpError) {
	result, err := v.Collection.InsertOne(context.Background(), v.SessionModel)
	if err != nil {
		return nil, e.NewHttpError("internal server error", 500)
	}
	sessionResult := NewSessionResult(result.InsertedID, v.SessionModel)
	return &sessionResult, nil
}

func (v SessionDb) DeleteAll() {
	v.Collection.DeleteMany(context.Background(), bson.D{})
}

func (v SessionDb) FindById(sessionToken string) (*SessionResult, *e.HttpError) {
	var result SessionResult
	sessionObjectId, err := primitive.ObjectIDFromHex(sessionToken)
	if err != nil {
		return nil, e.NewHttpError("unauthorized", 400)
	}
	filter := bson.D{{Key: "_id", Value: sessionObjectId}}
	err = v.Collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, e.NewHttpError("unauthorized", 400)
		}
		return nil, e.NewHttpError("internal server error", 500)
	}
	return &result, nil
}