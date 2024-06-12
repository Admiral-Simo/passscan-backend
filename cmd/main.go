package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"passport_card_analyser/internal/adapters/core/ocr"
	"passport_card_analyser/internal/adapters/core/types"
	"passport_card_analyser/internal/adapters/core/utilities"
)

func main() {
	flag.Parse()

	passports := flag.Args()

	if len(passports) == 0 {
		log.Fatalf("Usage: %s --passport example_image.jpeg", os.Args[0])
	}

	people := []*types.Person{}

	parser := ocr.NewParser()
	for _, card := range passports {
		parser.SetImage(card)
		person, err := parser.ParseCitizen()
		if err != nil {
			log.Println(err)
			continue
		}
		people = append(people, person)
	}

	for _, person := range people {
		fmt.Println("----------------------")
		utilities.PrintPerson(person)
	}
	fmt.Println("----------------------")
}
