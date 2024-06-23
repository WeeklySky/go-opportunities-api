package main

import (
	"github.com/guilhermemcardoso/go-opportunities-api/config"
	"github.com/guilhermemcardoso/go-opportunities-api/router"
)

var (
	logger config.Logger
)

func main() {
	logger = *config.GetLogger("main")
	// initialize config
	err := config.Init()
	if err != nil {
		logger.Errorf("config init error: %v", err)
		return
	}

	// initialize router
	router.Initialize()
}
