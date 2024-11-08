package main

import (
	"api-go/api"
	"log"
	"os"
)

func main() {
	pocketbaseURL := os.Getenv("POCKETBASE_URL")
	router := api.SetupRoutes(pocketbaseURL)

	log.Println("Starting server on :3000")
	if err := router.Run(":3000"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
