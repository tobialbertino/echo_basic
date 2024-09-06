package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func InitConfig() (res *AppConfig, err error) {
	err = godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	res = &AppConfig{
		GeminiAppKey: os.Getenv("GEMINI_API_KEY"),
	}

	return
}
