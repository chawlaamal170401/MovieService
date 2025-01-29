package database

import (
	models "github.com/razorpay/movie-service/internals/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type DB struct {
	*gorm.DB
}

func NewDB() (*gorm.DB, error) {
	dsn := "user=movie_user password=password@123 dbname=movies_db port=5432 sslmode=disable"

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
