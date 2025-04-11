package tmdb

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	
	models "github.com/popflix-live/api/src/lib/models/tmdb"
)

const (
	BaseURL = "https://api.themoviedb.org/3"
	ImageBaseURL = "https://image.tmdb.org/t/p/original"
)

type Client struct {
	apiKey string
	httpClient *http.Client
}

func New() (*Client, error) {
	apiKey := os.Getenv("TMDB_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("TMDB_API_KEY environment variable not set")
	}

	return &Client{
		apiKey: apiKey,
		httpClient: &http.Client{},
	}, nil
}

func (c *Client) SearchTV(query string) (*models.SearchResponse, error) {
	endpoint := fmt.Sprintf("%s/search/tv?api_key=%s&query=%s", 
		BaseURL, c.apiKey, url.QueryEscape(query))
	
	return c.getSearchResults(endpoint)
}

func (c *Client) SearchMulti(query string) (*models.SearchResponse, error) {
	endpoint := fmt.Sprintf("%s/search/multi?api_key=%s&query=%s", 
		BaseURL, c.apiKey, url.QueryEscape(query))
	
	return c.getSearchResults(endpoint)
}

func (c *Client) GetTVImages(id string) (*models.ImagesResponse, error) {
	endpoint := fmt.Sprintf("%s/tv/%s/images?api_key=%s", BaseURL, id, c.apiKey)
	return c.getImages(endpoint)
}

func (c *Client) GetMovieImages(id string) (*models.ImagesResponse, error) {
	endpoint := fmt.Sprintf("%s/movie/%s/images?api_key=%s", BaseURL, id, c.apiKey)
	return c.getImages(endpoint)
}

func (c *Client) getSearchResults(endpoint string) (*models.SearchResponse, error) {
	resp, err := c.httpClient.Get(endpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("TMDB API returned status: %d", resp.StatusCode)
	}

	var result models.SearchResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) getImages(endpoint string) (*models.ImagesResponse, error) {
	resp, err := c.httpClient.Get(endpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("TMDB API returned status: %d", resp.StatusCode)
	}

	var result models.ImagesResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}