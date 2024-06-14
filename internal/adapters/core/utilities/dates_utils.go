package utilities

import (
	"errors"
	"passport_card_analyser/internal/adapters/core/types"
	"sort"
	"time"
)

var ErrNoDate = errors.New("no date found")

const (
	slashLayout = "02/01/2006"
	dotLayout   = "02.01.2006"
	spaceLayout = "02 01 2006"
)

func ParseDate(word string) (time.Time, error) {
	var layouts = []string{slashLayout, dotLayout, spaceLayout}
	var parsers = []func(string) bool{
		ContainsTwoSlashes,
		ContainsTwoDots,
		AllDigits,
	}

	for i, parseFunc := range parsers {
		if parseFunc(word) && len(word) == 10 {
			dateTime, err := time.Parse(layouts[i], word[:10])
			if err != nil {
				return time.Time{}, err
			}
			return dateTime, nil
		}
	}

	return time.Time{}, ErrNoDate
}

func SortDates(dates []time.Time) {
	sort.Slice(dates, func(i, j int) bool {
		return dates[i].Before(dates[j])
	})
}

func AssignDates(person *types.Person, dates []time.Time) {
	if len(dates) > 0 {
		person.BirthDate = dates[0]
	}
	if len(dates) > 1 {
		person.ExpireDate = dates[1]
	}
}
