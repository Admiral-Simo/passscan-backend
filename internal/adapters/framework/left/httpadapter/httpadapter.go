package httpadapter

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"passport_card_analyser/internal/ports"
	"time"
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

func (httpa Adapter) HandleGetPassportData(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseMultipartForm(10 << 20) // 10 MB max
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	nationality := r.FormValue("nationality")

	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving the file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	if !checkImage(handler.Filename) {
		http.Error(w, "Error the file is not an image", http.StatusBadRequest)
		return
	}

	// extract extension

	outputFilePath := fmt.Sprintf("uploads/%d%s", time.Now().UnixNano(), extractExtension(handler.Filename))
	dst, err := os.Create(outputFilePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// parse citizen

	person, err := httpa.apia.GetPassportData(outputFilePath, nationality)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	fmt.Println("response:", person)
	fmt.Println("nationality:", nationality)

	person.Nationality = nationality

	json.NewEncoder(w).Encode(person)
}

func (httpa Adapter) Run(postString string) {
	http.HandleFunc("/get-passport-data", httpa.HandleGetPassportData)

	fmt.Printf("listening to port %s\n", postString)
	http.ListenAndServe(postString, enableCors(http.DefaultServeMux))
}