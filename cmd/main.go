package main

import (
	"log"

	"github.com/Shaurya213/Restaurant-Management/internal/db"
	"github.com/joho/godotenv"
)

func main() {
	// Optional: Load .env if you're using it
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found. Using system env variables...")
	}

	// Init DB connection and auto-migrate
	db.ConnectDB()

	log.Println("Restaurant Management backend started")
}
