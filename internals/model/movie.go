package models

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	Title    string  `json:"title"`
	Genre    string  `json:"genre"`
	Director string  `json:"director"`
	Year     string  `json:"year"`
	Rating   float64 `json:"rating"`
}
