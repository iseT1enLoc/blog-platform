package appconfig

import (
	"fmt"
	"log"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
func ConnectDatabaseWithRetryIn20s(cfg *AppConfig) (*gorm.DB, error) {
	const timeRetry = 20 * time.Second
	var connectDatabase = func(cfg *AppConfig) (*gorm.DB, error) {
		psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
			"password=%s dbname=%s sslmode=disable",
			cfg.DBHost, cfg.Port, cfg.DBUsername, cfg.DBPassword, cfg.DBDatabase)
		db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
		if err != nil {
			return nil, fmt.Errorf("failed to connect to database: %w", err)
		}
		return db.Debug(), nil
	}

	var db *gorm.DB
	var err error

	deadline := time.Now().Add(timeRetry)

	for time.Now().Before(deadline) {
		log.Println("Connecting to database...")
		db, err = connectDatabase(cfg)
		if err == nil {
			return db, nil
		}
		time.Sleep(time.Second)
	}

	return nil, fmt.Errorf("failed to connect to database after retrying for 20 seconds: %w", err)
}
