package services

import (
	"errors"
	"movie-tracker/database"
	"movie-tracker/models"
	"time"

	"gorm.io/gorm"
)

type FavoritesService struct{}

func NewFavoritesService() *FavoritesService {
	return &FavoritesService{}
}

func (s *FavoritesService) AddToFavorites(userID uint, tmdbMovie *models.TMDBMovie, status models.Status, rating *int, notes, recommendedBy string) (*models.FavoriteMovie, error) {
	db := database.GetDB()

	var releaseDate *time.Time
	if tmdbMovie.ReleaseDate != "" {
		if t, err := time.Parse("2006-01-02", tmdbMovie.ReleaseDate); err == nil {
			releaseDate = &t
		}
	}

	favorite := &models.FavoriteMovie{
		UserID:        userID,
		TMDBId:        tmdbMovie.ID,
		Title:         tmdbMovie.Title,
		Overview:      tmdbMovie.Overview,
		ReleaseDate:   releaseDate,
		PosterPath:    tmdbMovie.PosterPath,
		GenreIDs:      models.IntArray(tmdbMovie.GenreIDs),
		Status:        status,
		Rating:        rating,
		Notes:         notes,
		RecommendedBy: recommendedBy,
	}

	if err := db.Create(favorite).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, errors.New("movie is already in favorites")
		}
		return nil, err
	}

	return favorite, nil
}

func (s *FavoritesService) GetUserFavorites(userID uint, status *models.Status, offset, limit int) ([]models.FavoriteMovie, error) {
	db := database.GetDB()
	var favorites []models.FavoriteMovie

	query := db.Where("user_id = ?", userID)
	if status != nil {
		query = query.Where("status = ?", *status)
	}

	query = query.Order("added_at DESC")
	if limit > 0 {
		query = query.Offset(offset).Limit(limit)
	}

	if err := query.Find(&favorites).Error; err != nil {
		return nil, err
	}

	return favorites, nil
}

func (s *FavoritesService) GetFavoriteByID(id, userID uint) (*models.FavoriteMovie, error) {
	db := database.GetDB()
	var favorite models.FavoriteMovie

	if err := db.Where("id = ? AND user_id = ?", id, userID).First(&favorite).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("favorite movie not found")
		}
		return nil, err
	}

	return &favorite, nil
}

func (s *FavoritesService) UpdateFavorite(id, userID uint, updates map[string]interface{}) (*models.FavoriteMovie, error) {
	db := database.GetDB()
	
	favorite, err := s.GetFavoriteByID(id, userID)
	if err != nil {
		return nil, err
	}

	if status, ok := updates["status"]; ok && status == models.StatusWatched {
		if favorite.WatchedAt == nil {
			now := time.Now()
			updates["watched_at"] = now
		}
	}

	if err := db.Model(favorite).Updates(updates).Error; err != nil {
		return nil, err
	}

	return favorite, nil
}

func (s *FavoritesService) UpdateStatus(id, userID uint, status models.Status) (*models.FavoriteMovie, error) {
	updates := map[string]interface{}{
		"status": status,
	}
	return s.UpdateFavorite(id, userID, updates)
}

func (s *FavoritesService) UpdateRating(id, userID uint, rating int) (*models.FavoriteMovie, error) {
	if rating < 1 || rating > 10 {
		return nil, errors.New("rating must be between 1 and 10")
	}

	updates := map[string]interface{}{
		"rating": rating,
	}
	return s.UpdateFavorite(id, userID, updates)
}

func (s *FavoritesService) DeleteFavorite(id, userID uint) error {
	db := database.GetDB()
	
	result := db.Where("id = ? AND user_id = ?", id, userID).Delete(&models.FavoriteMovie{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("favorite movie not found")
	}

	return nil
}

func (s *FavoritesService) GetUserStats(userID uint) (map[string]int, error) {
	db := database.GetDB()
	stats := make(map[string]int)

	var total int64
	db.Model(&models.FavoriteMovie{}).Where("user_id = ?", userID).Count(&total)
	stats["total"] = int(total)

	var toBe int64
	db.Model(&models.FavoriteMovie{}).Where("user_id = ? AND status = ?", userID, models.StatusToBe).Count(&toBe)
	stats["por_ver"] = int(toBe)

	var watched int64
	db.Model(&models.FavoriteMovie{}).Where("user_id = ? AND status = ?", userID, models.StatusWatched).Count(&watched)
	stats["vista"] = int(watched)

	var recommended int64
	db.Model(&models.FavoriteMovie{}).Where("user_id = ? AND status = ?", userID, models.StatusRecommended).Count(&recommended)
	stats["recomendada"] = int(recommended)

	return stats, nil
}