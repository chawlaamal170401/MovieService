package database

import (
	"fmt"
	"github.com/razorpay/movie-service/internals/config"
	models "github.com/razorpay/movie-service/internals/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type DB struct {
	*gorm.DB
}

func NewDB() (*gorm.DB, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
		return nil, err
	}

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s port=%d sslmode=%s",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.DBName,
		cfg.Database.Port,
		cfg.Database.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
		return nil, err
	}

	if err := db.AutoMigrate(&models.Movie{}); err != nil {
		log.Fatalf("Error automigrating database: %v", err)
		return nil, err
	}

	log.Println("Successfully connected to the PostgreSQL database")

	return db, nil
}
