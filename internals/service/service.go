package service

import (
	"context"
	"encoding/json"
	"fmt"
	models "github.com/razorpay/movie-service/internals/model"
	pb "github.com/razorpay/movie-service/internals/proto"
	"github.com/razorpay/movie-service/internals/repository"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type MovieService struct {
	pb.UnimplementedMovieServiceServer
	repo *repository.MovieRepository
}

func NewMovieServer(repo *repository.MovieRepository) *MovieService {
	return &MovieService{repo: repo}
}

func (s *MovieService) GetAllMovies(ctx context.Context, in *pb.Empty) (*pb.MovieListResponse, error) {
	log.Println("Fetching all movies from the database...")

	movies, err := s.repo.GetAllMovies()
	if err != nil {
		log.Printf("Error fetching movies: %v", err)
		return nil, err
	}

	var movieList []*pb.Movie
	for _, movie := range movies {
		movieList = append(movieList, &pb.Movie{
			Id:       int64(movie.ID),
			Title:    movie.Title,
			Genre:    movie.Genre,
			Director: movie.Director,
			Year:     movie.Year,
			Rating:   float32(movie.Rating),
		})
	}

	return &pb.MovieListResponse{Movies: movieList}, nil
}

func (s *MovieService) GetMovieByID(ctx context.Context, in *pb.MovieIDRequest) (*pb.Movie, error) {
	log.Printf("Fetching movie with ID: %d from the database...", in.GetId())

	movie, err := s.repo.GetMovieByID(in.GetId())
	if err != nil {
		log.Printf("Error fetching movie by ID: %v", err)
		return nil, err
	}

	return &pb.Movie{
		Id:       int64(movie.ID),
		Title:    movie.Title,
		Genre:    movie.Genre,
		Director: movie.Director,
		Year:     movie.Year,
		Rating:   float32(movie.Rating),
	}, nil
}

func (s *MovieService) CreateMovie(ctx context.Context, in *pb.MovieRequest) (*pb.Movie, error) {
	log.Printf("Received an add-movie request")

	created_movie := &pb.Movie{
		Title:    in.GetTitle(),
		Genre:    in.GetGenre(),
		Director: in.GetDirector(),
		Year:     in.GetYear(),
		Rating:   in.GetRating(),
	}

	log.Printf("Received movie: %+v", created_movie)

	id, err := s.repo.SaveMovie(created_movie)
	if err != nil {
		log.Printf("Error saving movie to database: %v", err)
		return nil, err
	}

	saved_movie := &pb.Movie{
		Id:       int64(id),
		Title:    in.GetTitle(),
		Genre:    in.GetGenre(),
		Director: in.GetDirector(),
		Year:     in.GetYear(),
		Rating:   in.GetRating(),
	}

	log.Println("Movie Saved Successfully to DB", id)
	return saved_movie, nil
}

func (s *MovieService) DeleteMovieByID(ctx context.Context, in *pb.MovieIDRequest) (*pb.ResponseMessage, error) {
	log.Printf("Deleting movie with ID: %d", in.GetId())

	err := s.repo.DeleteMovieByID(in.GetId())
	if err != nil {
		log.Printf("Error deleting movie: %v", err)
		return nil, err
	}
	return &pb.ResponseMessage{Message: "Movie Deleted Successfully"}, err
}

func (s *MovieService) UpdateMovie(ctx context.Context, in *pb.UpdateMovieRequest) (*pb.Movie, error) {
	log.Printf("Updating movie with ID: %d", in.GetId())

	updatedMovie := &models.Movie{
		Title:    in.GetTitle(),
		Genre:    in.GetGenre(),
		Director: in.GetDirector(),
		Year:     in.GetYear(),
		Rating:   float64(in.GetRating()),
	}

	log.Printf("Updated Movie %v: ", updatedMovie)

	err := s.repo.UpdateMovieByID(in.GetId(), updatedMovie)
	if err != nil {
		log.Printf("Error updating movie: %v", err)
		return nil, err
	}

	return &pb.Movie{
		Id:       in.GetId(),
		Title:    updatedMovie.Title,
		Genre:    updatedMovie.Genre,
		Director: updatedMovie.Director,
		Year:     updatedMovie.Year,
		Rating:   float32(updatedMovie.Rating),
	}, nil
}

func (s *MovieService) InitializeMovies(apiURL string) error {
	movies, err := s.repo.GetAllMovies()
	if err != nil {
		log.Printf("Error checking movies in the database: %v", err)
		return err
	}

	if len(movies) > 0 {
		log.Println("Movies already exist in the database. Skipping external API fetch.")
		return nil
	}

	log.Println("No movies found in the database. Fetching from external API...")
	return s.FetchMoviesFromExternalAPI(apiURL)
}

func (s *MovieService) FetchMoviesFromExternalAPI(apiURL string) error {
	resp, err := http.Get(apiURL)
	if err != nil {
		log.Printf("Error calling external API: %v", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("external API returned status: %d", resp.StatusCode)
	}

	var externalMovies []struct {
		Title    string   `json:"title"`
		Genre    []string `json:"genre"`
		Director string   `json:"director"`
		Year     int      `json:"year"`
		Rating   float32  `json:"rating"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&externalMovies); err != nil {
		log.Printf("Error decoding API response: %v", err)
		return err
	}

	for _, extMovie := range externalMovies {
		yearString := strconv.Itoa(extMovie.Year)

		genres := strings.Join(extMovie.Genre, ", ")

		movie := models.Movie{
			Title:    extMovie.Title,
			Genre:    genres,
			Director: extMovie.Director,
			Year:     yearString,
			Rating:   float64(extMovie.Rating),
		}

		err := s.repo.SaveExternalMovie(&movie)
		if err != nil {
			log.Printf("Error saving movie to database: %v", err)
			continue
		}

		log.Printf("Saved movie: %s", movie.Title)
	}

	return nil
}
