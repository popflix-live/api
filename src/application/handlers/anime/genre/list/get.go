package genre

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	httpClient "github.com/popflix-live/api/src/lib/http"
)

type Genre struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Count int    `json:"count"`
}

func GetHandler(c *gin.Context) {
	client := httpClient.NewHttpClient("")

	resp, err := client.Get("http://localhost:3000/anime/gogoanime/genre/list", nil)
	if err != nil {
		respondWithError(c, "Failed to fetch genres", err)
		return
	}

	var genres []Genre
	if err := json.Unmarshal(resp.Body, &genres); err != nil {
		respondWithError(c, "Failed to parse genres", err)
		return
	}

	for i, genre := range genres {
		count, err := getAnimeCountFromKitsu(genre.Title)
		if err != nil {
			count = 0
		}
		genres[i].Count = count
	}

	c.JSON(http.StatusOK, genres)
}

func getAnimeCountFromKitsu(genreName string) (int, error) {
	url := fmt.Sprintf("https://kitsu.io/api/edge/anime?filter[categories]=%s&page[limit]=1", genreName)

	client := httpClient.NewHttpClient("")
	resp, err := client.Get(url, map[string]string{
		"Accept": "application/vnd.api+json",
	})

	if err != nil {
		return 0, err
	}

	var result struct {
		Data []interface{} `json:"data"`
		Meta struct {
			Count int `json:"count"`
		} `json:"meta"`
	}

	if err := json.Unmarshal(resp.Body, &result); err != nil {
		return 0, err
	}

	if result.Meta.Count == 0 {
		return len(result.Data), nil
	}

	return result.Meta.Count, nil
}

func respondWithError(c *gin.Context, message string, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": fmt.Sprintf("%s: %v", message, err),
	})
}
