package config

import (
	"github.com/gin-gonic/gin"

	animeGenre "github.com/popflix-live/api/src/application/handlers/anime/genre/list"
)

type RouteConfig struct {
	Method      string
	Path        string
	Handler     func(*gin.Context)
	Description string
}

func GetRoutes() []RouteConfig {
	return []RouteConfig{
		{
			Method:      "GET",
			Path:        "/anime/genre/list",
			Handler:     animeGenre.GetHandler,
			Description: "Get a list of anime genres",
		},
	}
}
