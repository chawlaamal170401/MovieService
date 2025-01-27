package controllers

import (
	"github.com/razorpay/movie-service/internals/model"
	services "github.com/razorpay/movie-service/internals/service"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MovieControllerInterface interface {
	CreateMovie(ctx *gin.Context)
	GetMovieByID(ctx *gin.Context)
	GetAllMovies(ctx *gin.Context)
	UpdateMovie(ctx *gin.Context)
	DeleteMovie(ctx *gin.Context)
}

type MovieController struct {
	service services.MovieService
}

func NewMovieController(service services.MovieService) *MovieController {
	return &MovieController{service: service}
}

func (c *MovieController) CreateMovie(ctx *gin.Context) {
	var movie model.Movie
	if err := ctx.ShouldBindJSON(&movie); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Printf("ReleaseYear from payload: %s", movie)
	if err := c.service.CreateMovie(&movie); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, movie)
}

func (c *MovieController) GetMovieByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	movie, err := c.service.GetMovieByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}
	ctx.JSON(http.StatusOK, movie)
}

func (c *MovieController) GetAllMovies(ctx *gin.Context) {
	movies, err := c.service.GetAllMovies()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, movies)
}

func (c *MovieController) UpdateMovie(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var movie model.Movie
	if err := ctx.ShouldBindJSON(&movie); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	movie.ID = int64(uint(id))
	if err := c.service.UpdateMovie(&movie); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, movie)
}

func (c *MovieController) DeleteMovie(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := c.service.DeleteMovieByID(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Movie deleted"})
}
