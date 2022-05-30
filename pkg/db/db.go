package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"test/internal/models"
	"time"
)

func Connect() *gorm.DB {
	db, err := gorm.Open(postgres.Open("postgresql://postgres:postgres@postgres:5432/"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %#v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil
	}
	sqlDB.SetMaxIdleConns(50)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	err = db.AutoMigrate(models.AllModels...)
	if err != nil {
		log.Fatalf("failed to connect database: %#v", err)
	}

	return db
}
