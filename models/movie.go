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

type Genre struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ProductionCompany struct {
	ID            int    `json:"id"`
	LogoPath      string `json:"logo_path"`
	Name          string `json:"name"`
	OriginCountry string `json:"origin_country"`
}

type ProductionCountry struct {
	ISO31661 string `json:"iso_3166_1"`
	Name     string `json:"name"`
}

type SpokenLanguage struct {
	EnglishName string `json:"english_name"`
	ISO6391     string `json:"iso_639_1"`
	Name        string `json:"name"`
}

type TMDBMovieDetail struct {
	Adult               bool                `json:"adult"`
	BackdropPath        string              `json:"backdrop_path"`
	Budget              int64               `json:"budget"`
	Genres              []Genre             `json:"genres"`
	Homepage            string              `json:"homepage"`
	ID                  int                 `json:"id"`
	IMDBId              string              `json:"imdb_id"`
	OriginCountry       []string            `json:"origin_country"`
	OriginalLanguage    string              `json:"original_language"`
	OriginalTitle       string              `json:"original_title"`
	Overview            string              `json:"overview"`
	Popularity          float64             `json:"popularity"`
	PosterPath          string              `json:"poster_path"`
	ProductionCompanies []ProductionCompany `json:"production_companies"`
	ProductionCountries []ProductionCountry `json:"production_countries"`
	ReleaseDate         string              `json:"release_date"`
	Revenue             int64               `json:"revenue"`
	Runtime             int                 `json:"runtime"`
	SpokenLanguages     []SpokenLanguage    `json:"spoken_languages"`
	Status              string              `json:"status"`
	Tagline             string              `json:"tagline"`
	Title               string              `json:"title"`
	Video               bool                `json:"video"`
	VoteAverage         float64             `json:"vote_average"`
	VoteCount           int                 `json:"vote_count"`
}

type TMDBResponse struct {
	Page         int         `json:"page"`
	Results      []TMDBMovie `json:"results"`
	TotalPages   int         `json:"total_pages"`
	TotalResults int         `json:"total_results"`
}