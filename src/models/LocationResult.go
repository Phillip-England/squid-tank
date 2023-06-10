package models

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LocationResult struct {
	Id string `bson:"_id"`
	User string `bson:"user"`
	Name string `bson:"name"`
	Number string `bson:"number"`
}

func NewLocationResult(_id interface{}, model LocationModel) LocationResult {
	objectId, _ := _id.(primitive.ObjectID)
	stringId := objectId.Hex()
	return LocationResult{
		Id: stringId,
		User: model.User,
		Name: model.Name,
		Number: model.Number,
	}
}

func (v LocationResult) Respond(c *gin.Context) {
	c.JSON(201, gin.H{
		"msg": "location created",
		"_id": v.Id,
		"name": v.Name,
		"number": v.Number,
	})
}