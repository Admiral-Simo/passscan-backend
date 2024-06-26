package db

import (
	"passport_card_analyser/types"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Adapter struct {
	db *gorm.DB
}

func NewAdapter() (*Adapter, error) {
	// Use the appropriate driver for your database (e.g., sqlite, postgres, mysql)
	// Here we're using sqlite for simplicity
	db, err := gorm.Open(sqlite.Open("passport_scanner.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto migrate the models
	err = db.AutoMigrate(&types.MRZData{})
	if err != nil {
		return nil, err
	}

	return &Adapter{
		db: db,
	}, nil
}

func (dba *Adapter) CloseDatabase() error {
	// Get the underlying sql.DB object and close it
	sqlDB, err := dba.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
