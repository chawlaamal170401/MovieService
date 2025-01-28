package main

import (
	"context"
	pb "github.com/razorpay/movie-service/internals/proto"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc"
	"log"
	"math/rand"
	"net"
)

const (
	port = ":8080"
)

func NewMovieServer() *MovieServer {
	return &MovieServer{
		movie_list: &pb.MovieListResponse{},
	}
}

type MovieServer struct {
	pb.UnimplementedMovieServiceServer
	movie_list *pb.MovieListResponse
}

func (server *MovieServer) Run() error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMovieServiceServer(s, server)
	log.Printf("Server Running on : %v", lis.Addr())
	return s.Serve(lis)
}

func (s *MovieServer) CreateMovie(ctx context.Context, in *pb.MovieRequest) (*pb.Movie, error) {
	log.Printf("Recieved NewMovie: %s", in.GetTitle())
	var movie_id int64 = int64(rand.Intn(1000))
	created_movie := &pb.Movie{Title: in.GetTitle(), Genre: in.GetGenre(), Id: movie_id, Director: in.GetDirector(), Year: in.GetYear(), Rating: in.GetRating()}
	s.movie_list.Movies = append(s.movie_list.Movies, created_movie)
	return created_movie, nil
}

func (s *MovieServer) GetAllMovies(ctx context.Context, in *pb.Empty) (*pb.MovieListResponse, error) {
	return s.movie_list, nil
}

func main() {
	var movie_server *MovieServer = NewMovieServer()
	if err := movie_server.Run(); err != nil {
		log.Fatalf("Failed to run %v", err)
	}
}
