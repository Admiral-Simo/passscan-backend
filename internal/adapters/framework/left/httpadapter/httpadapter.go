package httpadapter

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"passport_card_analyser/internal/ports"
	"passport_card_analyser/types"
	"strings"
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

	person.Nationality = nationality

	json.NewEncoder(w).Encode(getResponseFromPerson(*person))
}

func (httpa Adapter) HandleGetTemplateNationalities(w http.ResponseWriter, r *http.Request) {
	nationalities, err := httpa.apia.GetTempateNationalities()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	type Response struct {
		Nationalities []string `json:"nationalities"`
	}

	response := Response{
		Nationalities: nationalities,
	}

	json.NewEncoder(w).Encode(response)
}

func (httpa Adapter) Run(postString string) {
	http.HandleFunc("/get-passport-data", httpa.HandleGetPassportData)
	http.HandleFunc("/get-nationalities", httpa.HandleGetTemplateNationalities)

	fmt.Printf("listening to port %s\n", postString)
	http.ListenAndServe(postString, enableCors(http.DefaultServeMux))
}

type response struct {
	CNIE                 string    `json:"cin"`
	FirstName            string    `json:"first_name"`
	LastName             string    `json:"last_name"`
	City                 string    `json:"city"`
	Nationality          string    `json:"nationality"`
	BirthDate            time.Time `json:"birth_date"`
	ExpireDate           time.Time `json:"expire_date"`
	PossibleNamesAddress []string  `json:"possible_names_address"`
}

func getResponseFromPerson(person types.Person) response {
	var resp response
	resp.CNIE = person.CNIE
	resp.FirstName = person.FirstName
	resp.LastName = person.LastName
	resp.City = person.City
	resp.Nationality = person.Nationality
	resp.BirthDate = person.BirthDate
	resp.ExpireDate = person.ExpireDate
	for _, name := range person.PossibleNamesAddress {
		resp.PossibleNamesAddress = append(resp.PossibleNamesAddress, name.Name)
	}

	return resp
}
