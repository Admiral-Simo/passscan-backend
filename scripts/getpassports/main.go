package main

import (
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

	UpdateTemplate(store)

	bounds, err := GetTemplate(store, "MA")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("bounds:", bounds)
}

func CreateTemplate(store ports.DBPort) {
	var rectangles []types.Rectangle
	for i := 0; i < 7; i++ {
		rectangle := types.Rectangle{
			TopLeft: types.Point{
				X: 10,
				Y: 10,
			},
			TopRight: types.Point{
				X: 20,
				Y: 10,
			},
			BottomLeft: types.Point{
				X: 10,
				Y: 20,
			},
			BottomRight: types.Point{
				X: 20,
				Y: 20,
			},
		}
		rectangles = append(rectangles, rectangle)
	}
	template := types.OCRTemplate{
		Nationality: "MA",
		Bounds:      rectangles,
	}
	err := store.CreateTemplate(template)
	log.Fatal(err)
}

func UpdateTemplate(store ports.DBPort) {
	var rectangles []types.Rectangle
	for i := 0; i < 7; i++ {
		rectangle := types.Rectangle{
			TopLeft: types.Point{
				X: 10,
				Y: 10,
			},
			TopRight: types.Point{
				X: 20,
				Y: 10,
			},
			BottomLeft: types.Point{
				X: 10,
				Y: 20,
			},
			BottomRight: types.Point{
				X: 20,
				Y: 20,
			},
		}
		rectangles = append(rectangles, rectangle)
	}
	template := types.OCRTemplate{
		ID:          1,
		Nationality: "MA",
		Bounds:      rectangles,
	}
	store.UpdateTemplate(template)
}

func GetTemplate(store ports.DBPort, nationatlity string) ([]types.Rectangle, error) {
	template, err := store.GetTemplateByNationality(nationatlity)
	return template.Bounds, err
}

func GetTemplates(store ports.DBPort, nationatlity string) {
	templates, err := store.GetTemplates()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(templates)
}
