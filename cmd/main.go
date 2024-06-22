package main

import (
	"passport_card_analyser/internal/adapters/app/api"
	"passport_card_analyser/internal/adapters/core/ocrscanner"
	"passport_card_analyser/internal/adapters/framework/left/httpadapter"
	"passport_card_analyser/internal/adapters/framework/right/db"
	"passport_card_analyser/internal/ports"
)

func main() {
	var (
		portString = ":8080"
		apier      ports.APIPort
		store      ports.DBPort
		httpdriver ports.HttpPort
		ocradapter ports.OCRScannerPost
	)
	store = db.NewAdapter("database.txt")
	ocradapter = ocrscanner.NewAdapter()
	apier = api.NewAdapter(ocradapter, store)
	httpdriver = httpadapter.NewAdapter(apier)
	httpdriver.Run(portString)

	store.CloseDatabase()
}
