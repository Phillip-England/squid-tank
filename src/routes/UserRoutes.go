package routes

import (
	"cfa-tools-api/src/app"
	"cfa-tools-api/src/e"
	"cfa-tools-api/src/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func UserRoutes(r *gin.Engine, db *app.Database) {

	// creating a user
	r.POST("/user", func(c *gin.Context) {
		var userModel models.UserModel
		_ = c.BindJSON(&userModel)
		userModel.Format()
		err := userModel.Validate()
		if err != nil {
			err.HttpExit(c)
			return
		}
		err = userModel.HashPassword()
		if err != nil {
			err.HttpExit(c)
			return
		}
		userdb := models.NewUserDb(db.UserCollection, userModel)
		result, err := userdb.Insert()
		if err != nil {
			err.HttpExit(c)
			return
		}
		result.Respond(c)
	})

	// logging a user in
	r.POST("/login", func(c *gin.Context) {
		var userModel models.UserModel
		_ = c.BindJSON(&userModel)
		userModel.Format()
		userdb := models.NewUserDb(db.UserCollection, userModel)
		userResult, err := userdb.FindByEmail()
		if err != nil {
			errOveride := e.NewHttpError("invalid credentials", 400)
			errOveride.HttpExit(c)
			return
		}
		newErr := bcrypt.CompareHashAndPassword([]byte(userResult.Password), []byte(userModel.Password))
		if newErr != nil {
			err := e.NewHttpError("invalid credentials", 400)
			err.HttpExit(c)
			return
		}
		sessionModel := models.NewSessionModel(userResult.ID)
		sessiondb := models.NewSessionDb(db.SessionCollection, sessionModel)
		sessiondb.DeleteAll()
		sessionResult, err := sessiondb.Insert()
		if err != nil {
			err.HttpExit(c)
			return
		}
		c.SetCookie("session-token", sessionResult.SessionId, 86400, "/", "localhost", true, true)
		c.JSON(200, gin.H{
			"msg": "user logged in",
		})
	})

	// logging a user out
	r.GET("/logout", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "logging out",
		})
	})


}