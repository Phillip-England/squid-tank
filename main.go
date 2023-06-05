package main

import (
	"cfa-tools-api/src/app"
	"cfa-tools-api/src/routes"
	"context"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// loading our env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env file")
	}

	// loading in our database
	db, err := app.NewDatabase()
	if err != nil {
		log.Fatal("failed to connect to mongo db")
	}
	defer db.Client.Disconnect(context.Background())

	// getting our router
	r := gin.Default()

	// adding a cors layer
	r.Use(cors.New(cors.Config{
    AllowOrigins: []string{"http://localhost:3000"},
    AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
    AllowHeaders: []string{"Origin", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}))

	// mounting our routes
	routes.Mount(r, db)

	// running our server
	r.Run(":8080")

}