package backdrops

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/popflix-live/api/src/lib/tmdb"
)

func GetHandler(c *gin.Context) {
	animeName := c.Query("name")
	mediaID := c.Param("id")

	client, err := tmdb.New()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to initialize TMDB client: %v", err)})
		return
	}

	if animeName != "" {
		searchResp, err := client.SearchTV(animeName)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to search anime: %v", err)})
			return
		}

		if len(searchResp.Results) == 0 {
			searchResp, err = client.SearchMulti(animeName)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to search anime: %v", err)})
				return
			}

			if len(searchResp.Results) == 0 {
				c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("No results found for '%s'", animeName)})
				return
			}
		}

		mediaID = strconv.Itoa(searchResp.Results[0].ID)
	}

	if mediaID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Please provide either an ID parameter or a name query"})
		return
	}

	imagesResp, err := client.GetTVImages(mediaID)
	if err != nil || len(imagesResp.Backdrops) == 0 {
		imagesResp, err = client.GetMovieImages(mediaID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to fetch images: %v", err)})
			return
		}
	}

	if len(imagesResp.Backdrops) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No backdrop images found"})
		return
	}

	targetWidth, targetHeight := 2000, 1125
	closestBackdrop := imagesResp.Backdrops[0]
	closestDiff := calculateResolutionDifference(closestBackdrop.Width, closestBackdrop.Height, targetWidth, targetHeight)

	for _, backdrop := range imagesResp.Backdrops {
		if backdrop.Width == targetWidth && backdrop.Height == targetHeight {
			closestBackdrop = backdrop
			break
		}

		diff := calculateResolutionDifference(backdrop.Width, backdrop.Height, targetWidth, targetHeight)
		if diff < closestDiff {
			closestDiff = diff
			closestBackdrop = backdrop
		}
	}

	backdropURL := tmdb.ImageBaseURL + closestBackdrop.FilePath

	c.JSON(http.StatusOK, gin.H{
		"backdrop": backdropURL,
		"resolution": gin.H{
			"width":  closestBackdrop.Width,
			"height": closestBackdrop.Height,
		},
		"all_backdrops": len(imagesResp.Backdrops),
	})
}

func calculateResolutionDifference(width, height, targetWidth, targetHeight int) int {
	aspectRatioDiff := abs(float64(width)/float64(height) - float64(targetWidth)/float64(targetHeight))
	sizeDiff := abs(float64(width*height - targetWidth*targetHeight))

	return int(aspectRatioDiff*1000) + int(sizeDiff)/1000
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func GetRoute() (string, string, gin.HandlerFunc) {
	return "GET", "/anime/backdrop/:id", GetHandler
}
