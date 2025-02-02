package repository

import (
	models "github.com/razorpay/movie-service/internals/model"
	pb "github.com/razorpay/movie-service/internals/proto"
)

type MovieRepositoryInterface interface {
	GetMovieByID(id int64) (*models.Movie, error)
	SaveMovie(movie *pb.Movie) (uint, error)
	GetAllMovies() ([]models.Movie, error)
	DeleteMovieByID(id int64) error
	SaveExternalMovie(movie *models.Movie) error
	UpdateMovieByID(id int64, updatedMovie *models.Movie) error
}
