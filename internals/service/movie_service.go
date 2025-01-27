package services

import (
	"github.com/razorpay/movie-service/internals/model"
	repositories "github.com/razorpay/movie-service/internals/repository"
)

type MovieService interface {
	Create(movie *model.Movie) error
	GetByID(id uint) (*model.Movie, error)
	GetAll() ([]model.Movie, error)
	Update(movie *model.Movie) error
	Delete(id uint) error
}

type movieService struct {
	repo repositories.MovieRepository
}

// NewMovieService creates a new instance of the MovieService
func NewMovieService(repo repositories.MovieRepository) MovieService {
	return &movieService{repo: repo}
}

func (s *movieService) Create(movie *model.Movie) error {
	return s.repo.CreateMovie(movie)
}

func (s *movieService) GetByID(id uint) (*model.Movie, error) {
	return s.repo.GetMovieByID(id)
}

func (s *movieService) GetAll() ([]model.Movie, error) {
	return s.repo.GetAllMovies()
}

func (s *movieService) Update(movie *model.Movie) error {
	return s.repo.UpdateMovie(movie)
}

func (s *movieService) Delete(id uint) error {
	return s.repo.DeleteMovieByID(id)
}
