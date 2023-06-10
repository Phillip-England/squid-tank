package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SessionResult struct {
	Id string `bson:"_id"`
	User primitive.ObjectID `bson:"user"`
	Expiration time.Time `bson:"expiration"`
}

func NewSessionResult(sessionId interface{}, sessionModel SessionModel) SessionResult {
	objectId := sessionId.(primitive.ObjectID)
	stringId := objectId.Hex()
	return SessionResult{
		Id: stringId,
		User: sessionModel.User,
		Expiration: sessionModel.Expiration,
	}
}


