package models

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserResult struct {
	ID string 
	Email string
	Password string
}

func NewUserResult(_id interface{}, model UserModel) UserResult {
	objectId, _ := _id.(primitive.ObjectID)
	stringId := objectId.Hex()
	return UserResult{
		ID: stringId,
		Email: model.Email,
		Password: model.Password,
	}
}

func (v UserResult) Respond(c *gin.Context) {
	c.JSON(201, gin.H{
		"msg": "user created",
		"_id": v.ID,
		"email": v.Email,
	})
}