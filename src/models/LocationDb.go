package models

import (
	"cfa-tools-api/src/e"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type LocationDb struct {
	Collection *mongo.Collection
	LocationModel LocationModel
}

func NewLocationDb(coll *mongo.Collection, model LocationModel) LocationDb {
	return LocationDb{
		Collection: coll,
		LocationModel: model,
	}
}

func (v LocationDb) Insert() (*LocationResult, *e.HttpError) {
	userObjectId, _ := primitive.ObjectIDFromHex(v.LocationModel.User)
	result, err := v.Collection.InsertOne(context.Background(), bson.D{
		{Key: "user", Value: userObjectId},
		{Key: "name", Value: v.LocationModel.Name},
		{Key: "number", Value: v.LocationModel.Number},
	})
	fmt.Println(result)
	// return nil, nil
	if err != nil {
		return nil, e.NewHttpError("internal server error", 500)
	}
	locationResult := NewLocationResult(result.InsertedID, v.LocationModel)
	return &locationResult, nil
}