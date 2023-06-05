package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LocationResult struct {
	ID primitive.ObjectID
	Name string
	Number string
}