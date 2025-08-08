package main

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/Shaurya213/Restaurant-Management/internal/db"
	"github.com/Shaurya213/Restaurant-Management/internal/di"
)

func main() {
	_ = godotenv.Load()

	// 1) DB
	db.ConnectDB()

	// 2) DI
	app, err := di.InitializeApp()
	if err != nil {
		log.Fatalf("DI init failed: %v", err)
	}

	// Smoke log to confirm DI worked
	log.Printf("DI ready: AdminService=%T, MenuService=%T, OrderService=%T",
		app.AdminService, app.MenuService, app.OrderService)

	// gRPC server + handlers next...
}
