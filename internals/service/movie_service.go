package services

import (
	"github.com/razorpay/movie-service/internals/model"
	repositories "github.com/razorpay/movie-service/internals/repository"
)

type MovieService interface {
	CreateMovie(movie *model.Movie) error
	GetMovieByID(id uint) (*model.Movie, error)
	GetAllMovies() ([]model.Movie, error)
	UpdateMovie(movie *model.Movie) error
	DeleteMovieByID(id uint) error
}

type movieService struct {
	repo repositories.MovieRepository
}

func NewMovieService(repo repositories.MovieRepository) MovieService {
	return &movieService{repo: repo}
}

func (s *movieService) CreateMovie(movie *model.Movie) error {
	return s.repo.CreateMovie(movie)
}

func (s *movieService) GetMovieByID(id uint) (*model.Movie, error) {
	return s.repo.GetMovieByID(id)
}

func (s *movieService) GetAllMovies() ([]model.Movie, error) {
	return s.repo.GetAllMovies()
}

func (s *movieService) UpdateMovie(movie *model.Movie) error {
	return s.repo.UpdateMovie(movie)
}

func (s *movieService) DeleteMovieByID(id uint) error {
	return s.repo.DeleteMovieByID(id)
}
