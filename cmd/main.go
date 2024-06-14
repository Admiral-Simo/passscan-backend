package main

import (
	"log"
	"os"
	"passport_card_analyser/internal/adapters/app/api"
)

func main() {
	// Define the directory path
	uploadDir := "uploads"

	// Check if the directory exists
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		// Create the directory if it doesn't exist
		err := os.Mkdir(uploadDir, 0755) // 0755 is the permission for the directory
		if err != nil {
			log.Fatalf("Failed to create directory: %v", err)
		}
	}

	api := api.NewAdapter(":8080")
	api.Run()
}
