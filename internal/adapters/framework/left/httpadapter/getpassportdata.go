package httpadapter

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func (httpa Adapter) HandleGetDocumentData(w http.ResponseWriter, r *http.Request) {
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

	if !checkImage(handler.Filename) {
		http.Error(w, "Error the file is not an image", http.StatusBadRequest)
		return
	}

	// extract extension

	names := strings.Split(handler.Filename, ".")
	name := names[0] + "-"
	extension := "." + names[1]

	unixTiming := time.Now().UnixNano()

	outputFilePath := fmt.Sprintf("uploads/%s%d%s", name, unixTiming, extension)
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

	// parse document

	person, err := httpa.apia.GetDocumentData(outputFilePath)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(person)
}
