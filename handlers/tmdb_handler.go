package handlers

import (
	"movie-tracker/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TMDBHandler struct {
	tmdbService *services.TMDBService
}

func NewTMDBHandler() *TMDBHandler {
	return &TMDBHandler{
		tmdbService: services.NewTMDBService(),
	}
}

func (h *TMDBHandler) SearchMovies(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Query parameter 'q' is required"})
		return
	}

	pageStr := c.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}

	results, err := h.tmdbService.SearchMovies(query, page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if c.GetHeader("HX-Request") == "true" {
		c.HTML(http.StatusOK, "movie_card.html", gin.H{
			"movies": results.Results,
		})
		return
	}

	c.JSON(http.StatusOK, results)
}

func (h *TMDBHandler) GetPopularMovies(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}

	results, err := h.tmdbService.GetPopularMovies(page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if c.GetHeader("HX-Request") == "true" {
		c.HTML(http.StatusOK, "movie_card.html", gin.H{
			"movies": results.Results,
		})
		return
	}

	c.JSON(http.StatusOK, results)
}

func (h *TMDBHandler) GetTrendingMovies(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}

	results, err := h.tmdbService.GetTrendingMovies(page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if c.GetHeader("HX-Request") == "true" {
		c.HTML(http.StatusOK, "movie_card.html", gin.H{
			"movies": results.Results,
		})
		return
	}

	c.JSON(http.StatusOK, results)
}