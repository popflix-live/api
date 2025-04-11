package recent

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	instance "github.com/popflix-live/api/src/lib/http"
)

type Anime struct {
	AnimeName  string `json:"anime_name"`
	AnimeImage string `json:"anime_image"`
}
type AnimeList struct {
	Data []Anime `json:"data"`
}

func GetHandler(c *gin.Context) {
	url := "/series/?status=&type=&order=update"
	resp, err := instance.Client.R().
		Get(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data"})
		return
	} else if resp.StatusCode() != http.StatusOK {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data"})
		return
	} else if len(resp.Body()) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No data found"})
		return
	}
	var result AnimeList
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(string(resp.Body())))

	fmt.Println("Parsing HTML...")
	fmt.Println(doc.Find(".bs").Length())
	doc.Find(".bs").Each(func(i int, s *goquery.Selection) {
		s.Find(".bsx").Each(func(i int, s *goquery.Selection) {
			s.Find(".tip").Each(func(i int, s *goquery.Selection) {
				s.Find("img").Each(func(i int, s *goquery.Selection) {
					animeName, _ := s.Attr("title")
					animeImage, _ := s.Attr("src")
					result.Data = append(result.Data, Anime{
						AnimeName:  animeName,
						AnimeImage: animeImage,
					})
				})
			})

		})
	})
	fmt.Println(result)
	c.JSON(http.StatusOK, result.Data)
}
