package genre

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	instance "github.com/popflix-live/api/src/lib/http"
)

type Genre struct {
	MalID int    `json:"mal_id"`
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type GenreResponse struct {
	Data []Genre `json:"data"`
}

func GetHandler(c *gin.Context) {
	url := "https://api.jikan.moe/v4/genres/anime"

	resp, err := instance.Client.R().
		SetHeader("Accept", "application/json").
		Get(url)
	if err != nil {
		log.Fatal(err)
	}

	var result GenreResponse
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, result.Data)
}
