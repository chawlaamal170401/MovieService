package repository

import (
	"github.com/razorpay/movie-service/internals/database"
	"github.com/razorpay/movie-service/internals/proto"
	"gorm.io/gorm"
	"log"
)

type MovieRepository struct {
	db *gorm.DB
}

func NewMovieRepository(db *gorm.DB) *MovieRepository {
	return &MovieRepository{db: db}
}

func (r *MovieRepository) SaveMovie(movie *proto.Movie) error {
	dbMovie := database.Movie{
		Title:    movie.Title,
		Genre:    movie.Genre,
		Director: movie.Director,
		Year:     movie.Year,
		Rating:   float64(movie.Rating),
	}

	result := r.db.Create(&dbMovie)
	if result.Error != nil {
		log.Printf("Error saving movie to DB: %v", result.Error)
		return result.Error
	}

	log.Printf("Movie saved with ID: %d", dbMovie.ID)
	return nil
}
