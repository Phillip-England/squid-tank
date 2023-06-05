package routes

import (
	"cfa-tools-api/src/app"

	"github.com/gin-gonic/gin"
)

func Mount(r *gin.Engine, db *app.Database) {
	UserRoutes(r, db)
	LocationRoutes(r, db)
}