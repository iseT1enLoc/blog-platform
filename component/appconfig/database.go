package appconfig

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadConfig() (*AppConfig, error) {
	env, err := godotenv.Read()

	if err != nil {
		log.Fatalf("There is some error while loading .env %s", err)
	}
	return &AppConfig{
		Port:       env["PORT"],
		Env:        env["REMOTE_CONNECTION_STRING"],
		DBUsername: env["USERNAME"],
		DBPassword: env["PASSWORD"],
		DBHost:     env["HOST"],
		DBDatabase: env["DATABASE"],
		SecretKey:  env["SECRETKEY"],
	}, err
}
