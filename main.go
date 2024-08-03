package main

import (
	"log"
	"os"

	db "github.com/DevloperAmanSingh/app-backend/database"
	"github.com/DevloperAmanSingh/app-backend/router"
	"github.com/joho/godotenv"
)

func main() {
	app := router.SetupRouter()
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	dbURL := os.Getenv("MONGO_URI")
	dbName := os.Getenv("MONGO_DATABASE_NAME")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8001"
	}
	db.ConnectDatabase(dbURL, dbName)
	defer db.DisconnectDatabase()

	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

}
