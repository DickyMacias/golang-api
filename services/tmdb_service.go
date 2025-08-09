package services

import (
	"encoding/json"
	"fmt"
	"movie-tracker/models"
	"net/http"
	"net/url"
	"os"
)

type TMDBService struct {
	apiKey  string
	baseURL string
}

func NewTMDBService() *TMDBService {
	return &TMDBService{
		apiKey:  os.Getenv("TMDB_API_KEY"),
		baseURL: "https://api.themoviedb.org/3",
	}
}

func (s *TMDBService) SearchMovies(query string, page int) (*models.TMDBResponse, error) {
	if page <= 0 {
		page = 1
	}

	endpoint := fmt.Sprintf("%s/search/movie", s.baseURL)
	params := url.Values{
		"api_key": {s.apiKey},
		"query":   {query},
		"page":    {fmt.Sprintf("%d", page)},
	}

	resp, err := http.Get(endpoint + "?" + params.Encode())
	if err != nil {
		return nil, fmt.Errorf("error making request to TMDB: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("TMDB API returned status code: %d", resp.StatusCode)
	}

	var tmdbResponse models.TMDBResponse
	if err := json.NewDecoder(resp.Body).Decode(&tmdbResponse); err != nil {
		return nil, fmt.Errorf("error decoding TMDB response: %w", err)
	}

	return &tmdbResponse, nil
}

func (s *TMDBService) GetPopularMovies(page int) (*models.TMDBResponse, error) {
	if page <= 0 {
		page = 1
	}

	endpoint := fmt.Sprintf("%s/movie/popular", s.baseURL)
	params := url.Values{
		"api_key": {s.apiKey},
		"page":    {fmt.Sprintf("%d", page)},
	}

	resp, err := http.Get(endpoint + "?" + params.Encode())
	if err != nil {
		return nil, fmt.Errorf("error making request to TMDB: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("TMDB API returned status code: %d", resp.StatusCode)
	}

	var tmdbResponse models.TMDBResponse
	if err := json.NewDecoder(resp.Body).Decode(&tmdbResponse); err != nil {
		return nil, fmt.Errorf("error decoding TMDB response: %w", err)
	}

	return &tmdbResponse, nil
}

func (s *TMDBService) GetTrendingMovies(page int) (*models.TMDBResponse, error) {
	if page <= 0 {
		page = 1
	}

	endpoint := fmt.Sprintf("%s/trending/movie/week", s.baseURL)
	params := url.Values{
		"api_key": {s.apiKey},
		"page":    {fmt.Sprintf("%d", page)},
	}

	resp, err := http.Get(endpoint + "?" + params.Encode())
	if err != nil {
		return nil, fmt.Errorf("error making request to TMDB: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("TMDB API returned status code: %d", resp.StatusCode)
	}

	var tmdbResponse models.TMDBResponse
	if err := json.NewDecoder(resp.Body).Decode(&tmdbResponse); err != nil {
		return nil, fmt.Errorf("error decoding TMDB response: %w", err)
	}

	return &tmdbResponse, nil
}

func (s *TMDBService) GetMovieDetails(movieID int) (*models.TMDBMovie, error) {
	endpoint := fmt.Sprintf("%s/movie/%d", s.baseURL, movieID)
	params := url.Values{
		"api_key": {s.apiKey},
	}

	resp, err := http.Get(endpoint + "?" + params.Encode())
	if err != nil {
		return nil, fmt.Errorf("error making request to TMDB: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("TMDB API returned status code: %d", resp.StatusCode)
	}

	var movie models.TMDBMovie
	if err := json.NewDecoder(resp.Body).Decode(&movie); err != nil {
		return nil, fmt.Errorf("error decoding TMDB response: %w", err)
	}

	return &movie, nil
}

func (s *TMDBService) GetMovieFullDetails(movieID int) (*models.TMDBMovieDetail, error) {
	endpoint := fmt.Sprintf("%s/movie/%d", s.baseURL, movieID)
	params := url.Values{
		"api_key": {s.apiKey},
	}

	resp, err := http.Get(endpoint + "?" + params.Encode())
	if err != nil {
		return nil, fmt.Errorf("error making request to TMDB: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("TMDB API returned status code: %d", resp.StatusCode)
	}

	var movieDetail models.TMDBMovieDetail
	if err := json.NewDecoder(resp.Body).Decode(&movieDetail); err != nil {
		return nil, fmt.Errorf("error decoding TMDB response: %w", err)
	}

	return &movieDetail, nil
}