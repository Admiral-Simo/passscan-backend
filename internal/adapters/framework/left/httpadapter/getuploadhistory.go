package httpadapter

import (
	"encoding/json"
	"net/http"
	"os"
)

func (httpa Adapter) HandleGetUploadHistory(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	const dirPath = "./uploads"
    urlPrefix := os.Getenv("IP_ADDRESS") + os.Getenv("PORT") + "/uploads/"

	images, err := getFilesByDate(dirPath, urlPrefix)
	if err != nil {
		http.Error(w, "no such file or directory with the name of "+dirPath, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(images)
}

func getFilesByDate(dirPath string, urlPrefix string) (map[string][]string, error) {
	// Initialize the map to hold files grouped by modification date
	filesByDate := make(map[string][]string)

	// Read the directory contents
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		// Skip directories
		if entry.IsDir() {
			continue
		}

		// Get the file's stat
		fileInfo, err := entry.Info()
		if err != nil {
			return nil, err
		}

		// Get the modification time of the file
		modTime := fileInfo.ModTime()
		dateStr := modTime.Format("2006-01-02") // Format date as "YYYY-MM-DD"

		// Add the file to the map under the corresponding date
		filesByDate[dateStr] = append(filesByDate[dateStr], urlPrefix+entry.Name())
	}

	return filesByDate, nil
}
