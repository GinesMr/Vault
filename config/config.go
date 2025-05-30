package config

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadEnvConfig() error {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
		return err
	}
	return nil
}
