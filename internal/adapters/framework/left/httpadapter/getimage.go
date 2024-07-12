package httpadapter

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func (httpa Adapter) HandleGetImage(w http.ResponseWriter, r *http.Request) {
	filename := strings.TrimPrefix(r.URL.Path, "/uploads/")

	if filename == "" {
		http.Error(w, "Filename is required", http.StatusBadRequest)
		return
	}

	filePath := filepath.Join("uploads", filename)

	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, "File is not found", http.StatusNotFound)
		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()

	if err != nil {
		http.Error(w, "Error retreiving file information", http.StatusInternalServerError)
		return
	}

	contentType := "application/octet-stream"
	ext := filepath.Ext(filename)
	switch ext {
	case ".jpg", ".jpeg":
		contentType = "image/jpeg"
	case ".png":
		contentType = "image/png"
	case ".gif":
		contentType = "image/gif"
	}

	w.Header().Set("Content-Type", contentType)
	w.Header().Set("Content-Length", string(fileInfo.Size()))

	http.ServeFile(w, r, filePath)
}
