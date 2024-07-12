package httpadapter

import (
	"net/http"
)

func (httpa Adapter) HandleGetUploadHistory(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}
