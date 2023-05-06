package config

import (
	"os"

	"github.com/LucasRufo/golang-first-api/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeDb() (*gorm.DB, error) {
	logger := GetLogger()
	dbPath := "./db/opportunity.db"

	_, err := os.Stat(dbPath)

	if os.IsNotExist(err) {
		logger.Info("creating database and directory")

		err = os.MkdirAll("./db", os.ModePerm)

		if err != nil {
			return nil, err
		}

		file, err := os.Create(dbPath)

		if err != nil {
			return nil, err
		}

		file.Close()
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

	if err != nil {
		logger.Errorf("failed to connect database: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Opportunity{})

	if err != nil {
		logger.Errorf("failed to execute AutoMigrate: %v", err)
		return nil, err
	}

	return db, nil
}
