package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SessionModel struct {
	User primitive.ObjectID `bson:"user"`
	Expiration time.Time `bson:"expiration"`
}

func NewSessionModel(user string) SessionModel {
	objectid, _ := primitive.ObjectIDFromHex(user)
	return SessionModel{
		User: objectid,
		Expiration: time.Now().Add(24 * time.Hour),
	}
}
