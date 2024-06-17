package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"passport_card_analyser/internal/adapters/core/ocr"
	"passport_card_analyser/internal/adapters/core/types"
	"passport_card_analyser/internal/adapters/core/utilities"
	"time"
)

type Adapter struct {
	Port string
}

func NewAdapter(port string) *Adapter {
	return &Adapter{
		Port: port,
	}
}

func (adap *Adapter) getPassportData(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseMultipartForm(10 << 20) // 10 MB max
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving the file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	if !utilities.CheckImage(handler.Filename) {
		http.Error(w, "Error the file is not an image", http.StatusBadRequest)
		return
	}

	// extract extension

	outputFilePath := fmt.Sprintf("uploads/%d%s", time.Now().UnixNano(), utilities.ExtractExtension(handler.Filename))
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

	parser := ocr.NewParser()
	parser.SetImage(outputFilePath)
	person, names, err := parser.ParseCitizen()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	type responseType struct {
		Person               *types.Person `json:"person"`
		PossibleNamesAddress []string      `json:"possible_names_address"`
	}

	response := responseType{
		Person:               person,
		PossibleNamesAddress: names,
	}

	json.NewEncoder(w).Encode(response)
}

func (adap *Adapter) Run() error {
	http.HandleFunc("/get-passport-data", adap.getPassportData)

	fmt.Printf("listening to port %s\n", adap.Port)
	err := http.ListenAndServe(adap.Port, enableCors(http.DefaultServeMux))
	return err
}

func enableCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
