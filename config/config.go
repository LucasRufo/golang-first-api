package config

import "gorm.io/gorm"

var (
	db     *gorm.DB
	logger *Logger
)

func Initialize() error {
	var err error

	db, err = InitializeDb()

	if err != nil {
		return err
	}

	return nil
}

func GetDb() *gorm.DB {
	return db
}

func GetLogger() *Logger {
	logger = NewLogger()
	return logger
}
