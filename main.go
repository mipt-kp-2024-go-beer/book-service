package main

import (
	"context"
	"log"

	"github.com/mipt-kp-2024-go-beer/book-service/internal/app"
)

func main() {
	config, err := app.NewConfig("configs/config.yml")
	if err != nil {
		log.Fatalf("Failed to open config: %s", err)
	}

	// Create a new app instance
	ctx := context.Background()
	appInstance, err := app.New(ctx, config)
	if err != nil {
		log.Fatalf("Failed to initialize app: %v", err)
	}

	// Setup the application (initialize DB, services, and routes)
	if err := appInstance.Setup(ctx); err != nil {
		log.Fatalf("Failed to setup app: %v", err)
	}

	// Run the app (start the HTTP server)
	if err := appInstance.Start(); err != nil {
		log.Fatalf("Failed to start app: %v", err)
	}

	log.Println("Application has started successfully")
}
