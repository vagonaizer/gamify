package main

import (
	"fmt"
	"gamify/internal/app"
	"gamify/internal/config"
	"log"
)

func main() {
	fmt.Println("📝 test log")

	// Initializing the config
	cfg, err := config.LoadConfig("./")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	log.Print("config loaded")

	application, err := app.NewApp(cfg)
	if err != nil {
		log.Fatalf("failed to initialize app: %v", err)
	}

	if err := application.Run(); err != nil {
		log.Fatalf("failed to run app: %v", err)
	}
}
