package main

import (
	"log"

	"github.com/yusuftatli/hepsiburada/config"
	"github.com/yusuftatli/hepsiburada/jobs"
)

func main() {
	//init environment
	env := config.GetEnvironment()

	//get config
	cfg, err := config.NewDefaultConfig()
	if err != nil {
		log.Panicln("failed to load config")
	}

	jobs.GetCurrencyJob(cfg)

	InitializeRoutes(env)
}
