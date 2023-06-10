package routes

import (
	"cfa-tools-api/src/app"
	"cfa-tools-api/src/middleware"
	"cfa-tools-api/src/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

func LocationRoutes(r *gin.Engine, db *app.Database) {
	
	// creating a location
	r.POST("/location", middleware.Auth(db), func(c *gin.Context) {
		var locationModel models.LocationModel
		_ = c.BindJSON(&locationModel)
		locationModel.Format()
		err := locationModel.Validate()
		if err != nil {
			err.Exit(c)
			return
		}
		locationDb := models.NewLocationDb(db.LocationCollection, locationModel)
		fmt.Println(locationDb.LocationModel)
		locationResult, err := locationDb.Insert()
		if err != nil {
			err.Exit(c)
			return
		}
		locationResult.Respond(c)
	})

	// getting user locations
	r.GET("/locations", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "success",
		})


	})

}