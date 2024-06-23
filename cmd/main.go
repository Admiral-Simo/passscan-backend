package main

import (
	"log"
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
	store, err := db.NewAdapter()
	if err != nil {
		log.Fatal(err)
	}
	ocradapter = ocrscanner.NewAdapter()
	apier = api.NewAdapter(ocradapter, store)
	httpdriver = httpadapter.NewAdapter(apier)
	httpdriver.Run(portString)

	err = store.CloseDatabase()
	if err != nil {
		log.Fatal(err)
	}
}
