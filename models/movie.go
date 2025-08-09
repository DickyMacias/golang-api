package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type IntArray []int

func (ia *IntArray) Scan(value interface{}) error {
	if value == nil {
		*ia = []int{}
		return nil
	}
	
	switch v := value.(type) {
	case []byte:
		return json.Unmarshal(v, ia)
	case string:
		return json.Unmarshal([]byte(v), ia)
	default:
		return errors.New("cannot scan non-string into IntArray")
	}
}

func (ia IntArray) Value() (driver.Value, error) {
	if len(ia) == 0 {
		return "[]", nil
	}
	return json.Marshal(ia)
}

type TMDBMovie struct {
	ID           int      `json:"id"`
	Title        string   `json:"title"`
	Overview     string   `json:"overview"`
	ReleaseDate  string   `json:"release_date"`
	PosterPath   string   `json:"poster_path"`
	GenreIDs     []int    `json:"genre_ids"`
	VoteAverage  float64  `json:"vote_average"`
	VoteCount    int      `json:"vote_count"`
	Adult        bool     `json:"adult"`
	OriginalTitle string  `json:"original_title"`
	Popularity   float64  `json:"popularity"`
}

type TMDBResponse struct {
	Page         int         `json:"page"`
	Results      []TMDBMovie `json:"results"`
	TotalPages   int         `json:"total_pages"`
	TotalResults int         `json:"total_results"`
}