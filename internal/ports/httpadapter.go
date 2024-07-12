package ports

import "net/http"

type HttpPort interface {
	HandleGetDocumentData(w http.ResponseWriter, r *http.Request)
	HandleGetUploadHistory(w http.ResponseWriter, r *http.Request)
	HandleGetImage(w http.ResponseWriter, r *http.Request)
	Run(portString string)
}
