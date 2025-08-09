package models

import (
	"time"

	"gorm.io/gorm"
)

type Status string

const (
	StatusToBe         Status = "por_ver"
	StatusWatched      Status = "vista"
	StatusRecommended  Status = "recomendada"
)

type FavoriteMovie struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	UserID        uint           `gorm:"not null;index" json:"user_id"`
	TMDBId        int            `gorm:"not null" json:"tmdb_id"`
	Title         string         `gorm:"not null;size:255" json:"title"`
	Overview      string         `gorm:"type:text" json:"overview"`
	ReleaseDate   *time.Time     `json:"release_date"`
	PosterPath    string         `gorm:"size:255" json:"poster_path"`
	GenreIDs      IntArray       `gorm:"type:jsonb" json:"genre_ids"`
	Status        Status         `gorm:"type:varchar(20);default:'por_ver'" json:"status"`
	Rating        *int           `gorm:"check:rating >= 1 AND rating <= 10" json:"rating"`
	Notes         string         `gorm:"type:text" json:"notes"`
	RecommendedBy string         `gorm:"size:100" json:"recommended_by"`
	AddedAt       time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"added_at"`
	WatchedAt     *time.Time     `json:"watched_at"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
	
	User User `gorm:"foreignKey:UserID" json:"user,omitempty"`
}

func (FavoriteMovie) TableName() string {
	return "favorite_movies"
}

func (fm *FavoriteMovie) BeforeCreate(tx *gorm.DB) error {
	if fm.AddedAt.IsZero() {
		fm.AddedAt = time.Now()
	}
	return nil
}

func (fm *FavoriteMovie) BeforeUpdate(tx *gorm.DB) error {
	if fm.Status == StatusWatched && fm.WatchedAt == nil {
		now := time.Now()
		fm.WatchedAt = &now
	}
	return nil
}