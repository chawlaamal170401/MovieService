package models

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	Title    string  `gorm:"type:varchar(100);not null"`
	Genre    string  `gorm:"type:varchar(50);not null"`
	Director string  `gorm:"type:varchar(100);not null"`
	Year     string  `gorm:"type:varchar(100);not null"`
	Rating   float64 `gorm:"type:decimal(3,1);not null"`
}
