package config

import "gorm.io/gorm"

var (
	db     *gorm.DB
	logger *Logger
)

func Initialize() error {
	return nil
}

func GetLogger() *Logger {
	logger = NewLogger()
	return logger
}
