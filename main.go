package main

import (
	"fmt"
	"log"
	"passport_card_analyser/internal/adapters/core/ocr"
	"passport_card_analyser/internal/adapters/core/utilities"
)

func main() {
	parser := ocr.NewParser()
	cards := []string{"Bensalem-Alouakili-.jpg"}
	for i, card := range cards {
		parser.SetImage(card)
		person, err := parser.ParseCitizen()
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Println("----------------------")
		fmt.Printf("card number %d: ", i+1)
		utilities.PrintPerson(person)
	}
	fmt.Println("----------------------")
}
