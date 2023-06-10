package models

import (
	"cfa-tools-api/src/e"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserDb struct {
	Collection *mongo.Collection
	UserModel *UserModel
}

func NewUserDb(coll *mongo.Collection, model *UserModel) UserDb {
	return UserDb{
		Collection: coll,
		UserModel: model,
	}
}

func (v UserDb) Insert() (*UserResult, *e.HttpError) {
	err := v.AssertUnique()
	if err != nil {
		return nil, err
	}
	result, _ := v.Collection.InsertOne(context.Background(), v.UserModel)
	userResult := NewUserResult(result.InsertedID, v.UserModel.Email, v.UserModel.Password)
	return &userResult, nil
}

func (v UserDb) AssertUnique() *e.HttpError {
	var userExists UserResult
	err := v.Collection.FindOne(context.Background(), bson.D{{Key: "email", Value: v.UserModel.Email}}).Decode(&userExists)
	if err == nil && err != mongo.ErrNoDocuments {
		return e.NewHttpError("user already exists", 400)
	}
	return nil
}

func (v UserDb) FindByEmail() (*UserResult, *e.HttpError) {
	var result UserResult
	filter := bson.D{{Key: "email", Value: v.UserModel.Email}}
	err := v.Collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, e.NewHttpError("user does not exist", 400)
		}
		return nil, e.NewHttpError("internal server error", 500)
	}
	return &result, nil
}

func (v UserDb) FindById(userid primitive.ObjectID) (*UserResult, *e.HttpError) {
	var result UserResult
	filter := bson.D{{Key: "_id", Value: userid}}
	err := v.Collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, e.NewHttpError("user does not exist", 400)
		}
		return nil, e.NewHttpError("internal server error", 400)
	}
	return &result, nil
}