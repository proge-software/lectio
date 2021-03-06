package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	bot "github.com/proge-software/lectio/internal/moderatebot"
	"github.com/proge-software/lectio/internal/tgconf"
	"github.com/proge-software/lectio/pkg/envext"
	"github.com/proge-software/lectio/pkg/wssface"
	"github.com/proge-software/lectio/pkg/wssformrecognizer"
	"github.com/proge-software/lectio/pkg/wssmoderator"
	"github.com/proge-software/lectio/pkg/wsssentiment"
	"github.com/proge-software/lectio/pkg/wsstranslator"
	"github.com/proge-software/lectio/pkg/wssvision"
)

func main() {
	log.Println("Configuring Moderate bot")
	loadEnvs() // load the env vars from .env file

	conf, err := getConfigurations()
	if err != nil {
		log.Fatalln(err)
	}

	fbot, err := bot.New(*conf)
	if err != nil {
		log.Fatalf("can not instantiate bot: %v", err)
	}

	log.Println("Starting Moderate bot")
	// start bot
	fbot.Start()
}

func getConfigurations() (*bot.Configuration, error) {
	c := bot.Configuration{}

	{ // Telegram Configuration
		conf, err := tgconf.GetConfigurationsFromEnv()
		if err != nil {
			return nil, fmt.Errorf("error retrieving telegram configuration: %v", err)
		}
		c.TelegramConf = *conf
	}
	{ // Vision: Face
		faceConf, err := wssface.BuildConfigurationFromEnvs()
		if err != nil {
			log.Printf("error retrieving face service configuration: %v", err)
		}
		c.FaceConf = faceConf
	}
	{ // Vision: Computer Vision
		visionConf, err := wssvision.BuildConfigurationFromEnvs()
		if err != nil {
			log.Printf("error retrieving vision service configuration: %v", err)
		}
		c.VisionConf = visionConf
	}
	{ // Vision: Form Recognizer
		formRecognizerConf, err := wssformrecognizer.BuildConfigurationFromEnvs()
		if err != nil {
			log.Printf("error retrieving form recognizer service configuration: %v", err)
		}
		c.FormRecognizerConf = formRecognizerConf
	}
	{ // Language: Text Analytics
		textAnalyticsConf, err := wsssentiment.BuildConfigurationFromEnvs()
		if err != nil {
			log.Printf("error retrieving text analitycs service configuration: %v", err)
		}
		c.TextAnalyticsConf = textAnalyticsConf
	}
	{ // Language: Translator
		translatorConf, err := wsstranslator.BuildConfigurationFromEnvs()
		if err != nil {
			log.Printf("error retrieving translator service configuration: %v", err)
		}
		c.TranslatorConf = translatorConf
	}
	{ // Decision: Moderator
		moderatorConf, err := wssmoderator.BuildConfigurationFromEnvs()
		if err != nil {
			log.Printf("error retrieving moderator service configuration: %v", err)
		}
		c.ModeratorConf = moderatorConf
	}

	return &c, nil
}

// AppEnvKey Environment variable key where is stored the environment to use for the app
// execution. default is `development`
const AppEnvKey = "WSSBOT_ENV"

func loadEnvs() {
	env := envext.GetEnvOrDefault(AppEnvKey, "development")
	godotenv.Load(".env." + env + ".local")
	if "test" != env {
		godotenv.Load(".env.local")
	}

	envfile := ".env." + env
	if err := godotenv.Load(envfile); err != nil {
		cwd, _ := os.Getwd()
		log.Printf("error loading file %v (%s): %v", envfile, cwd, err)
	}

	if err := godotenv.Load(); err != nil { // The Original .env
		cwd, _ := os.Getwd()
		log.Printf("error loading .env (%v) file: %v", cwd, err)
	}
}
