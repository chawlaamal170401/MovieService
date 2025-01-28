package service

import (
	"context"
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
