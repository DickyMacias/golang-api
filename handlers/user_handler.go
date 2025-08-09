package handlers

import (
	"movie-tracker/models"
	"movie-tracker/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	favoritesService *services.FavoritesService
}

func NewUserHandler() *UserHandler {
	return &UserHandler{
		favoritesService: services.NewFavoritesService(),
	}
}

func (h *UserHandler) GetStats(c *gin.Context) {
	user, _ := c.Get("user")
	userModel := user.(*models.User)

	stats, err := h.favoritesService.GetUserStats(userModel.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get stats"})
		return
	}

	c.JSON(http.StatusOK, stats)
}