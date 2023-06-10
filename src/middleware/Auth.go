package middleware

import (
	"cfa-tools-api/src/app"
	"cfa-tools-api/src/e"
	"cfa-tools-api/src/models"
	"time"

	"github.com/gin-gonic/gin"
)

func Auth(db *app.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		// checking if the user provided the authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || len(authHeader) < 8 {
			httpErr := e.NewHttpError("unauthorized", 401)
			httpErr.Exit(c)
			c.Abort()
			return
		}
		// extracting the session token out of the header
		sessionToken := c.GetHeader("Authorization")[7:]
		// using the session token to load the session
		sessionModel := models.NewSessionModel("")
		sessiondb := models.NewSessionDb(db.SessionCollection, sessionModel)
		sessionResult, err := sessiondb.FindById(sessionToken)
		if err != nil {
			err.Exit(c)
			c.Abort()
			return
		}
		// checking if our session is expired
		if sessionResult.Expiration.Before(time.Now()) {
			e.NewHttpError("session expired", 401).Exit(c)
			c.Abort()
			return
		}
		// creating a user db from our user ID
		userdb := models.NewUserDb(db.UserCollection, nil)
		// getting our userResult from our userdb
		userResult, err := userdb.FindById(sessionResult.User)
		if err != nil {
			e.NewHttpError("unauthorized", 401).Exit(c)
			c.Abort()
			return
		}
		c.Set("user", userResult)
		c.Next()
	}
}