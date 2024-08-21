package main

import (
	"log"
	"os"
	"passport_card_analyser/internal/adapters/app/api"
	"passport_card_analyser/internal/adapters/core/ocrscanner"
	"passport_card_analyser/internal/adapters/framework/left/httpadapter"
	"passport_card_analyser/internal/adapters/framework/right/db"
	"passport_card_analyser/internal/ports"

	"github.com/joho/godotenv"
)

func getTessearctTextExtractor() ports.OCRTextExtractor {
	return ocrscanner.TesseractTextExtractor{}
}

func main() {
	err := godotenv.Load()
	var (
		portString = os.Getenv("PORT")
		apier      ports.APIPort
		store      ports.DBPort
		httpdriver ports.HttpPort
		ocradapter ports.OCRScannerPost
	)
	store, err = db.NewAdapter()
	if err != nil {
		log.Fatal(err)
	}
	ocradapter = ocrscanner.NewAdapter(getTessearctTextExtractor())
	apier = api.NewAdapter(ocradapter, store)
	httpdriver = httpadapter.NewAdapter(apier)
	httpdriver.Run(portString)

	err = store.CloseDatabase()
	if err != nil {
		log.Fatal(err)
	}
}
