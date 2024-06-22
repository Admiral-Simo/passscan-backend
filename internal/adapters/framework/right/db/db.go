package db

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"passport_card_analyser/types"
)

type Adapter struct {
	file *os.File
}

func NewAdapter(filename string) *Adapter {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		log.Fatal(err)
	}
	return &Adapter{
		file: file,
	}
}

func (dba Adapter) CreatePassport(person types.PersonWithNames) error {
	data, err := json.Marshal(person)
	if err != nil {
		return err
	}
	_, err = dba.file.Write(append(data, '\n'))
	return err
}

func (dba Adapter) GetPassports() ([]*types.PersonWithNames, error) {
	var persons []*types.PersonWithNames
	_, err := dba.file.Seek(0, 0)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(dba.file)
	for scanner.Scan() {
		var person types.PersonWithNames
		if err := json.Unmarshal(scanner.Bytes(), &person); err != nil {
			return nil, err
		}
		persons = append(persons, &person)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return persons, nil
}

func (dba Adapter) GetPassport(cne string) (*types.PersonWithNames, error) {
	persons, err := dba.GetPassports()
	if err != nil {
		return nil, err
	}
	for _, person := range persons {
		if person.Person.CIN == cne {
			return person, nil
		}
	}
	return nil, nil
}

func (dba Adapter) CloseDatabase() error {
	return dba.file.Close()
}
