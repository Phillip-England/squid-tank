package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SessionResult struct {
	SessionId string
	UserId string
	Expiration time.Time
}

func NewSessionResult(sessionId interface{}, sessionModel SessionModel) SessionResult {
	objectId := sessionId.(primitive.ObjectID)
	stringId := objectId.Hex()
	return SessionResult{
		SessionId: stringId,
		UserId: sessionModel.UserId,
		Expiration: sessionModel.Expiration,
	}
}


