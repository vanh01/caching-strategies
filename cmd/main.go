package main

import (
	"log"

	"github.com/vanh01/caching-strategies/config"
	"github.com/vanh01/caching-strategies/internal/app"
)

func main() {
	// load config from file
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config :%s", err.Error())
	}
	config.Instance = cfg

	// run application
	app.Run()
}
