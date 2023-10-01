package main

import (
	"github.com/marcelomartinsdev/Golang-API/config"
	"github.com/marcelomartinsdev/Golang-API/router"
)

var (
	logger *config.Logger
)

func main() {

	logger = config.GetLogger("main")
	// Initialize Configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	// Initialize Router
	router.Initialize()
}
