package routes

import (
	"cfa-tools-api/src/app"

	"github.com/gin-gonic/gin"
)

func LocationRoutes(r *gin.Engine, db *app.Database) {
	
	// creating a user
	r.POST("/location", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "creating a user..",
		})
	})

	// getting user locations
	r.GET("/locations", func(c *gin.Context) {
		


	})

}