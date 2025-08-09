package handlers

import (
	"movie-tracker/models"
	"movie-tracker/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FavoritesHandler struct {
	favoritesService *services.FavoritesService
	tmdbService      *services.TMDBService
}

func NewFavoritesHandler() *FavoritesHandler {
	return &FavoritesHandler{
		favoritesService: services.NewFavoritesService(),
		tmdbService:      services.NewTMDBService(),
	}
}

func (h *FavoritesHandler) ShowFavorites(c *gin.Context) {
	user, _ := c.Get("user")
	userModel := user.(*models.User)

	statusParam := c.Query("status")
	var status *models.Status
	if statusParam != "" {
		s := models.Status(statusParam)
		status = &s
	}

	favorites, err := h.favoritesService.GetUserFavorites(userModel.ID, status, 0, 0)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "favorites.html", gin.H{
			"title": "Favorites",
			"error": "Error loading favorites",
			"user":  userModel,
		})
		return
	}

	stats, _ := h.favoritesService.GetUserStats(userModel.ID)

	c.HTML(http.StatusOK, "favorites.html", gin.H{
		"title":     "Favorites",
		"user":      userModel,
		"favorites": favorites,
		"stats":     stats,
		"status":    statusParam,
	})
}

func (h *FavoritesHandler) ShowSearch(c *gin.Context) {
	user, _ := c.Get("user")
	userModel := user.(*models.User)

	c.HTML(http.StatusOK, "search.html", gin.H{
		"title": "Search Movies",
		"user":  userModel,
	})
}

func (h *FavoritesHandler) AddToFavorites(c *gin.Context) {
	user, _ := c.Get("user")
	userModel := user.(*models.User)

	tmdbIDStr := c.PostForm("tmdb_id")
	tmdbID, err := strconv.Atoi(tmdbIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid TMDB ID"})
		return
	}

	tmdbMovie, err := h.tmdbService.GetMovieDetails(tmdbID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching movie details"})
		return
	}

	status := models.Status(c.DefaultPostForm("status", string(models.StatusToBe)))
	notes := c.PostForm("notes")
	recommendedBy := c.PostForm("recommended_by")

	var rating *int
	if ratingStr := c.PostForm("rating"); ratingStr != "" {
		if r, err := strconv.Atoi(ratingStr); err == nil && r >= 1 && r <= 10 {
			rating = &r
		}
	}

	favorite, err := h.favoritesService.AddToFavorites(userModel.ID, tmdbMovie, status, rating, notes, recommendedBy)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if c.GetHeader("HX-Request") == "true" {
		c.HTML(http.StatusOK, "alert.html", gin.H{
			"type":    "success",
			"message": "Movie added to favorites!",
		})
		return
	}

	c.JSON(http.StatusCreated, favorite)
}

func (h *FavoritesHandler) UpdateStatus(c *gin.Context) {
	user, _ := c.Get("user")
	userModel := user.(*models.User)

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	status := models.Status(c.PostForm("status"))
	if status != models.StatusToBe && status != models.StatusWatched && status != models.StatusRecommended {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status"})
		return
	}

	_, err = h.favoritesService.UpdateStatus(uint(id), userModel.ID, status)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if c.GetHeader("HX-Request") == "true" {
		c.Header("HX-Refresh", "true")
		c.Status(http.StatusOK)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Status updated successfully"})
}

func (h *FavoritesHandler) UpdateRating(c *gin.Context) {
	user, _ := c.Get("user")
	userModel := user.(*models.User)

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	ratingStr := c.PostForm("rating")
	rating, err := strconv.Atoi(ratingStr)
	if err != nil || rating < 1 || rating > 10 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Rating must be between 1 and 10"})
		return
	}

	_, err = h.favoritesService.UpdateRating(uint(id), userModel.ID, rating)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if c.GetHeader("HX-Request") == "true" {
		c.Header("HX-Refresh", "true")
		c.Status(http.StatusOK)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Rating updated successfully"})
}

func (h *FavoritesHandler) DeleteFavorite(c *gin.Context) {
	user, _ := c.Get("user")
	userModel := user.(*models.User)

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = h.favoritesService.DeleteFavorite(uint(id), userModel.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if c.GetHeader("HX-Request") == "true" {
		c.Header("HX-Refresh", "true")
		c.Status(http.StatusOK)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Favorite deleted successfully"})
}