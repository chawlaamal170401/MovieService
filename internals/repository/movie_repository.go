package repository

import (
	models "github.com/razorpay/movie-service/internals/model"
	pb "github.com/razorpay/movie-service/internals/proto"
	"gorm.io/gorm"
	"log"
)

type MovieRepository struct {
	db *gorm.DB
}

func NewMovieRepository(db *gorm.DB) *MovieRepository {
	return &MovieRepository{db: db}
}

func (r *MovieRepository) GetMovieByID(id int64) (*models.Movie, error) {
	var movie models.Movie
	result := r.db.First(&movie, id)

	if result.Error != nil {
		log.Printf("Error retrieving movie by ID %d: %v", id, result.Error)
		return nil, result.Error
	}

	return &movie, nil
}

func (r *MovieRepository) SaveMovie(movie *pb.Movie) error {
	dbMovie := models.Movie{
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

func (r *MovieRepository) GetAllMovies() ([]models.Movie, error) {
	var movies []models.Movie
	result := r.db.Find(&movies)

	if result.Error != nil {
		log.Printf("Error retrieving movies: %v", result.Error)
		return nil, result.Error
	}

	return movies, nil
}

func (r *MovieRepository) DeleteMovieByID(id int64) error {
	result := r.db.Delete(&models.Movie{}, id)
	if result.Error != nil {
		log.Printf("Error deleting movie by ID %d: %v", id, result.Error)
		return result.Error
	}
	return nil
}

func (repo *MovieRepository) SaveExternalMovie(movie *models.Movie) error {
	result := repo.db.Create(movie)
	if result.Error != nil {
		log.Printf("Error saving movie: %v", result.Error)
		return result.Error
	}
	return nil
}

func (repo *MovieRepository) UpdateMovieByID(id int64, updatedMovie *models.Movie) error {
	var movie models.Movie
	if err := repo.db.First(&movie, id).Error; err != nil {
		log.Printf("Movie not found: %v", err)
		return err
	}

	movie.Title = updatedMovie.Title
	movie.Genre = updatedMovie.Genre
	movie.Director = updatedMovie.Director
	movie.Year = updatedMovie.Year
	movie.Rating = float64(updatedMovie.Rating)

	if err := repo.db.Save(&movie).Error; err != nil {
		log.Printf("Error updating movie: %v", err)
		return err
	}

	return nil
}
