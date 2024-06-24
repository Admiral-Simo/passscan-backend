package ports

import "net/http"

type HttpPort interface {
	HandleGetPassportData(w http.ResponseWriter, r *http.Request)
	HandleGetTemplateNationalities(w http.ResponseWriter, r *http.Request)
	Run(portString string)
}
