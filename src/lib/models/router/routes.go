package router

import (
	"github.com/gin-gonic/gin"
)

type RouteConfig struct {
	Method      string
	Path        string
	Handler     func(*gin.Context)
	Description string
}