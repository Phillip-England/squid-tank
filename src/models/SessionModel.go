package models

import "time"

type SessionModel struct {
	UserId string
	Expiration time.Time
}

func NewSessionModel(userId string) SessionModel {
	return SessionModel{
		UserId: userId,
		Expiration: time.Now().Add(24 * time.Hour),
	}
}