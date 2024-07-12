package httpadapter

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"passport_card_analyser/internal/ports"
)

type Adapter struct {
	apia ports.APIPort
}

func NewAdapter(apia ports.APIPort) *Adapter {
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

	return &Adapter{
		apia: apia,
	}
}

func (httpa Adapter) Run(postString string) {
	http.HandleFunc("/get-document-data", httpa.HandleGetDocumentData)
	http.HandleFunc("/get-upload-history", httpa.HandleGetUploadHistory)
	http.HandleFunc("/uploads/", httpa.HandleGetImage)

	fmt.Printf("listening to port %s\n", postString)
	http.ListenAndServe(postString, enableCors(http.DefaultServeMux))
}
