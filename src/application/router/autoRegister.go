// This file is auto-generated. Do not edit manually.
package router

import (
    "fmt"

    "github.com/gin-gonic/gin"
	backdrop "github.com/popflix-live/api/src/application/handlers/anime/backdrop"
	list "github.com/popflix-live/api/src/application/handlers/anime/genre/list"
	recent "github.com/popflix-live/api/src/application/handlers/anime/recent"
    "github.com/popflix-live/api/src/lib/models/router"
)

// AutoRegisterRoutes automatically registers all handler functions from the handlers directory
func AutoRegisterRoutes(r *gin.Engine) {
    routes := []router.RouteConfig{
        {
            Method:      "GET",
            Path:        "/anime/backdrop",
            Handler:     backdrop.GetHandler,
            Description: "Auto-generated route for /anime/backdrop",
        },
        {
            Method:      "GET",
            Path:        "/anime/genre/list",
            Handler:     list.GetHandler,
            Description: "Auto-generated route for /anime/genre/list",
        },
        {
            Method:      "GET",
            Path:        "/anime/recent",
            Handler:     recent.GetHandler,
            Description: "Auto-generated route for /anime/recent",
        },
    }
    
    for _, route := range routes {
        registerRoute(r, route)
    }
    
    fmt.Println("Auto-registered", len(routes), "routes")
}
func registerRoute(r *gin.Engine, route router.RouteConfig) {
    switch route.Method {
    case "GET":
        r.GET(route.Path, route.Handler)
    case "POST":
        r.POST(route.Path, route.Handler)
    case "PUT":
        r.PUT(route.Path, route.Handler)
    case "DELETE":
        r.DELETE(route.Path, route.Handler)
    default:
        fmt.Println("Unsupported method:", route.Method)
    }
}
