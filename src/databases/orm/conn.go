package orm

import (
	"errors"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New() (*gorm.DB, error) {
	DB_HOST := os.Getenv("DB_HOST")
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_NAME := os.Getenv("DB_NAME")

	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s", DB_HOST, DB_USER, DB_PASSWORD, DB_NAME)

	gormDb, err := gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		return nil, errors.New("gagal init database")
	}

	db, err := gormDb.DB()
	if err != nil {
		return nil, errors.New("gagal init database")
	}

	db.SetConnMaxIdleTime(10)
	db.SetMaxOpenConns(100)

	return gormDb, nil
}
