package config

import (
	animeBackdrops "github.com/popflix-live/api/src/application/handlers/anime/backdrop"
	animeGenre "github.com/popflix-live/api/src/application/handlers/anime/genre/list"
	"github.com/popflix-live/api/src/lib/models/router"
)

func GetRoutes() []router.RouteConfig {
	return []router.RouteConfig{
		{
			Method:      "GET",
			Path:        "/anime/genre/list",
			Handler:     animeGenre.GetHandler,
			Description: "Get a list of anime genres",
		},
		{
			Method:      "GET",
			Path:        "/anime/backdrop/:id",
			Handler:     animeBackdrops.GetHandler,
			Description: "Get backdrop images for a movie or TV show by ID",
		},
		{
			Method:      "GET",
			Path:        "/anime/backdrop",
			Handler:     animeBackdrops.GetHandler,
			Description: "Get backdrop images for a movie or TV show by name",
		},
	}
}
