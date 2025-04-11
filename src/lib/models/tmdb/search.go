package tmdb

type SearchResult struct {
	ID          int    `json:"id"`
	Title       string `json:"title,omitempty"`
	Name        string `json:"name,omitempty"`
	ReleaseDate string `json:"release_date,omitempty"`
	MediaType   string `json:"media_type"`
}

type SearchResponse struct {
	Results []SearchResult `json:"results"`
}