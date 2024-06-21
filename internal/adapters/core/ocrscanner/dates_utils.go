package ocrscanner

import (
	"errors"
	"passport_card_analyser/types"
	"sort"
	"time"
)

var errNoDate = errors.New("no date found")

func sortDates(dates []time.Time) {
	sort.Slice(dates, func(i, j int) bool {
		return dates[i].Before(dates[j])
	})
}

func assignDates(person *types.Person, dates []time.Time) {
	if len(dates) > 0 {
		person.BirthDate = dates[0]
	}
	if len(dates) > 1 {
		person.ExpireDate = dates[1]
	}
}
