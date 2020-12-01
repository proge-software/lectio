package main

import (
	"log"

	bot "github.com/proge-software/lectio/internal/simplebot"
	"github.com/proge-software/lectio/pkg/envext"
)

// AppEnvKey Environment variable key where is stored the environment to use for the app
// execution. default is `development`
const AppEnvKey = "WSSBOT_ENV"

func main() {
	log.Println("Starting simple telegram bot")

	envext.LoadDotenvs(AppEnvKey) // load the env vars from .env file

	// alternative: https://github.com/kelseyhightower/envconfig
	fbot, err := bot.NewFromEnv()
	if err != nil {
		log.Fatalf("can not instantiate bot: %v", err)
	}

	log.Println("Simple Telegram bot configured")

	// start bot
	fbot.Start()
}
