package main

import (
	"encoding/json"
	"fmt"
	"log"
	"passport_card_analyser/internal/adapters/framework/right/db"
	"passport_card_analyser/internal/ports"
	"passport_card_analyser/types"
)

func main() {
	var store ports.DBPort
	store, err := db.NewAdapter()
	if err != nil {
		log.Fatal(err)
	}
	response, err := store.GetPassports()
	if err != nil {
		log.Fatal(err)
	}
	b, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
	CreateTemplate(store)
	err = store.CloseDatabase()
	if err != nil {
		log.Fatal(err)
	}
}

func CreateTemplate(store ports.DBPort) {
	var rectangles []types.Rectangle
	for i := 0; i < 7; i++ {
		rectangle := types.Rectangle{
			TopLeft:     10,
			TopRight:    10,
			BottomLeft:  10,
			BottomRight: 10,
		}
		rectangles = append(rectangles, rectangle)
	}
	template := types.OCRTemplate{
		Nationality: "MA",
		Bounds:      rectangles,
	}
	store.CreateTemplate(template)
}
