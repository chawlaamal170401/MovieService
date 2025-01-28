package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type DB struct {
	*gorm.DB
}

type Movie struct {
	gorm.Model
	Title    string  `gorm:"type:varchar(100);not null"`
	Genre    string  `gorm:"type:varchar(50);not null"`
	Director string  `gorm:"type:varchar(100);not null"`
	Year     string  `gorm:"type:varchar(100);not null"`
	Rating   float64 `gorm:"type:decimal(3,1);not null"`
}

func NewDB() (*gorm.DB, error) {
	dsn := "user=movie_user password=password@123 dbname=movies_db port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
		return nil, err
	}

	if err := db.AutoMigrate(&Movie{}); err != nil {
		log.Fatalf("Error automigrating database: %v", err)
		return nil, err
	}

	log.Println("Successfully connected to the PostgreSQL database")

	return db, nil
}
