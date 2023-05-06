package main

import (
	"github.com/LucasRufo/golang-first-api/config"
	"github.com/LucasRufo/golang-first-api/router"
)

var (
	logger *config.Logger
)

func main() {

	logger = config.GetLogger()

	err := config.Initialize()

	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	router.Initialize()
}
