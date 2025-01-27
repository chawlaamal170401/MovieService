package repositories

import (
	"github.com/razorpay/movie-service/internals/model"
	"gorm.io/gorm"
)

type MovieRepository interface {
	CreateMovie(movie *model.Movie) error
	GetMovieByID(id uint) (*model.Movie, error)
	GetAllMovies() ([]model.Movie, error)
	UpdateMovie(movie *model.Movie) error
	DeleteMovieByID(id uint) error
}

type movieRepository struct {
	db *gorm.DB
}

func NewMovieRepository(db *gorm.DB) MovieRepository {
	return &movieRepository{db: db}
}

func (r *movieRepository) CreateMovie(movie *model.Movie) error {
	return r.db.Create(movie).Error
}

func (r *movieRepository) GetMovieByID(id uint) (*model.Movie, error) {
	var movie model.Movie
	if err := r.db.First(&movie, id).Error; err != nil {
		return nil, err
	}
	return &movie, nil
}

func (r *movieRepository) GetAllMovies() ([]model.Movie, error) {
	var movies []model.Movie
	if err := r.db.Find(&movies).Error; err != nil {
		return nil, err
	}
	return movies, nil
}

func (r *movieRepository) UpdateMovie(movie *model.Movie) error {
	return r.db.Save(movie).Error
}

func (r *movieRepository) DeleteMovieByID(id uint) error {
	return r.db.Delete(&model.Movie{}, id).Error
}
