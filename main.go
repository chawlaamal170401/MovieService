package main

import (
	"github.com/razorpay/movie-service/internals/controller"
	"github.com/razorpay/movie-service/internals/database"
	"github.com/razorpay/movie-service/internals/repository"
	"github.com/razorpay/movie-service/internals/service"

	"github.com/gin-gonic/gin"
)

func main() {
	dsn := "host=localhost port=5432 user=movie_user password=password@123 dbname=movies_db sslmode=disable"
	database := database.ConnectDB(dsn)

	movieRepo := repositories.NewMovieRepository(database)
	movieService := services.NewMovieService(movieRepo)
	movieController := controllers.NewMovieController(movieService)

	router := gin.Default()

	router.POST("/movie", movieController.CreateMovie)
	router.GET("/movie/:id", movieController.GetMovieByID)
	router.GET("/movies", movieController.GetAllMovies)
	router.PUT("/movie/:id", movieController.UpdateMovie)
	router.DELETE("/movie/:id", movieController.DeleteMovie)

	router.Run(":8080")
}
