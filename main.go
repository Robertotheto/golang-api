package main

import (
	"github.com/Robertotheto/golang-api/config"
	"github.com/Robertotheto/golang-api/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	// Initialize Config
	err := config.Init()
	if err != nil {
		logger.Errorf("Error initializing config: %v", err)
		return
	}
	// Initialize Router
	router.Initialize()
}
