package main

import "passport_card_analyser/internal/adapters/app/api"

func main() {
    api := api.NewAdapter(":8080")
	api.Run()
}
