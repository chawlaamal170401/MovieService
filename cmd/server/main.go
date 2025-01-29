package main

import (
	"github.com/razorpay/movie-service/internals/database"
	pb "github.com/razorpay/movie-service/internals/proto"
	"github.com/razorpay/movie-service/internals/repository"
	"github.com/razorpay/movie-service/internals/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"
	"log"
	"net"
)

func setupDatabase() (*gorm.DB, error) {
	db, err := database.NewDB()
	if err != nil {
		return nil, err
	}
	log.Println("Successfully connected to the database")
	return db, nil
}

func setupGRPCServer(addr string, movieService *service.MovieService) (*grpc.Server, net.Listener, error) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, nil, err
	}

	server := grpc.NewServer()
	reflection.Register(server)
	pb.RegisterMovieServiceServer(server, movieService)

	log.Printf("gRPC server is listening on %v", addr)
	return server, listener, nil
}

func main() {
	const addr = "0.0.0.0:8080"

	log.Println("Starting gRPC server...")

	db, err := setupDatabase()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			log.Printf("Error getting underlying sql.DB: %v", err)
		} else {
			if err := sqlDB.Close(); err != nil {
				log.Printf("Error closing database connection: %v", err)
			} else {
				log.Println("Database connection closed successfully")
			}
		}
	}()

	movieRepo := repository.NewMovieRepository(db)
	movieService := service.NewMovieServer(movieRepo)

	apiURL := "https://freetestapi.com/api/v1/movies"
	if err := movieService.InitializeMovies(apiURL); err != nil {
		log.Fatalf("Failed to initialize movies: %v", err)
	}

	server, listener, err := setupGRPCServer(addr, movieService)
	if err != nil {
		log.Fatalf("Failed to setup gRPC server: %v", err)
	}

	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
