package service

import (
	"context"
	"github.com/razorpay/movie-service/internals/database"
	pb "github.com/razorpay/movie-service/internals/proto"
	"github.com/razorpay/movie-service/internals/repository"
	"google.golang.org/grpc/metadata"
	"log"
	"math/rand"
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

	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		log.Printf("Received metadata: %v", md)
	} else {
		log.Println("No metadata received")
	}

	var movie_id = int64(rand.Intn(1000))
	created_movie := &pb.Movie{
		Title:    in.GetTitle(),
		Genre:    in.GetGenre(),
		Id:       movie_id,
		Director: in.GetDirector(),
		Year:     in.GetYear(),
		Rating:   in.GetRating(),
	}

	log.Printf("Received movie: %+v", created_movie)

	err := s.repo.SaveMovie(created_movie)
	if err != nil {
		log.Printf("Error saving movie to database: %v", err)
		return nil, err
	}

	log.Println("Movie Saved Successfully to DB")
	return created_movie, nil
}

func (s *MovieService) DeleteMovieByID(ctx context.Context, in *pb.MovieIDRequest) (*pb.Empty, error) {
	log.Printf("Deleting movie with ID: %d", in.GetId())

	err := s.repo.DeleteMovieByID(in.GetId())
	if err != nil {
		log.Printf("Error deleting movie: %v", err)
		return nil, err
	}
	return &pb.Empty{}, err
}

func (s *MovieService) UpdateMovie(ctx context.Context, in *pb.UpdateMovieRequest) (*pb.Movie, error) {
	log.Printf("Updating movie with ID: %d", in.GetId())

	updatedMovie := &database.Movie{
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
