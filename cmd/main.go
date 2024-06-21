package main

import (
	"passport_card_analyser/internal/adapters/app/api"
	"passport_card_analyser/internal/adapters/core/ocrscanner"
	"passport_card_analyser/internal/adapters/framework/left/httpadapter"
	"passport_card_analyser/internal/ports"
)

func main() {
	var (
		portString = ":8080"
		apier      ports.APIPort
		httpdriver ports.HttpPort
		ocradapter ports.OCRScannerPost
	)
	ocradapter = ocrscanner.NewAdapter()
	apier = api.NewAdapter(ocradapter)
	httpdriver = httpadapter.NewAdapter(apier)
	httpdriver.Run(portString)
}
