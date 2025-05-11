package main

import (
	"github.com/pierrialexandervidmar/go-jobs/config"
	"github.com/pierrialexandervidmar/go-jobs/router"
)

var (
	logger *config.Logger
)

func main() {

	logger = config.GetLogger("main")

	// Initialize config
	err := config.Init()

	if err != nil {
		logger.Errorf("Config initialization error: %v", err)
		return
	}

	// Initialize Router
	router.Initialize()

}
