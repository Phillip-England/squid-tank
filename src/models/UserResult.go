package models

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserResult struct {
	Id string `bson:"_id"`
	Email string `bson:"email"`
	Password string `bson:"password"`
}

func NewUserResult(_id interface{}, email string, password string) UserResult {
	objectId, _ := _id.(primitive.ObjectID)
	stringId := objectId.Hex()
	return UserResult{
		Id: stringId,
		Email: email,
		Password: password,
	}
}

func (v UserResult) Respond(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"msg": message,
		"_id": v.Id,
		"email": v.Email,
	})
}