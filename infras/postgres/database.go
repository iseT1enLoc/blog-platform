package database

import (
	"blog-platform-go/component/appconfig"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabaseWithRetryIn20s(cfg *appconfig.AppConfig) (*gorm.DB, error) {
	const timeRetry = 20 * time.Second
	var connectDatabase = func(cfg *appconfig.AppConfig) (*gorm.DB, error) {
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
