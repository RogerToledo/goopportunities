package main

import (
	"github.com/me/goopportunities/config"
	"github.com/me/goopportunities/router"
)

var logger *config.Logger

func main() {
	logger = config.GetLogger("main")
	
	err := config.Init()
	if err != nil {
		logger.Errorf("config error: %v", err)
		return
	}

	router.Initialize()
}
