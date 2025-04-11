package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	routesData "github.com/popflix-live/api/src/lib/data"
	"github.com/popflix-live/api/src/lib/models/router"
)

func RegisterRoutes(r *gin.Engine) {
	routes := routesData.GetRoutes()

	for _, route := range routes {
		registerRoute(r, route)
	}
}

func registerRoute(r *gin.Engine, route router.RouteConfig) {
	switch route.Method {
	case "GET":
		r.GET(route.Path, func(c *gin.Context) {
			route.Handler(c)
		})
	case "POST":
		r.POST(route.Path, func(c *gin.Context) {
			route.Handler(c)
		})
	case "PUT":
		r.PUT(route.Path, func(c *gin.Context) {
			route.Handler(c)
		})
	case "DELETE":
		r.DELETE(route.Path, func(c *gin.Context) {
			route.Handler(c)
		})
	case "PATCH":
		r.PATCH(route.Path, func(c *gin.Context) {
			route.Handler(c)
		})
	default:
		fmt.Printf("Unsupported HTTP method: %s\n", route.Method)
	}
}
