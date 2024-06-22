package main

import (
	"encoding/json"
	"fmt"
	"log"
	"passport_card_analyser/internal/adapters/framework/right/db"
	"passport_card_analyser/internal/ports"
)

func main() {
	var store ports.DBPort
	store = db.NewAdapter("database.txt")
	response, err := store.GetPassports()
	if err != nil {
		log.Fatal(err)
	}
	b, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
	err = store.CloseDatabase()
	if err != nil {
		log.Fatal(err)
	}
}
