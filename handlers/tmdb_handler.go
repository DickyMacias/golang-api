package handlers

import (
	"movie-tracker/services"
	"net/http"
	"strconv"
	"strings"

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

// formatNumber formats a number with comma separators
func formatNumber(n int64) string {
	s := strconv.FormatInt(n, 10)
	if len(s) <= 3 {
		return s
	}
	
	result := ""
	for i, char := range s {
		if i > 0 && (len(s)-i)%3 == 0 {
			result += ","
		}
		result += string(char)
	}
	return result
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

func (h *TMDBHandler) GetMovieDetail(c *gin.Context) {
	movieIDStr := c.Param("id")
	movieID, err := strconv.Atoi(movieIDStr)
	if err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Error",
			"error": "Invalid movie ID",
		})
		return
	}

	movieDetail, err := h.tmdbService.GetMovieFullDetails(movieID)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"title": "Error",
			"error": "Error loading movie details",
		})
		return
	}

	c.HTML(http.StatusOK, "movie_detail.html", gin.H{
		"title":          movieDetail.Title,
		"movie":          movieDetail,
		"user":           c.MustGet("user"),
		"formatBudget":   formatNumber(movieDetail.Budget),
		"formatRevenue":  formatNumber(movieDetail.Revenue),
	})
}